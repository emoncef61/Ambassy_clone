package api

import (
	"auth-service/internal/api/handlers"
	"auth-service/internal/api/middleware"
	"auth-service/internal/api/routes"
	"auth-service/internal/config"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Api object is a router with a config, and a waitgroup to wait for backround tasks when shutting down
type API struct {
	Router *chi.Mux
	Config *config.Config
	Logger *zap.SugaredLogger
	Wg     *sync.WaitGroup
}

// takes logger, config, and handler and returns an api object
func NewAPI(log *zap.SugaredLogger, cfg *config.Config, h *handlers.Handlers) *API {

	router := chi.NewRouter()

	router.Use(middleware.RecoverPanic)

	routes.SetupRoutes(router, h)

	return &API{
		Router: router,
		Logger: log,
		Config: cfg,
		Wg:     &sync.WaitGroup{},
	}
}

// a methode that implements api object
func (a *API) Run() error {
	srv := &http.Server{
		Addr:         a.Config.Port,
		Handler:      a.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() { //go routine to recieve signals
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sign := <-quit

		a.Logger.Infow("Caught signal", "signal", sign.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		a.Logger.Infow("Completing background tasks", "addr", srv.Addr)

		a.Wg.Wait()
		shutdownError <- nil
	}()

	a.Logger.Infow("Starting server", "addr", srv.Addr, "env", a.Config.Env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	a.Logger.Infow("Stopped server", "addr", srv.Addr)

	return nil
}
