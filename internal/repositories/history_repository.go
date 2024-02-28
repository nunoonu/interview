package repositories

import (
	"context"
	"github.com/nunoonu/interview/internal/core/domain"
	"github.com/nunoonu/interview/internal/core/ports"
	"gorm.io/gorm"
	"time"
)

type History struct {
	ID            string    `gorm:"column:id; default:uuid_generate_v4()"`
	AppointmentID string    `gorm:"column:appointment_id"`
	CardName      string    `gorm:"column:card_name"`
	Message       string    `gorm:"column:message"`
	Status        string    `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	CreatedBy     string    `gorm:"column:created_by"`
}

type Histories []History

func (h Histories) toDomain() domain.Histories {
	histories := make(domain.Histories, 0)
	for _, v := range h {
		histories = append(histories, *v.toDomain())
	}
	return histories
}

func (History) TableName() string {
	return "history"
}

func (h History) toDomain() *domain.History {
	return &domain.History{
		ID:            h.ID,
		AppointmentID: h.AppointmentID,
		CardName:      h.CardName,
		Message:       h.Message,
		Status:        h.Status,
		CreatedAt:     h.CreatedAt,
		CreatedBy:     h.CreatedBy,
	}
}

func (h *History) fromDomain(card domain.Card, userID string) {
	h.AppointmentID = card.AppointmentID
	h.CardName = card.CardName
	h.Message = card.Message
	h.Status = card.Status
	h.CreatedAt = time.Now()
	h.CreatedBy = userID
}

type historyDBRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) ports.HistoryRepository {
	return &historyDBRepository{
		db: db,
	}
}

func (h historyDBRepository) Get(ctx context.Context, appointmentID string) (domain.Histories, error) {
	var his Histories
	if err := h.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&his, History{
			AppointmentID: appointmentID,
		}).Error; err != nil {
		return nil, err
	}
	return his.toDomain(), nil
}

func (h historyDBRepository) Create(ctx context.Context, card domain.Card, userID string) error {
	var his History
	his.fromDomain(card, userID)
	result := h.db.WithContext(ctx).
		Create(&his)
	if result.RowsAffected != 1 {
		return domain.ErrorConflict
	}
	return result.Error
}
