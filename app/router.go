package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qwerty2265/go-chi-subscription-manager/internal/subscription"
)

func NewRouter(subscriptionHandler *subscription.SubscriptionHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/subscriptions", subscription.SubscriptionRouter(*subscriptionHandler))
	})

	return r
}
