package entity

import "time"

type Question struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Quest     string    `gorm:"type:varchar(255);not null" json:"quest"`
	Answer    string    `gorm:"type:text;not null" json:"answer"`
	Poin      int       `gorm:"type:int;not null" jsson:"poin"`
	TaskID    int       `gorm:"type:int;not null" json:"task_id"`
	UserID    int       `gorm:"type:int;not null" json:"user_id"`
	AdminID   int       `gorm:"type:int;not null" json:"admin_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type QuestionRequest struct {
	ID     int    `json:"id"`
	Quest  string `json:"quest"`
	Answer string `json:"answer"`
	TaskID int    `json:"task_id"`
}
