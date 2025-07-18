package subscription

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionService interface {
	CreateSubscription(subscription *Subscription) (*Subscription, error)
	GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error)
	GetSubscriptionByID(id uuid.UUID) (*Subscription, error)
	GetTotalPrice(userId uuid.UUID, serviceName string, from, to time.Time) (int, error)
	UpdateSubscription(id uuid.UUID, subscription *SubscriptionUpdateDTO) (*Subscription, error)
	DeleteSubscriptionByID(id uuid.UUID) error
}

type subscriptionService struct {
	repo SubscriptionRepository
}

func NewSubscriptionService(repo SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

// -------------------------- service methods --------------------------

func (s *subscriptionService) CreateSubscription(subscription *Subscription) (*Subscription, error) {
	if err := subscription.Validate(); err != nil {
		return nil, err
	}
	return s.repo.CreateSubscription(subscription)
}

func (s *subscriptionService) GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error) {
	return s.repo.GetAllSubscriptionsByUserID(userId)
}

func (s *subscriptionService) GetSubscriptionByID(id uuid.UUID) (*Subscription, error) {
	return s.repo.GetSubscriptionByID(id)
}

func (s *subscriptionService) GetTotalPrice(userId uuid.UUID, serviceName string, from, to time.Time) (int, error) {
	return s.repo.GetTotalPrice(userId, serviceName, from, to)
}

func (s *subscriptionService) UpdateSubscription(id uuid.UUID, subscription *SubscriptionUpdateDTO) (*Subscription, error) {
	existing, err := s.repo.GetSubscriptionByID(id)
	if err != nil {
		return nil, err
	}

	existing.UpdateFields(*subscription)
	if err := existing.Validate(); err != nil {
		return nil, err
	}

	updatedSubscription, err := s.repo.UpdateSubscription(existing)
	if err != nil {
		return nil, err
	}
	return updatedSubscription, nil
}

func (s *subscriptionService) DeleteSubscriptionByID(id uuid.UUID) error {
	return s.repo.DeleteSubscriptionByID(id)
}
