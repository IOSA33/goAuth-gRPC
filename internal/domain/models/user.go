package models

// So what is the reason that we have domain -> models -> user.
// Defines the core business entity (User) that is shared across all layers.
// Most important is preventing direct database exposure to handlers and services.
// Serves as a universal data contract for the entire application.

type User struct {
	ID       int64
	Email    string
	PassHash []byte
}
