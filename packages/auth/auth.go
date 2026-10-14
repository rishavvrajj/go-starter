package auth

import (
	"errors"
	"fmt"
	"time"
)

// Session represents a user session.
// The fields are exported (Uppercase) so they can be accessed by the main package.
type Session struct {
	Token     string
	ExpiresAt time.Time
	UserID    string
}

// internal session store (unexported - only accessible within this package)
var activeSessions = make(map[string]Session)

// ErrInvalidCredentials is a custom error for authentication failures
var ErrInvalidCredentials = errors.New("invalid username or password")

// LoginWithCred authenticates a user and creates a session.
// This is an exported function.
func LoginWithCred(username, password string) (*Session, error) {
	fmt.Printf("Attempting to login user: %s\n", username)

	// Mock authentication logic
	if username == "admin" && password == "password123" {
		token := fmt.Sprintf("tok_%d", time.Now().UnixNano())
		session := Session{
			Token:     token,
			ExpiresAt: time.Now().Add(1 * time.Hour),
			UserID:    "user_123",
		}

		// Store session internally
		activeSessions[token] = session
		fmt.Println("Login successful!")
		return &session, nil
	}

	return nil, ErrInvalidCredentials
}

// GetSession retrieves a session by token.
// Demonstrates how to interact with unexported package state.
func GetSession(token string) (Session, bool) {
	session, exists := activeSessions[token]
	return session, exists
}

// Logout removes a session from the internal store.
func Logout(token string) {
	delete(activeSessions, token)
	fmt.Println("User logged out.")
}

// internal helper function to validate token (unexported)
func validateToken(token string) bool {
	_, exists := activeSessions[token]
	return exists
}