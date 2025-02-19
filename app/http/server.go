package http

import (
	"time"

	api_import "github.com/TicketsBot-cloud/import-api/app/http/endpoints/api/export"
	"github.com/TicketsBot-cloud/import-api/app/http/middleware"
	"github.com/TicketsBot-cloud/import-api/app/http/session"
	"github.com/TicketsBot-cloud/import-api/config"
	"github.com/TicketsBot/common/permission"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"go.uber.org/zap"
)

func StartServer(logger *zap.Logger) {
	logger.Info("Starting HTTP server")

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logging(logger))
	router.Use(middleware.ErrorHandler)

	router.RemoteIPHeaders = config.Conf.Server.RealIpHeaders
	if err := router.SetTrustedProxies(config.Conf.Server.TrustedProxies); err != nil {
		panic(err)
	}

	// Sessions
	session.Store = session.NewRedisStore()

	router.Use(rl(middleware.RateLimitTypeIp, 60, time.Minute))
	router.Use(rl(middleware.RateLimitTypeIp, 20, time.Second*10))
	router.Use(rl(middleware.RateLimitTypeUser, 60, time.Minute))
	router.Use(rl(middleware.RateLimitTypeGuild, 600, time.Minute*5))

	router.Use(middleware.Cors(config.Conf))

	// Metrics
	if len(config.Conf.Server.MetricHost) > 0 {
		monitor := ginmetrics.GetMonitor()
		monitor.UseWithoutExposingEndpoint(router)
		monitor.SetMetricPath("/metrics")

		metricRouter := gin.New()
		metricRouter.Use(gin.Recovery())
		metricRouter.Use(middleware.Logging(logger))

		monitor.Expose(metricRouter)

		go func() {
			if err := metricRouter.Run(config.Conf.Server.MetricHost); err != nil {
				panic(err)
			}
		}()
	}

	// util endpoints
	router.GET("/robots.txt", func(ctx *gin.Context) {
		ctx.String(200, "Disallow: /")
	})

	apiGroup := router.Group("/api", middleware.VerifyXTicketsHeader, middleware.AuthenticateToken, middleware.UpdateLastSeen)

	guildAuthApiAdmin := apiGroup.Group("/:id", middleware.AuthenticateGuild(permission.Admin))
	{
		guildAuthApiAdmin.POST("/import", api_import.ImportHandler)
		guildAuthApiAdmin.GET("/import/presign", api_import.PresignTranscriptURL)
	}

	if err := router.Run(config.Conf.Server.Host); err != nil {
		panic(err)
	}
}

func rl(rlType middleware.RateLimitType, limit int, period time.Duration) func(*gin.Context) {
	return middleware.CreateRateLimiter(rlType, limit, period)
}
