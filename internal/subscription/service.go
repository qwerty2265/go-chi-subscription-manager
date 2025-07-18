package subscription

import "github.com/google/uuid"

type SubscriptionService interface {
	CreateSubscription(subscription *Subscription) (*Subscription, error)
	GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error)
	GetSubscriptionByID(id uuid.UUID) (*Subscription, error)
	UpdateSubscription(subscription *Subscription) (*Subscription, error)
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
	return s.repo.CreateSubscription(subscription)
}

func (s *subscriptionService) GetAllSubscriptionsByUserID(userId uuid.UUID) ([]Subscription, error) {
	return s.repo.GetAllSubscriptionsByUserID(userId)
}

func (s *subscriptionService) GetSubscriptionByID(id uuid.UUID) (*Subscription, error) {
	return s.repo.GetSubscriptionByID(id)
}

func (s *subscriptionService) UpdateSubscription(subscription *Subscription) (*Subscription, error) {
	return s.repo.UpdateSubscription(subscription)
}

func (s *subscriptionService) DeleteSubscriptionByID(id uuid.UUID) error {
	return s.repo.DeleteSubscriptionByID(id)
}
