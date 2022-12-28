package user

import (
	"time"
)

type User struct {
	ID           uint64    `json:"id,omitempty"`
	Email        string    `json:"email,omitempty"`
	PasswordHash string    `json:"-"`
	AtCreate     time.Time `json:"create_at,omitempty"`
	AtUpdate     time.Time `json:"update_at,omitempty"`
}
