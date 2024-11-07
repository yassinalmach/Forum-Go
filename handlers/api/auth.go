package api

import (
	"net/http"
)

func MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_id")
		if err != nil {
			responseJSON(w, http.StatusUnauthorized, resp{Err: "Unauthorized"})
			return
		}

		// Call next handler
		next.ServeHTTP(w, r)
	}
}
