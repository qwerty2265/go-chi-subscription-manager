package subscription

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	CreateSubscription(subscription *Subscription) (*Subscription, error)
	GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error)
	GetSubscriptionByID(id uuid.UUID) (*Subscription, error)
	GetTotalPrice(userId uuid.UUID, serviceName string, from, to time.Time) (int, error)
	UpdateSubscription(subscription *Subscription) (*Subscription, error)
	DeleteSubscriptionByID(id uuid.UUID) error
}

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

// -------------------------- repository methods --------------------------

func (r *subscriptionRepository) CreateSubscription(subscription *Subscription) (*Subscription, error) {
	if err := r.db.Create(subscription).Error; err != nil {
		return nil, err
	}
	return subscription, nil
}

func (r *subscriptionRepository) GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error) {
	var subscriptions []Subscription
	if err := r.db.Where("user_id = ?", userId).Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (r *subscriptionRepository) GetSubscriptionByID(id uuid.UUID) (*Subscription, error) {
	var subscription Subscription
	if err := r.db.First(&subscription, id).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *subscriptionRepository) GetTotalPrice(userId uuid.UUID, serviceName string, from, to time.Time) (int, error) {
	var total sql.NullInt64
	query := r.db.Model(&Subscription{})

	if userId != uuid.Nil {
		query = query.Where("user_id = ?", userId)
	}
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}
	if !from.IsZero() {
		query = query.Where("start_date >= ?", from)
	}
	if !to.IsZero() {
		query = query.Where("start_date <= ?", to)
	}

	if err := query.Select("SUM(price)").Scan(&total).Error; err != nil {
		return 0, err
	}
	if !total.Valid {
		return 0, nil
	}
	return int(total.Int64), nil
}

func (r *subscriptionRepository) UpdateSubscription(subscription *Subscription) (*Subscription, error) {
	if err := r.db.Save(subscription).Error; err != nil {
		return nil, err
	}
	return subscription, nil
}

func (r *subscriptionRepository) DeleteSubscriptionByID(id uuid.UUID) error {
	if err := r.db.Delete(&Subscription{}, id).Error; err != nil {
		return err
	}
	return nil
}
