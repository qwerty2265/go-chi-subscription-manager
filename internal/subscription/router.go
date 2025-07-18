package subscription

import (
	"github.com/go-chi/chi/v5"
	"github.com/qwerty2265/go-chi-subscription-manager/internal/common/middleware"
)

func SubscriptionRouter(subscriptionHandler SubscriptionHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/create", middleware.ErrorWrapper(subscriptionHandler.CreateSubscription))
	r.Get("/{id}", middleware.ErrorWrapper(subscriptionHandler.GetSubscriptionByID))
	r.Get("/", middleware.ErrorWrapper(subscriptionHandler.GetAllSubscriptionsByUserID))
	r.Get("/total-price", middleware.ErrorWrapper(subscriptionHandler.GetTotalPrice))
	r.Put("/{id}", middleware.ErrorWrapper(subscriptionHandler.UpdateSubscription))
	r.Delete("/{id}", middleware.ErrorWrapper(subscriptionHandler.DeleteSubscriptionByID))

	return r
}
