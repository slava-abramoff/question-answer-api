package services

import (
	"fmt"

	"github.com/slava-abramoff/question-answer-api/application/internals/models"
	"github.com/slava-abramoff/question-answer-api/application/internals/repostiories"
)

type QuestionService struct {
	questionRepo *repostiories.QuestionRepository
	answerRepo   *repostiories.AnswerRepository
}

func NewQuestionService(qRepo *repostiories.QuestionRepository, aRepo *repostiories.AnswerRepository) *QuestionService {
	return &QuestionService{
		questionRepo: qRepo,
		answerRepo:   aRepo,
	}
}

func (s *QuestionService) Create(text string) (*models.Question, error) {
	if text == "" {
		return nil, fmt.Errorf("text is undefined")
	}

	q := &models.Question{Text: text}
	if err := s.questionRepo.Create(q); err != nil {
		return nil, err
	}

	return q, nil
}

func (s *QuestionService) GetByID(id uint) (*models.Question, error) {
	q, err := s.questionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *QuestionService) GetAll() ([]models.Question, error) {
	q, err := s.questionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *QuestionService) Update(id uint, text string) (*models.Question, error) {
	q := &models.Question{ID: id, Text: text}
	if err := s.questionRepo.Update(q); err != nil {
		return nil, err
	}
	return q, nil

}

func (s *QuestionService) Delete(id uint) error {
	err := s.questionRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
