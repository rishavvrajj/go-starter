package main

import (
	"fmt"

	// Third-party package
	"github.com/fatih/color"

	// Local custom packages inside this module
	"github.com/rishavvrajj/go-starter/packages/auth"
	"github.com/rishavvrajj/go-starter/packages/mathutil"
	"github.com/rishavvrajj/go-starter/packages/user"
)

/*
================================================================================
PACKAGES IN GO - CHEAT SHEET & REFERENCE
================================================================================

1. Package Declaration:
   - Every Go file starts with `package <name>`.
   - Executable programs require `package main` with a `func main()`.
   - Reusable libraries define their package name (e.g. `package auth`).
   - All files in the same directory MUST share the exact same package name.

2. Exported vs Unexported Identifiers (Visibility / Encapsulation):
   - Uppercase first letter = Exported (public - accessible outside package):
     e.g., `mathutil.Add`, `user.User`, `auth.LoginWithCred`
   - Lowercase first letter = Unexported (package-private - accessible only in package):
     e.g., `user.password`, `mathutil.internalCounter`, `auth.validateToken`

3. Import Paths:
   - Standard Library: "fmt", "math", "time", "errors"
   - Third-Party: "github.com/fatih/color" (fetched with `go get` / `go mod tidy`)
   - Local packages: "<module_name>/<path_to_subfolder>"
     e.g., "github.com/rishavvrajj/go-starter/packages/mathutil"

4. Initialization (`init()` function):
   - Executes automatically when the package is imported, before `main()`.
   - Useful for setting up connections, counters, or default configurations.
================================================================================
*/

func main() {
	color.Cyan("==========================================")
	color.Cyan("       Go Packages Demo & Reference       ")
	color.Cyan("==========================================")

	// ---------------------------------------------------------
	// 1. MathUtil Package (Constants, Functions, Internal state)
	// ---------------------------------------------------------
	color.Yellow("\n[1] MathUtil Package Demo:")
	fmt.Printf("mathutil.Pi constant: %.4f\n", mathutil.Pi)
	fmt.Printf("mathutil.Version variable: %s\n", mathutil.Version)

	sum := mathutil.Add(15, 25)
	diff := mathutil.Subtract(100, 35)
	prod := mathutil.Multiply(4, 8)
	quot, err := mathutil.Divide(50, 4)

	fmt.Printf("Add(15, 25) = %.1f\n", sum)
	fmt.Printf("Subtract(100, 35) = %.1f\n", diff)
	fmt.Printf("Multiply(4, 8) = %.1f\n", prod)
	if err == nil {
		fmt.Printf("Divide(50, 4) = %.2f\n", quot)
	}

	// Division by zero error handling
	_, divErr := mathutil.Divide(10, 0)
	if divErr != nil {
		color.Red("Handled Division error: %v", divErr)
	}

	// Geometry and powers
	fmt.Printf("Circle Area (r=7): %.2f\n", mathutil.AreaOfCircle(7))
	fmt.Printf("2^10 = %.0f\n", mathutil.Power(2, 10))

	// Tracking internal unexported state via exported accessors
	fmt.Printf("Total math operations recorded: %d\n", mathutil.GetInternalCounter())
	mathutil.ResetInternalCounter()
	fmt.Printf("Internal counter after reset: %d\n", mathutil.GetInternalCounter())

	// ---------------------------------------------------------
	// 2. User Package (Structs, Constructors, Encapsulation)
	// ---------------------------------------------------------
	color.Yellow("\n[2] User Package Demo:")

	// Direct struct instantiation (only for exported fields)
	userA := user.User{
		ID:    "u001",
		Name:  "Rishav Raj",
		Email: "rishav@example.com",
		Role:  user.AdminRole,
		Profile: user.Profile{
			Bio:     "Software Engineer & Gopher",
			Website: "https://github.com/rishavvrajj",
		},
	}
	userA.GetInfo()

	// Using Constructor to set unexported fields securely
	userB := user.NewUser("u002", "Jane Doe", "jane@example.com", "secret456", user.UserRole)
	userB.GetInfo()

	// Verify password via method (password field itself is private/unexported)
	fmt.Printf("Password match for 'wrong': %t\n", userB.CheckPassword("wrong"))
	fmt.Printf("Password match for 'secret456': %t\n", userB.CheckPassword("secret456"))

	// Update password
	if err := userB.UpdatePassword("secret456", "newSecret999"); err != nil {
		color.Red("Failed to update password: %v", err)
	} else {
		color.Green("Password updated successfully!")
		fmt.Printf("Verification with updated password: %t\n", userB.CheckPassword("newSecret999"))
	}

	// ---------------------------------------------------------
	// 3. Auth Package (Session Management & Multi-package flow)
	// ---------------------------------------------------------
	color.Yellow("\n[3] Auth Package Demo:")

	// Failed login
	_, authErr := auth.LoginWithCred("admin", "wrongpassword")
	if authErr != nil {
		color.Red("Login failed as expected: %v", authErr)
	}

	// Successful login
	session, err := auth.LoginWithCred("admin", "password123")
	if err != nil {
		color.Red("Unexpected login error: %v", err)
		return
	}

	color.Green("Session token received: %s", session.Token)
	fmt.Printf("Session UserID: %s, ExpiresAt: %s\n", session.UserID, session.ExpiresAt.Format("15:04:05"))

	// Verify session existence
	retrievedSession, exists := auth.GetSession(session.Token)
	if exists {
		fmt.Printf("Retrieved active session for user: %s\n", retrievedSession.UserID)
	}

	// Logout
	auth.Logout(session.Token)
	_, stillExists := auth.GetSession(session.Token)
	fmt.Printf("Session still active after logout? %t\n", stillExists)

	color.Cyan("\n==========================================")
	color.Cyan("         Packages Demo Complete           ")
	color.Cyan("==========================================")
}