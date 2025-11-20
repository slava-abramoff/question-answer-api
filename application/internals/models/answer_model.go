package models

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID         uint   `gorm:"primaryKey"`
	QuestionID uint   `gorm:"not null"`
	UserID     string `gorm:"type:uuid;not null"`
	Text       string `gorm:"type:text;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Question Question `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

// id: int
// question_id: int (ссылка на Question)
// user_id: str (идентификатор пользователя, например uuid)
// text: str (текст ответа)
// created_at: datetime
