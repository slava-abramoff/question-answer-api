package models

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	QuestionID uint   `gorm:"not null"`
	UserID     string `gorm:"type:uuid;not null"`
	Text       string `gorm:"type:text;not null"`
	CreatedAt  time.Time

	Question Question `gorm:"foreignKey:QuestionID"`
}

// id: int
// question_id: int (ссылка на Question)
// user_id: str (идентификатор пользователя, например uuid)
// text: str (текст ответа)
// created_at: datetime
