package subscription

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ServiceName string     `gorm:"not null" json:"service_name"`
	Price       int        `gorm:"not null" json:"price"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	StartDate   time.Time  `gorm:"not null" json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (s *Subscription) UpdateFields(updatedData Subscription) error {
	if updatedData.ServiceName != "" {
		s.ServiceName = updatedData.ServiceName
	}
	if updatedData.Price > 0 {
		s.Price = updatedData.Price
	}
	if !updatedData.StartDate.IsZero() {
		s.StartDate = updatedData.StartDate
	}
	if updatedData.EndDate != nil {
		s.EndDate = updatedData.EndDate
	}
	s.UpdatedAt = time.Now()
	return nil
}
