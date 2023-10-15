package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/app/synchronizer"
	"github.com/nikita5637/quiz-fetcher/internal/config"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/elasticsearch"
	synclog "github.com/nikita5637/quiz-fetcher/internal/pkg/facade/sync_log"
	game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game"
	quiz_please_game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/quiz_please"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/sixty_seconds"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/squiz/v2"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/middleware"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/syncer"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
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
	pflag.Parse()

	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	ctx := context.Background()

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
	logger.Infof(ctx, "initialized logger with log level: %s", logLevel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		oscall := <-c
		logger.Infof(ctx, "system call recieved: %+v", oscall)
		cancel()
	}()

	cc, err := grpc.Dial(config.GetRegistratorAPIAddress(), grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(
		middleware.ModuleNameInterceptor,
	))
	if err != nil {
		logger.Fatalf(ctx, "could not connect: %s", err.Error())
	}

	gameServiceClient := gamepb.NewServiceClient(cc)

	driverName := viper.GetString("database.driver")
	db, err := storage.NewDB(ctx, driverName)
	if err != nil {
		logger.Fatalf(ctx, "connecting to DB error: %s", err.Error())
	}
	defer db.Close()

	txManager := tx.NewManager(db)

	gameTypeMatchStorage := storage.NewGameTypeMatchStorage(driverName, txManager)

	syncLogStorage := storage.NewSyncLogStorage(driverName, txManager)
	syncLogFacadeConfig := synclog.Config{
		SyncLogStorage: syncLogStorage,
		TxManager:      txManager,
	}
	syncLogFacade := synclog.New(syncLogFacadeConfig)

	placeStorage := storage.NewPlaceStorage(driverName, txManager)

	quizPleaseGamesFetcherConfig := quiz_please_game_fetcher.Config{
		GameInfoPathFormat: quiz_please_game_fetcher.GameInfoPathFormat,
		GamesListPath:      quiz_please_game_fetcher.GamesListPath,
		Name:               quiz_please_game_fetcher.FetcherName,
		PlaceStorage:       placeStorage,
		URL:                quiz_please_game_fetcher.URL,
	}
	quizPleaseGamesFetcher := quiz_please_game_fetcher.New(quizPleaseGamesFetcherConfig)

	squizFetcherConfig := squiz.Config{
		GamesListPath:        squiz.GamesListPath,
		GameTypeMatchStorage: gameTypeMatchStorage,
		Name:                 squiz.FetcherName,
		PlaceStorage:         placeStorage,
		URL:                  squiz.URL,
	}
	squizGamesFetcher := squiz.NewGamesFetcher(squizFetcherConfig)

	sixtySecondsFetcherConfig := sixty_seconds.Config{
		OpenLeagueGamesListPath:  sixty_seconds.OpenLeagueGamesListPath,
		FirstLeagueGamesListPath: sixty_seconds.FirstLeagueGamesListPath,
		Name:                     sixty_seconds.FetcherName,
		PlaceStorage:             placeStorage,
		URL:                      sixty_seconds.URL,
	}
	sixtySecondsFetcher := sixty_seconds.NewGamesFetcher(sixtySecondsFetcherConfig)

	gamesSyncerCfg := syncer.GamesSyncerConfig{
		Enabled: viper.GetBool("fetcher.games.enabled"),
		GamesFetchers: []game_fetcher.Fetcher{
			quizPleaseGamesFetcher,
			squizGamesFetcher,
			sixtySecondsFetcher,
		},
		DisabledGamesFetchers: viper.GetStringSlice("fetcher.games.disabled_leagues"),
		Period:                viper.GetDuration("fetcher.games.period") * time.Second,
		GameServiceClient:     gameServiceClient,
		SyncLogFacade:         syncLogFacade,
	}

	gamesSyncer, err := syncer.NewGamesSyncer(ctx, gamesSyncerCfg)
	if err != nil {
		logger.Fatalf(ctx, "creating games syncer error: %s", err.Error())
	}

	syncronizer := synchronizer.NewSynchronizer(
		gamesSyncer,
	)
	if err := syncronizer.Start(ctx); err != nil {
		logger.Fatalf(ctx, "starting syncronizer error: %s", err.Error())
	}
}
