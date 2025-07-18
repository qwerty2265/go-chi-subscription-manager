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
// @Summary      Создать подписку
// @Description  Создаёт новую запись о подписке
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body      SubscriptionCreateDTO  true  "Данные подписки"
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
// @Summary      Получить все подписки пользователя
// @Description  Возвращает список всех подписок по user-id
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        user-id  query     string  true  "ID пользователя"
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
// @Summary      Получить подписку по ID
// @Description  Возвращает одну подписку по её ID
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID подписки"
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
// @Summary      Получить сумму подписок
// @Description  Считает сумму подписок пользователя за период
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        user-id      query     string  true  "ID пользователя"
// @Param        service-name query     string  false "Название сервиса"
// @Param        from         query     string  false "Дата начала (MM-YYYY)"
// @Param        to           query     string  false "Дата окончания (MM-YYYY)"
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
// @Summary      Обновить подписку
// @Description  Обновляет данные существующей подписки
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id    path      string   true  "ID подписки"
// @Param        subscription  body      SubscriptionUpdateDTO  true  "Данные подписки"
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
// @Summary      Удалить подписку
// @Description  Удаляет подписку по её ID
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID подписки"
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
