package models

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Text      string `gorm:"type:text;not null"`
	CreatedAt time.Time

	Answers []Answer `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

// id: int
// text: str (текст вопроса)
// created_at: datetime
