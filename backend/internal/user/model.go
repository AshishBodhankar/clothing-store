// Defines the User struct corresponding to your users table schema.

package user

import (
	"time"
)

const (
	RoleAdmin    = "admin"
	RoleCustomer = "customer"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
