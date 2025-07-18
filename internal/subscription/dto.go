package subscription

import "github.com/google/uuid"

type SubscriptionCreateDTO struct {
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   MonthYear  `json:"start_date"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
}

type SubscriptionUpdateDTO struct {
	ServiceName *string    `json:"service_name"`
	Price       *int       `json:"price"`
	StartDate   *MonthYear `json:"start_date"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
}

func fromCreateDTOtoSubscription(dto *SubscriptionCreateDTO) *Subscription {
	return &Subscription{
		ID:          uuid.New(),
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		UserID:      dto.UserID,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
}
