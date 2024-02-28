package repositories

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/interview/internal/core/domain"
	"github.com/nunoonu/interview/internal/core/ports"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID            string    `gorm:"column:id; default:uuid_generate_v4()"`
	AppointmentID string    `gorm:"column:appointment_id"`
	Message       string    `gorm:"column:message"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	CreatedBy     string    `gorm:"column:created_by"`
}

type Comments []Comment

func (c Comments) toDomain() domain.Comments {
	comments := make(domain.Comments, 0)
	for _, v := range c {
		comments = append(comments, *v.toDomain())
	}
	return comments
}

func (Comment) TableName() string {
	return "comment"
}

func (i Comment) toDomain() *domain.Comment {
	return &domain.Comment{
		ID:            i.ID,
		AppointmentID: i.AppointmentID,
		Message:       i.Message,
		CreatedAt:     i.CreatedAt,
		CreatedBy:     i.CreatedBy,
	}
}

func (i *Comment) FromDomain(appointmentID string, message string, userID string) {
	i.AppointmentID = appointmentID
	i.Message = message
	i.CreatedBy = userID
}

type commentDBRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ports.CommentRepository {
	return &commentDBRepository{
		db: db,
	}
}

func (c commentDBRepository) Get(ctx context.Context, appointmentID string) (domain.Comments, error) {
	var com Comments
	if err := c.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&com, Comment{
			AppointmentID: appointmentID,
		}).Error; err != nil {
		return nil, err
	}
	return com.toDomain(), nil
}

func (c commentDBRepository) Create(ctx *gin.Context, appointmentID string, message string, userID string) (*string, error) {
	var com Comment
	com.FromDomain(appointmentID, message, userID)
	result := c.db.WithContext(ctx).
		Create(&com)
	if result.RowsAffected != 1 {
		return nil, domain.ErrorConflict
	}
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &com.ID, nil
}

func (c commentDBRepository) Update(ctx *gin.Context, commentID string, message string, userID string) error {
	result := c.db.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ? AND created_by = ?", commentID, userID).
		Updates(Comment{
			Message: message,
		})
	if result.RowsAffected != 1 {
		return domain.ErrorConflict
	}
	return nil
}

func (c commentDBRepository) Delete(ctx *gin.Context, commentID string, userID string) error {
	result := c.db.WithContext(ctx).
		Where("created_by = ?", userID).
		Delete(&Comment{ID: commentID})
	if result.RowsAffected != 1 {
		return domain.ErrorConflict
	}
	return result.Error
}
