package app

import (
	"github.com/nunoonu/interview/helpers"
	"github.com/nunoonu/interview/internal/core/usecases"
	"github.com/nunoonu/interview/internal/handlers"
	"github.com/nunoonu/interview/internal/repositories"
	"log/slog"
	"os"
)

func InitializeApp() *App {

	setLogLevel()

	params := helpers.NewDBParams()
	dbCon := helpers.NewDB(params)
	appRepo := repositories.NewAppointmentRepository(dbCon)
	comRepo := repositories.NewCommentRepository(dbCon)
	hisRepo := repositories.NewHistoryRepository(dbCon)
	appUsc := usecases.NewAppointmentUsecase(appRepo, comRepo, hisRepo)
	appHdl := handlers.NewAppointmentHandler(appUsc)
	router := handlers.NewRouter(appHdl)

	httpServParams := handlers.NewHTTPServiceParams()
	httpServ := handlers.NewHTTPService(httpServParams, router)
	return ProvideApp(httpServ)
}

func setLogLevel() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(l)
}
