package user

import "fmt"

// Role defines the access level for a user (custom type based on string)
type Role string

const (
	AdminRole Role = "ADMIN"
	UserRole  Role = "USER"
)

// Profile stores additional user metadata
type Profile struct {
	Bio     string
	Website string
}

// User represents a system user.
// Exported struct with a mix of exported fields and unexported internal details.
type User struct {
	ID       string
	Name     string
	Email    string
	Role     Role
	Profile  Profile
	password string // unexported field: cannot be accessed directly from outside the package
}

// NewUser creates a new User instance, handling initialization and unexported fields.
func NewUser(id, name, email, password string, role Role) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Role:     role,
		password: password, // setting unexported field within the package
	}
}

// CheckPassword verifies the provided password against the unexported password field.
// This demonstrates how unexported fields can be securely manipulated or accessed via exported methods.
func (u *User) CheckPassword(p string) bool {
	return u.password == p
}

// UpdatePassword updates the unexported password field securely.
func (u *User) UpdatePassword(oldPassword, newPassword string) error {
	if !u.CheckPassword(oldPassword) {
		return fmt.Errorf("incorrect old password")
	}
	u.password = newPassword
	return nil
}

// GetInfo prints user details (demonstrating method on user struct)
func (u *User) GetInfo() {
	fmt.Printf("User: %s (%s) [Role: %s]\n", u.Name, u.Email, u.Role)
}