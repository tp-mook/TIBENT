package models

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Location    string    `gorm:"type:varchar(255)"`
	Date        time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
