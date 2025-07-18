package subscription

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/qwerty2265/go-chi-subscription-manager/internal/common"
)

type SubscriptionHandler struct {
	subscriptionService SubscriptionService
}

func NewSubscriptionHandler(subscriptionService SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{subscriptionService: subscriptionService}
}

// -------------------- handler methods ----------------

func (h *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) error {
	var subscription Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		return err
	}

	createdSubscription, err := h.subscriptionService.CreateSubscription(&subscription)
	if err != nil {
		return err
	}

	response := common.Response{
		Success: true,
		Message: "subscription created",
		Data:    createdSubscription,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	return nil
}

func (h *SubscriptionHandler) GetAllSubscriptionsByUserID(w http.ResponseWriter, r *http.Request) error {
	userIdStr := r.URL.Query().Get("user-id")
	if userIdStr == "" {
		return errors.New("user-id query parameter is required")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return errors.New("invalid user-id format")
	}

	subscriptions, err := h.subscriptionService.GetAllSubscriptionsByUserID(userId)
	if err != nil {
		return err
	}

	response := common.Response{
		Success: true,
		Data:    subscriptions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return nil
}

func (h *SubscriptionHandler) GetSubscriptionByID(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return errors.New("subscription ID is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid subscription ID format")
	}

	subscription, err := h.subscriptionService.GetSubscriptionByID(id)
	if err != nil {
		return err
	}

	response := common.Response{
		Success: true,
		Data:    subscription,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return nil
}

func (h *SubscriptionHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) error {
	var subscription Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		return err
	}

	updatedSubscription, err := h.subscriptionService.UpdateSubscription(&subscription)
	if err != nil {
		return err
	}

	response := common.Response{
		Success: true,
		Data:    updatedSubscription,
		Message: "subscription updated",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return nil
}

func (h *SubscriptionHandler) DeleteSubscriptionByID(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return errors.New("subscription ID is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid subscription ID format")
	}

	if err := h.subscriptionService.DeleteSubscriptionByID(id); err != nil {
		return err
	}

	response := common.Response{
		Success: true,
		Message: "subscription deleted",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return nil
}
