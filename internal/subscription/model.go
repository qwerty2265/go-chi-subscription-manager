package subscription

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id"`
	ServiceName string     `gorm:"not null" json:"service_name"`
	Price       int        `gorm:"not null" json:"price"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;<-:create" json:"user_id"`
	StartDate   MonthYear  `gorm:"not null" json:"start_date"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (s *Subscription) Validate() (err error) {
	if s.EndDate != nil && s.EndDate.ToTime().Before(s.StartDate.ToTime()) {
		return errors.New("end date cannot be before start date")
	}

	if s.Price < 0 {
		return errors.New("price cannot be negative")
	}
	return nil
}

func (s *Subscription) UpdateFields(updatedData SubscriptionUpdateDTO) {
	if updatedData.ServiceName != nil {
		s.ServiceName = *updatedData.ServiceName
	}
	if updatedData.Price != nil {
		s.Price = *updatedData.Price
	}
	if updatedData.StartDate != nil {
		s.StartDate = *updatedData.StartDate
	}
	s.EndDate = updatedData.EndDate
}
