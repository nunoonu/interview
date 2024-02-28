package repositories

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/interview/internal/core/domain"
	"github.com/nunoonu/interview/internal/core/ports"
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	ID        string    `gorm:"column:id; primary_key"`
	CardName  string    `gorm:"column:card_name"`
	Status    string    `gorm:"column:status"`
	Message   string    `gorm:"column:message"`
	IsActive  bool      `gorm:"column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
}

type Appointments []Appointment

func (i Appointments) toDomain() []domain.Appointment {
	apps := make([]domain.Appointment, 0)
	for _, v := range i {
		apps = append(apps, *v.toDomain())
	}
	return apps
}

func (Appointment) TableName() string {
	return "appointment"
}

func (i Appointment) toDomain() *domain.Appointment {
	return &domain.Appointment{
		Card: domain.Card{
			AppointmentID: i.ID,
			CardName:      i.CardName,
			Status:        i.Status,
			Message:       i.Message,
		},
		IsActive:  i.IsActive,
		CreatedAt: i.CreatedAt,
		CreatedBy: i.CreatedBy,
	}
}

type appointmentDBRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) ports.AppointmentRepository {
	return &appointmentDBRepository{
		db: db,
	}
}

func (i appointmentDBRepository) List(ctx context.Context, limit int, offset int) ([]domain.Appointment, *domain.PageInfo, error) {
	var apps Appointments
	var count int64
	result := i.db.WithContext(ctx).
		Table("appointment").
		Where("is_active = ?", true).
		Count(&count).
		Scopes(paginate(limit, offset)).
		Order("created_at ASC").
		Find(&apps)

	err := result.Error
	if err != nil {
		return nil, nil, err
	}
	isLastPage := offset+limit >= int(count)
	return apps.toDomain(), &domain.PageInfo{IsLastPage: isLastPage, Total: count}, nil
}

func (a appointmentDBRepository) Get(ctx context.Context, appointmentID string) (*domain.Appointment, error) {
	var app Appointment
	if err := a.db.WithContext(ctx).
		First(&app, Appointment{
			IsActive: true,
			ID:       appointmentID,
		}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrorNotFound
		}
		return nil, err
	}
	return app.toDomain(), nil
}

func (a appointmentDBRepository) IsValid(ctx *gin.Context, appointmentID string) bool {
	var app Appointment
	if err := a.db.WithContext(ctx).
		First(&app, Appointment{
			IsActive: true,
			ID:       appointmentID,
		}).Error; err != nil {
		return false
	}
	return true
}

func (a appointmentDBRepository) Update(ctx context.Context, card domain.Card) error {
	result := a.db.WithContext(ctx).
		Model(&Appointment{}).
		Where("id = ? AND is_active = ?", card.AppointmentID, true).
		Updates(Appointment{
			CardName: card.CardName,
			Status:   card.Status,
			Message:  card.Message,
		})

	if result.RowsAffected != 1 {
		return domain.ErrorConflict
	}

	return result.Error
}

func (a appointmentDBRepository) Keep(ctx context.Context, appointmentID string) error {
	result := a.db.WithContext(ctx).
		Model(&Appointment{}).
		Where("id = ? AND is_active = ?", appointmentID, true).
		Update("is_active", false)

	if result.RowsAffected != 1 {
		return domain.ErrorConflict
	}
	return result.Error
}

func paginate(limit int, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			limit = 3
		}
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(limit)
	}
}
