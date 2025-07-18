package subscription

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	CreateSubscription(subscription *Subscription) (*Subscription, error)
	GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error)
	GetSubscriptionByID(id uuid.UUID) (*Subscription, error)
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
