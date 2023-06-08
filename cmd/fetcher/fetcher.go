package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/nikita5637/quiz-fetcher/internal/app/synchronizer"
	"github.com/nikita5637/quiz-fetcher/internal/config"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/elasticsearch"
	game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game"
	quiz_please_game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/quiz_please"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/sixty_seconds"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/squiz"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/middleware"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/syncer"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./config.toml", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	var err error
	err = config.ParseConfigFile(configPath)
	if err != nil {
		panic(err)
	}

	logsCombiner := &logger.Combiner{}
	logsCombiner = logsCombiner.WithWriter(os.Stdout)

	elasticLogsEnabled := config.GetValue("ElasticLogsEnabled").Bool()
	if elasticLogsEnabled {
		var elasticClient *elasticsearch.Client
		elasticClient, err = elasticsearch.New(elasticsearch.Config{
			ElasticAddress: config.GetElasticAddress(),
			ElasticIndex:   config.GetValue("ElasticIndex").String(),
		})
		if err != nil {
			panic(err)
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", "fetcher"),
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
		logger.Panic(ctx, fmt.Errorf("could not connect: %w", err))
	}

	registratorServiceClient := registrator.NewRegistratorServiceClient(cc)

	db, err := storage.NewDB(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	txManager := tx.NewManager(db)

	gameTypeMatchStorage := storage.NewGameTypeMatchStorage(txManager)

	syncLogStorage := storage.NewSyncLogStorage(txManager)
	syncerConfig := syncer.Config{
		SyncLogStorage: syncLogStorage,
	}
	syncerFacade := syncer.NewFacade(syncerConfig)

	quizPleaseGamesFetcherConfig := quiz_please_game_fetcher.Config{
		GameInfoPathFormat:       quiz_please_game_fetcher.GameInfoPathFormat,
		GamesListPath:            quiz_please_game_fetcher.GamesListPath,
		Name:                     quiz_please_game_fetcher.FetcherName,
		RegistratorServiceClient: registratorServiceClient,
		URL:                      quiz_please_game_fetcher.URL,
	}
	quizPleaseGamesFetcher := quiz_please_game_fetcher.New(quizPleaseGamesFetcherConfig)

	squizFetcherConfig := squiz.Config{
		GamesListPath:            squiz.GamesListPath,
		GameTypeMatchStorage:     gameTypeMatchStorage,
		Name:                     squiz.FetcherName,
		RegistratorServiceClient: registratorServiceClient,
		URL:                      squiz.URL,
	}
	squizGamesFetcher := squiz.NewGamesFetcher(squizFetcherConfig)

	sixtySecondsFetcherConfig := sixty_seconds.Config{
		OpenLeagueGamesListPath:  sixty_seconds.OpenLeagueGamesListPath,
		FirstLeagueGamesListPath: sixty_seconds.FirstLeagueGamesListPath,
		Name:                     sixty_seconds.FetcherName,
		RegistratorServiceClient: registratorServiceClient,
		URL:                      sixty_seconds.URL,
	}
	sixtySecondsFetcher := sixty_seconds.NewGamesFetcher(sixtySecondsFetcherConfig)

	gamesSyncerCfg := syncer.GamesSyncerConfig{
		Enabled: config.GetValue("GamesSyncerEnabled").Bool(),
		GamesFetchers: []game_fetcher.Fetcher{
			quizPleaseGamesFetcher,
			squizGamesFetcher,
			sixtySecondsFetcher,
		},
		DisabledGamesFetchers:    config.GetValue("DisabledGamesFetchers").Strings(),
		Period:                   config.GetValue("GamesSyncerPeriod").Uint64(),
		RegistratorServiceClient: registratorServiceClient,
		SyncerFacade:             *syncerFacade,
	}

	gamesSyncer, err := syncer.NewGamesSyncer(ctx, gamesSyncerCfg)
	if err != nil {
		logger.Panic(ctx, err)
	}

	syncronizer := synchronizer.NewSynchronizer(
		gamesSyncer,
	)
	if err := syncronizer.Start(ctx); err != nil {
		logger.Panic(ctx, err)
	}
}
