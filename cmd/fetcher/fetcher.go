package main

import (
	"os"
	"runtime/debug"
	"time"

	_ "github.com/nikita5637/quiz-fetcher/internal/app/info"
	"github.com/nikita5637/quiz-fetcher/internal/app/synchronizer"
	"github.com/nikita5637/quiz-fetcher/internal/config"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/elasticsearch"
	synclog "github.com/nikita5637/quiz-fetcher/internal/pkg/facade/sync_log"
	quiz_please_game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/quiz_please"
	sixty_seconds_game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/sixty_seconds"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/squiz/v2"
	quiz_please_result_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/result/quiz_please"
	sixty_seconds_result_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/result/sixty_seconds"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/middleware"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	gamessyncer "github.com/nikita5637/quiz-fetcher/internal/pkg/syncer/games"
	resultssyncer "github.com/nikita5637/quiz-fetcher/internal/pkg/syncer/results"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/posener/ctxutil"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	pflag.StringP("config", "c", "", "path to config file")
	_ = viper.BindPFlag("config", pflag.Lookup("config"))
}

func main() {
	ctx := ctxutil.Interrupt()

	pflag.Parse()

	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	logsCombiner := &logger.Combiner{}
	logsCombiner = logsCombiner.WithWriter(os.Stdout)

	elasticLogsEnabled := viper.GetBool("log.elastic.enabled")
	if elasticLogsEnabled {
		var elasticClient *elasticsearch.Client
		elasticClient, err := elasticsearch.New(elasticsearch.Config{
			ElasticAddress: config.GetElasticAddress(),
			ElasticIndex:   viper.GetString("log.elastic.index"),
		})
		if err != nil {
			panic(err)
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", viper.GetString("log.module_name")),
	)))
	logger.InfoKV(ctx, "initialized logger", zap.String("log_level", logLevel.String()))

	cc, err := grpc.Dial(config.GetRegistratorAPIAddress(), grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(
		middleware.ModuleNameInterceptor,
	))
	if err != nil {
		logger.FatalKV(ctx, "dialing registrator-api error", zap.Error(err))
	}

	gameServiceClient := gamepb.NewServiceClient(cc)
	gameResultManagerServiceClient := gameresultmanagerpb.NewServiceClient(cc)

	driverName := viper.GetString("database.driver")
	db, err := storage.NewDB(ctx, driverName)
	if err != nil {
		logger.FatalKV(ctx, "connecting to DB error", zap.Error(err))
	}
	defer db.Close()

	txManager := tx.NewManager(db)

	gameTypeMatchStorage := storage.NewGameTypeMatchStorage(driverName, txManager)

	syncLogStorage := storage.NewSyncLogStorage(driverName, txManager)
	syncLogFacadeConfig := synclog.Config{
		SyncLogStorage: syncLogStorage,
		TxManager:      txManager,
	}
	syncLogsFacade := synclog.New(syncLogFacadeConfig)

	placeStorage := storage.NewPlaceStorage(driverName, txManager)

	quizPleaseGamesFetcherConfig := quiz_please_game_fetcher.Config{
		PlaceStorage: placeStorage,
	}
	quizPleaseGamesFetcher := quiz_please_game_fetcher.New(quizPleaseGamesFetcherConfig)

	squizGamesFetcherConfig := squiz.Config{
		GameTypeMatchStorage: gameTypeMatchStorage,
		PlaceStorage:         placeStorage,
	}
	squizGamesFetcher := squiz.New(squizGamesFetcherConfig)

	sixtySecondsGamesFetcherConfig := sixty_seconds_game_fetcher.Config{
		PlaceStorage: placeStorage,
	}
	sixtySecondsGamesFetcher := sixty_seconds_game_fetcher.New(sixtySecondsGamesFetcherConfig)

	gamesSyncerName := viper.GetString("synchronizer.syncer.games.name")
	logger.InfoKV(ctx, "initialize syncer", zap.String("syncer", gamesSyncerName))
	gamesSyncerCfg := gamessyncer.Config{
		DisabledGamesFetchers: viper.GetStringSlice("synchronizer.syncer.games.disabled_leagues"),
		Enabled:               viper.GetBool("synchronizer.syncer.games.enabled"),
		Fetchers: []gamessyncer.Fetcher{
			quizPleaseGamesFetcher,
			squizGamesFetcher,
			sixtySecondsGamesFetcher,
		},
		GameServiceClient: gameServiceClient,
		Name:              gamesSyncerName,
		Period:            viper.GetDuration("synchronizer.syncer.games.period") * time.Second,
	}
	gamesSyncer := gamessyncer.New(gamesSyncerCfg)
	logger.InfoKV(ctx, "syncer initialization done",
		zap.String("syncer", gamesSyncer.GetName()),
		zap.Duration("sync_period", gamesSyncer.GetPeriod()),
		zap.Bool("enabled", gamesSyncer.Enabled()))

	quizPleaseResultsFetcherConfig := quiz_please_result_fetcher.Config{
		Team: viper.GetString("synchronizer.syncer.results.team"),
	}
	quizPleaseResultsFetcher := quiz_please_result_fetcher.New(quizPleaseResultsFetcherConfig)

	sixtySecondsResultsFetcherConfig := sixty_seconds_result_fetcher.Config{
		Team: "Улица плохих снов",
	}
	sixtySecondsResultsFetcher := sixty_seconds_result_fetcher.New(sixtySecondsResultsFetcherConfig)

	resultsSyncerName := viper.GetString("synchronizer.syncer.results.name")
	logger.InfoKV(ctx, "initialize syncer", zap.String("syncer", resultsSyncerName))
	resultsSyncerConfig := resultssyncer.Config{
		DisabledResultsFetchers: viper.GetStringSlice("synchronizer.syncer.results.disabled_leagues"),
		Enabled:                 viper.GetBool("synchronizer.syncer.results.enabled"),
		Fetchers: []resultssyncer.Fetcher{
			quizPleaseResultsFetcher,
			sixtySecondsResultsFetcher,
		},
		GameServiceClient:              gameServiceClient,
		GameResultManagerServiceClient: gameResultManagerServiceClient,
		Name:                           resultsSyncerName,
		Period:                         viper.GetDuration("synchronizer.syncer.results.period") * time.Second,
	}
	resultsSyncer := resultssyncer.New(resultsSyncerConfig)
	logger.InfoKV(ctx, "initialize syncer done",
		zap.String("syncer", resultsSyncer.GetName()),
		zap.Duration("sync_period", resultsSyncer.GetPeriod()),
		zap.Bool("enabled", resultsSyncer.Enabled()))

	logger.Info(ctx, "initialize synchronizer")
	timeout := viper.GetDuration("synchronizer.sync.timeout") * time.Second
	synchronizerConfig := synchronizer.Config{
		Syncers: []synchronizer.Syncer{
			gamesSyncer,
			resultsSyncer,
		},
		SyncLogsfacade: syncLogsFacade,
		Timeout:        timeout,
	}
	synchronizer := synchronizer.New(synchronizerConfig)
	logger.InfoKV(ctx, "initialize synchronizer done", zap.Duration("timeout", timeout))

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range buildInfo.Settings {
			if setting.Key == "vcs.revision" {
				logger.InfoKV(ctx, "application started", zap.String("vcs.revision", setting.Value))
			}
		}
	}

	if err := synchronizer.Sync(ctx); err != nil {
		logger.FatalKV(ctx, "Sync error", zap.Error(err))
	}

	logger.Info(ctx, "synchronizer gracefully stopped")
}
