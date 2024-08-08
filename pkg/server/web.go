package server

import (
	"context"
	"fmt"
	"golang-example/pkg/apm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Web interface {
	Group(s string, handlers ...gin.HandlerFunc) *gin.RouterGroup
	Method(method, path string, handler gin.HandlerFunc)
	Shutdown(ctx context.Context)
	Run(ctx context.Context)
	ServeHTTP(wr http.ResponseWriter, r *http.Request)
}

type webServer struct {
	serverConfig    *http.Server
	shutdownTimeout time.Duration
	*router
}

type router struct {
	r  *gin.Engine
	mw []gin.HandlerFunc
}

func NewWebServer(ctx context.Context, config ConfigServer, nr apm.NewRelic) Web {
	webLayer := gin.New()
	webLayer.Use(gin.Recovery(), loggingMiddleware(), healthCheckMiddleware())

	if config.Env == "local" {
		webLayer.Use(localMetricsMiddleware())
	} else {
		webLayer.Use(addNewRelicToContext(nr))
	}

	cfg := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      webLayer,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	r := &router{r: webLayer}
	srv := &webServer{cfg, config.ShutdownTimeout, r}
	return srv
}

func addNewRelicToContext(nr apm.NewRelic) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		tx := nr.GetApplication().StartTransaction(c.FullPath())
		defer tx.End()

		ctx = context.WithValue(ctx, "newRelicTransaction", tx)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func localMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		fmt.Printf("Local Metrics - Request: %s %s took %v\n", c.Request.Method, c.Request.URL, latency)
	}
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		fmt.Printf("Request: %s %s took %v\n", c.Request.Method, c.Request.URL, latency)
	}
}

func healthCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/health" {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (w *webServer) Run(ctx context.Context) {
	go func() {
		if err := w.serverConfig.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	w.Shutdown(context.Background())
}

func (w *webServer) Shutdown(ctx context.Context) {
	shutdownCtx, cancel := context.WithTimeout(ctx, w.shutdownTimeout)
	defer cancel()
	if err := w.serverConfig.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Server forced to shutdown: %s\n", err)
	}
}

func (w *webServer) Method(method, path string, handler gin.HandlerFunc) {
	w.r.Handle(method, path, handler)
}

func (w *webServer) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	w.r.ServeHTTP(wr, r)
}

func (w *webServer) Group(s string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return w.router.r.Group(s, handlers...)
}
