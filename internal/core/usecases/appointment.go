package usecases

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/interview/internal/core/domain"
	"github.com/nunoonu/interview/internal/core/ports"
	"log/slog"
)

type appointmentUsecase struct {
	appointmentRepository ports.AppointmentRepository
	commentRepository     ports.CommentRepository
	historyRepository     ports.HistoryRepository
}

func NewAppointmentUsecase(
	appointmentRepository ports.AppointmentRepository,
	commentRepository ports.CommentRepository,
	historyRepository ports.HistoryRepository) ports.AppointmentUsecase {
	return &appointmentUsecase{
		appointmentRepository: appointmentRepository,
		commentRepository:     commentRepository,
		historyRepository:     historyRepository,
	}
}

func (a appointmentUsecase) ListAppointments(ctx context.Context, req *domain.ListAppointmentsRequest) (*domain.ListInterviewsResponse, error) {
	if res, pageInfo, err := a.appointmentRepository.List(ctx, req.Limit, req.Offset); err != nil {
		slog.Error("File store fail", slog.String("Err", err.Error()))
		return nil, err
	} else {
		return &domain.ListInterviewsResponse{Appointments: res, PageInfo: pageInfo}, nil
	}
}

func (a appointmentUsecase) GetAppointment(ctx context.Context, req *domain.GetAppointmentRequest) (*domain.GetAppointmentResponse, error) {
	appID := req.AppointmentID
	app, err := a.appointmentRepository.Get(ctx, appID)
	if err != nil {
		return nil, err
	}

	com, err := a.commentRepository.Get(ctx, appID)
	if err != nil {
		return nil, err
	}

	his, err := a.historyRepository.Get(ctx, appID)
	if err != nil {
		return nil, err
	}

	return composeDomain(app, com, his), nil
}

func (a appointmentUsecase) UpdateAppointment(ctx context.Context, req *domain.UpdateAppointmentRequest) (*domain.UpdateAppointmentResponse, error) {

	err := a.appointmentRepository.Update(ctx, req.Card)
	if err != nil {
		return nil, err
	}

	err = a.historyRepository.Create(ctx, req.Card, req.UserID)
	if err != nil {
		return nil, err
	}
	return &domain.UpdateAppointmentResponse{}, nil
}

func (a appointmentUsecase) KeepAppointment(ctx context.Context, req *domain.KeepAppointmentRequest) (*domain.KeepAppointmentResponse, error) {
	err := a.appointmentRepository.Keep(ctx, req.AppointmentID)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (a appointmentUsecase) CreateAppointmentComment(ctx *gin.Context, req *domain.CreateAppointmentCommentRequest) (*domain.CreateAppointmentCommentResponse, error) {

	if ok := a.appointmentRepository.IsValid(ctx, req.AppointmentID); !ok {
		return nil, domain.ErrorConflict
	}
	id, err := a.commentRepository.Create(ctx, req.AppointmentID, req.Message, req.UserID)
	if err != nil {
		return nil, err
	}
	return &domain.CreateAppointmentCommentResponse{CommentID: id}, nil

}

func (a appointmentUsecase) UpdateAppointmentComment(ctx *gin.Context, req *domain.UpdateAppointmentCommentRequest) (*domain.UpdateAppointmentCommentResponse, error) {

	if ok := a.appointmentRepository.IsValid(ctx, req.AppointmentID); !ok {
		return nil, domain.ErrorConflict
	}
	err := a.commentRepository.Update(ctx, req.CommentID, req.Message, req.UserID)
	if err != nil {
		return nil, err
	}
	return &domain.UpdateAppointmentCommentResponse{}, nil
}

func (a appointmentUsecase) DeleteAppointmentComment(ctx *gin.Context, req *domain.DeleteAppointmentCommentRequest) (*domain.DeleteAppointmentCommentResponse, error) {

	if ok := a.appointmentRepository.IsValid(ctx, req.AppointmentID); !ok {
		return nil, domain.ErrorConflict
	}
	err := a.commentRepository.Delete(ctx, req.CommentID, req.UserID)
	if err != nil {
		return nil, err
	}
	return &domain.DeleteAppointmentCommentResponse{}, nil
}

func composeDomain(app *domain.Appointment, com domain.Comments, his domain.Histories) *domain.GetAppointmentResponse {
	return &domain.GetAppointmentResponse{
		Appointment: *app,
		Comments:    com,
		Histories:   his,
	}
}
