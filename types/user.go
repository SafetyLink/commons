package types

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	AvatarID  int64     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Chat      []*Chat   `json:"chat"`
	Security  *Security `json:"security,omitempty"`
}

type Security struct {
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	DeviceID  int64     `json:"device_id"`
}
