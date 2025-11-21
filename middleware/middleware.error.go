package middleware

import (
	"net/http"

	"github.com/edwinjordan/e-canteen-backend/pkg/exceptions"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				exceptions.ErrorHadler(w, r, err)

			}

		}()

		next.ServeHTTP(w, r)

	})
}
