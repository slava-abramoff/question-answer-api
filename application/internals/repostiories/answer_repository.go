package repostiories

import (
	"github.com/slava-abramoff/question-answer-api/application/internals/models"
	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (r *AnswerRepository) Create(a *models.Answer) error {
	return r.db.Create(a).Error
}

func (r *AnswerRepository) GetByID(id uint) (*models.Answer, error) {
	var ans models.Answer
	err := r.db.First(&ans, id).Error
	return &ans, err
}

func (r *AnswerRepository) GetByQuestion(questionID uint) ([]models.Answer, error) {
	var answers []models.Answer
	err := r.db.Where("question_id = ?", questionID).Find(&answers).Error
	return answers, err
}

func (r *AnswerRepository) Update(a *models.Answer) error {
	return r.db.Save(a).Error
}

func (r *AnswerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Answer{}, id).Error
}
