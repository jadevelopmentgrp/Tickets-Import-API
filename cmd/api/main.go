package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/jadevelopmentgrp/Tickets-Utilities/observability"
	"github.com/jadevelopmentgrp/Tickets-Utilities/secureproxy"

	archiverclient "github.com/jadevelopmentgrp/Tickets-Archiver-Client"
	app "github.com/jadevelopmentgrp/Tickets-Import-API/app/http"
	"github.com/jadevelopmentgrp/Tickets-Import-API/config"
	"github.com/jadevelopmentgrp/Tickets-Import-API/database"
	"github.com/jadevelopmentgrp/Tickets-Import-API/log"
	"github.com/jadevelopmentgrp/Tickets-Import-API/redis"
	"github.com/jadevelopmentgrp/Tickets-Import-API/rpc/cache"
	"github.com/jadevelopmentgrp/Tickets-Import-API/s3"
	"github.com/jadevelopmentgrp/Tickets-Import-API/utils"
	"github.com/jadevelopmentgrp/Tickets-Worker/i18n"
	"github.com/rxdn/gdl/rest/request"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/joho/godotenv/autoload"
)

var Logger *zap.Logger

func main() {
	startPprof()

	cfg, err := config.LoadConfig()
	utils.Must(err)
	config.Conf = cfg

	var logger *zap.Logger
	if config.Conf.JsonLogs {
		loggerConfig := zap.NewProductionConfig()
		loggerConfig.Level.SetLevel(config.Conf.LogLevel)

		logger, err = loggerConfig.Build(
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
			zap.WrapCore(observability.ZapAdapter()),
		)
	} else {
		loggerConfig := zap.NewDevelopmentConfig()
		loggerConfig.Level.SetLevel(config.Conf.LogLevel)
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		logger, err = loggerConfig.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	}

	if err != nil {
		panic(fmt.Errorf("failed to initialise zap logger: %w", err))
	}

	log.Logger = logger

	logger.Info("Connecting to database")
	database.ConnectToDatabase()

	logger.Info("Connecting to cache")
	cache.Instance = cache.NewCache()

	logger.Info("Connecting to import S3")
	s3.ConnectS3(config.Conf.S3Import.Endpoint, config.Conf.S3Import.AccessKey, config.Conf.S3Import.SecretKey)

	logger.Info("Initialising microservice clients")
	utils.ArchiverClient = archiverclient.NewArchiverClient(archiverclient.NewProxyRetriever(config.Conf.Bot.ObjectStore), []byte(config.Conf.Bot.AesKey))
	utils.SecureProxyClient = secureproxy.NewSecureProxy(config.Conf.SecureProxyUrl)

	utils.LoadEmoji()

	i18n.Init()

	if config.Conf.Bot.ProxyUrl != "" {
		request.RegisterHook(utils.ProxyHook)
	}

	logger.Info("Connecting to Redis")
	redis.Client = redis.NewRedisClient()

	logger.Info("Starting server")
	app.StartServer(logger)
}

func startPprof() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/{action}", pprof.Index)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	go func() {
		http.ListenAndServe(":6060", mux)
	}()
}
