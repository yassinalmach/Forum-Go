package api

import (
	"context"
	"forum/controller"
	"net/http"
)

type contextKey string
const UserIDKey contextKey = "userId"

func MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			responseJSON(w, http.StatusUnauthorized, resp{Err: "Unauthorized"})
			return
		}

		userID, err := controller.GetSession(cookie.Value)
		if err != nil {
			responseJSON(w, http.StatusUnauthorized, resp{Err: err.Error()})
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		// Call next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// GetUserID helper function to get user ID from context
func GetUserID(r *http.Request) (int64, bool) {
    userID, ok := r.Context().Value(UserIDKey).(int64)
    return userID, ok
}
