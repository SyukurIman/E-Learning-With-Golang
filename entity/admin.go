package entity

import "time"

type Admin struct {
	ID        int       `gorm:"primarykey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
