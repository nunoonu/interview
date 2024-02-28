package ports

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/interview/internal/core/domain"
)

type AppointmentUsecase interface {
	ListAppointments(ctx context.Context, req *domain.ListAppointmentsRequest) (*domain.ListInterviewsResponse, error)
	GetAppointment(ctx context.Context, req *domain.GetAppointmentRequest) (*domain.GetAppointmentResponse, error)
	UpdateAppointment(ctx context.Context, req *domain.UpdateAppointmentRequest) (*domain.UpdateAppointmentResponse, error)
	KeepAppointment(ctx context.Context, req *domain.KeepAppointmentRequest) (*domain.KeepAppointmentResponse, error)
	CreateAppointmentComment(ctx *gin.Context, req *domain.CreateAppointmentCommentRequest) (*domain.CreateAppointmentCommentResponse, error)
	UpdateAppointmentComment(ctx *gin.Context, req *domain.UpdateAppointmentCommentRequest) (*domain.UpdateAppointmentCommentResponse, error)
	DeleteAppointmentComment(ctx *gin.Context, req *domain.DeleteAppointmentCommentRequest) (*domain.DeleteAppointmentCommentResponse, error)
}
