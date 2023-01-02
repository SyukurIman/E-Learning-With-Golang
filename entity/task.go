package entity

import "time"

type Task struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	AdminID   int       `json:"admin_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type TaskData struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Question []Question `json:"question"`
}
