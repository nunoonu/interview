package app

import (
	"github.com/nunoonu/interview/internal/handlers"
)

func ProvideApp(httpServer *handlers.HTTPService) *App {
	return &App{
		HTTPServer: *httpServer,
	}
}

type App struct {
	HTTPServer handlers.HTTPService
}
