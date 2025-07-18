package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/qwerty2265/go-chi-subscription-manager/docs" // путь к docs, если docs в корне
	"github.com/qwerty2265/go-chi-subscription-manager/internal/subscription"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(subscriptionHandler *subscription.SubscriptionHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/subscriptions", subscription.SubscriptionRouter(*subscriptionHandler))
	})

	return r
}
