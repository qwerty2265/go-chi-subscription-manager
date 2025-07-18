package subscription

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

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

// CreateSubscription godoc
// @Summary      Create subscription
// @Description  Creates a new subscription record
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body      SubscriptionCreateDTO  true  "Subscription data"
// @Success      201  {object}  common.Response
// @Router       /api/subscriptions [post]
func (h *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) error {
	var subscription SubscriptionCreateDTO
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

// GetAllSubscriptionsByUserID godoc
// @Summary      Get all user subscriptions
// @Description  Returns a list of all subscriptions by user-id
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        user-id  query     string  true  "User ID"
// @Success      200  {object}  common.Response
// @Router       /api/subscriptions [get]
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

// GetSubscriptionByID godoc
// @Summary      Get subscription by ID
// @Description  Returns a subscription by its ID
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Success      200  {object}  common.Response
// @Router       /api/subscriptions/{id} [get]
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

// GetTotalPrice godoc
// @Summary      Get total subscription price
// @Description  Calculates the total price of user subscriptions for a period
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        user-id      query     string  true  "User ID"
// @Param        service-name query     string  false "Service name"
// @Param        from         query     string  false "Start date (MM-YYYY)"
// @Param        to           query     string  false "End date (MM-YYYY)"
// @Success      200  {object}  common.Response
// @Router       /api/subscriptions/total-price [get]
func (h *SubscriptionHandler) GetTotalPrice(w http.ResponseWriter, r *http.Request) error {
	userIdStr := r.URL.Query().Get("user-id")
	serviceName := r.URL.Query().Get("service-name")
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	var userId uuid.UUID
	if userIdStr == "" {
		return errors.New("user-id query parameter is required")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return errors.New("invalid user-id format")
	}

	var from, to time.Time
	if fromStr != "" {
		t, err := time.Parse(monthYearLayout, fromStr)
		if err != nil {
			return errors.New("invalid from date format")
		}
		from = t
	}

	if toStr != "" {
		t, err := time.Parse(monthYearLayout, toStr)
		if err != nil {
			return errors.New("invalid to date format")
		}
		to = t
	}

	totalPrice, err := h.subscriptionService.GetTotalPrice(userId, serviceName, from, to)
	if err != nil {
		return err
	}

	response := common.Response{
		Success: true,
		Data:    totalPrice,
		Message: "total price calculated",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return nil
}

// UpdateSubscription godoc
// @Summary      Update subscription
// @Description  Updates an existing subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id    path      string   true  "Subscription ID"
// @Param        subscription  body      SubscriptionUpdateDTO  true  "Subscription data"
// @Success      200  {object}  common.Response
// @Router       /api/subscriptions/{id} [put]
func (h *SubscriptionHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return errors.New("subscription ID is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid subscription ID format")
	}

	var subscription SubscriptionUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		return err
	}

	updatedSubscription, err := h.subscriptionService.UpdateSubscription(id, &subscription)
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

// DeleteSubscriptionByID godoc
// @Summary      Delete subscription
// @Description  Deletes a subscription by its ID
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Success      200  {object}  common.Response
// @Router       /api/subscriptions/{id} [delete]
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
