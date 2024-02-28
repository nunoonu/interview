package ports

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/interview/internal/core/domain"
)

type AppointmentRepository interface {
	List(ctx context.Context, limit int, offset int) ([]domain.Appointment, *domain.PageInfo, error)
	Get(ctx context.Context, appointmentID string) (*domain.Appointment, error)
	Update(ctx context.Context, card domain.Card) error
	IsValid(ctx *gin.Context, appointmentID string) bool
	Keep(ctx context.Context, appointmentID string) error
}

type CommentRepository interface {
	Get(ctx context.Context, appointmentID string) (domain.Comments, error)
	Create(ctx *gin.Context, appointmentID string, message string, userID string) (*string, error)
	Update(ctx *gin.Context, commentID string, message string, userID string) error
	Delete(ctx *gin.Context, commentID string, userID string) error
}

type HistoryRepository interface {
	Get(ctx context.Context, appointmentID string) (domain.Histories, error)
	Create(ctx context.Context, card domain.Card, userID string) error
}
