package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/qwerty2265/go-chi-subscription-manager/internal/common"
	"gorm.io/gorm"
)

type HandlerFuncWithError func(w http.ResponseWriter, r *http.Request) error

func ErrorWrapper(next HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			var statusCode int
			var response common.Response

			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				if r.Method == http.MethodGet {
					statusCode = http.StatusOK
					response = common.Response{
						Success: true,
						Message: "record not found",
						Data:    struct{}{},
					}
				} else {
					statusCode = http.StatusNotFound
					response = common.Response{
						Success: false,
						Message: "record not found",
					}
				}
			default:
				statusCode = http.StatusBadRequest
				response = common.Response{
					Success: false,
					Message: err.Error(),
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(response)
		}
	}
}
