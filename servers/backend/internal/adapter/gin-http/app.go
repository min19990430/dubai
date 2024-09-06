package ginhttp

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auto-monitoring/internal/adapter/gin-http/middleware"
	"auto-monitoring/internal/adapter/gin-http/router"
	"auto-monitoring/internal/adapter/gin-http/validator"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Option struct {
	RunMode           string
	HTTPPort          string
	CancelTimeout     time.Duration
	ReadHeaderTimeout time.Duration
}

func defaultOption() Option {
	return Option{
		RunMode:           gin.DebugMode,
		HTTPPort:          "80",
		CancelTimeout:     5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func NewOption(conf *viper.Viper) Option {
	defaultOption := defaultOption()
	if conf.IsSet("app.run_mode") {
		defaultOption.RunMode = conf.GetString("app.run_mode")
	}
	if conf.IsSet("server.http_port") {
		defaultOption.HTTPPort = conf.GetString("server.http_port")
	}
	if conf.IsSet("app.cancel_timeout") {
		defaultOption.CancelTimeout = conf.GetDuration("app.cancel_timeout")
	}
	if conf.IsSet("app.read_header_timeout") {
		defaultOption.ReadHeaderTimeout = conf.GetDuration("app.read_header_timeout")
	}

	return defaultOption
}

type App struct {
	routes      []router.IRoute
	middlewares []middleware.IMiddleware
}

func NewApp(routes []router.IRoute, middleware []middleware.IMiddleware) *App {
	return &App{
		routes:      routes,
		middlewares: middleware,
	}
}

func Run(a *App, option Option) {
	gin.SetMode(option.RunMode)

	validator.InitValidator()

	router := gin.Default()

	for _, middleware := range a.middlewares {
		router.Use(middleware.Middleware)
	}

	for _, route := range a.routes {
		route.Setup(router.Group("/api"))
	}

	if option.RunMode == "debug" {
		pprof.Register(router)
	}

	// FIXME: 這裡應該要確定是否有反向代理
	// router.SetTrustedProxies([]string{"127.0.0.1"})
	// router.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}

	srv := &http.Server{
		Addr:              ":" + option.HTTPPort,
		Handler:           router,
		ReadHeaderTimeout: option.ReadHeaderTimeout,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), option.CancelTimeout)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	defer cancel()

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
