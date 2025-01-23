package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Event represents the event model
type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null" validate:"required,min=3,max=255"`
	Description string    `json:"description" gorm:"type:text;not null" validate:"required,min=10,max=1000"`
	Location    string    `json:"location" gorm:"type:varchar(255);not null" validate:"required"`
	Date        time.Time `json:"date" gorm:"not null" validate:"required"`
	Category    string    `json:"category" gorm:"type:varchar(255);not null" validate:"required"`
	Organizer   string    `json:"organizer" gorm:"type:varchar(255);not null" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// バリデーション用のグローバル変数
var validate = validator.New()

// Validate validates the Event struct fields
func (e *Event) Validate() error {
	err := validate.Struct(e)
	if err != nil {
		// バリデーションエラーを返す
		return err
	}
	return nil
}
