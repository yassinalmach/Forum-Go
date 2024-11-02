package controller

import (
	"fmt"
	"forum/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// this func creates a new session with an expiration time, to a specific user
// and set a session cookie
func NewSession(userId int, w http.ResponseWriter) error {
	// Delete any existing sessions for this user
	if _, err := database.DB.Exec("DELETE FROM sessions WHERE user_id = ?", userId); err != nil {
		return fmt.Errorf("error deleting session: %v", err)
	}

	// generate uuid token
	sessionID, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("error using uuid: %v", err)
	}

	// expiration time for this session
	expiredAt := time.Now().Add(24 * time.Hour)

	if _, err := database.DB.Exec("INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?) ",
		sessionID.String(), userId, expiredAt); err != nil {
		return fmt.Errorf("error inserting new session: %v", err)
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID.String(),
		Expires:  expiredAt,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
