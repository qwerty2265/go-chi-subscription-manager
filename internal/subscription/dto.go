package subscription

type SubscriptionUpdateDTO struct {
	ServiceName *string    `json:"service_name"`
	Price       *int       `json:"price"`
	StartDate   *MonthYear `json:"start_date"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
}
