package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/rs/zerolog/log"
)

func RunServer(h http.Handler, env *constants.EnvConfig) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", env.Server.Host, env.Server.Port),
		Handler:      h,
		ReadTimeout:  env.Server.ReadTimeout,
		WriteTimeout: env.Server.WriteTimeout,
		IdleTimeout:  env.Server.IdleTimeout,
	}

	go func() {
		log.Info().Str("port", srv.Addr).Msg("Starting http server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Could not start http server")
		}
	}()

	<-ctx.Done()

	stop()
	log.Info().Msg("Starting graceful shutdown (press CTRL+C to force)")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), env.Server.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Err(err).Msg("Server shutdown failed")
	}

	log.Info().Msg("Server exiting")
}
