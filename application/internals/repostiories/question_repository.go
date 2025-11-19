package repostiories

import (
	"github.com/slava-abramoff/question-answer-api/application/internals/models"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(question *models.Question) error {
	return r.db.Create(question).Error
}

func (r *QuestionRepository) GetByID(id uint) (*models.Question, error) {
	var q models.Question
	err := r.db.Preload("Answers").First(&q, id).Error
	return &q, err
}

func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Preload("Answers").Find(&questions).Error
	return questions, err
}

func (r *QuestionRepository) Update(q *models.Question) error {
	return r.db.Save(q).Error
}

func (r *QuestionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Question{}, id).Error
}
