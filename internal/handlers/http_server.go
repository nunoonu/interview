package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPServiceParams struct {
	port string
}

func NewHTTPServiceParams() *HTTPServiceParams {
	return &HTTPServiceParams{
		port: ":1321",
	}
}

func NewHTTPService(
	params *HTTPServiceParams,
	a *gin.Engine,
) *HTTPService {
	return &HTTPService{
		params: *params,
		app:    a,
	}
}

type HTTPService struct {
	params HTTPServiceParams
	app    *gin.Engine
}

func (h *HTTPService) Start() {

	server := &http.Server{
		Addr:    h.params.port,
		Handler: h.app,
	}

	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Error while starting server:", slog.String("Err", err.Error()))
		}
	}()

	<-done
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", slog.String("Err", err.Error()))
	} else {
		slog.Info("Server shutdown completed")
	}
}
