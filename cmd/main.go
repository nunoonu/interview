package main

import (
	"github.com/nunoonu/interview/cmd/app"
	"log/slog"
	"sync"
)

func main() {
	ap := app.InitializeApp()
	slog.Info("App is initialized")
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func(a app.App) {
		defer wg.Done()
		a.HTTPServer.Start()
	}(*ap)

	wg.Wait()
}
