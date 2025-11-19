package services

import (
	"fmt"

	"github.com/slava-abramoff/question-answer-api/application/internals/models"
	"github.com/slava-abramoff/question-answer-api/application/internals/repostiories"
)

type AnswerService struct {
	answerRepo   *repostiories.AnswerRepository
	questionRepo *repostiories.QuestionRepository
}

func NewAnswerService(aRepo *repostiories.AnswerRepository, qRepo *repostiories.QuestionRepository) *AnswerService {
	return &AnswerService{
		answerRepo:   aRepo,
		questionRepo: qRepo,
	}
}

func (s *AnswerService) Create(questionID uint, userID, text string) (*models.Answer, error) {
	if text == "" {
		return nil, fmt.Errorf("text is undefined")
	}
	if userID == "" {
		return nil, fmt.Errorf("userID is undefined")
	}

	_, err := s.questionRepo.GetByID(questionID)
	if err != nil {
		return nil, fmt.Errorf("question not found")
	}

	answer := &models.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
	}

	if err := s.answerRepo.Create(answer); err != nil {
		return nil, err
	}

	return answer, nil
}

func (s *AnswerService) GetByID(id uint) (*models.Answer, error) {
	a, err := s.answerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AnswerService) GetAll(id uint) ([]models.Answer, error) {
	a, err := s.answerRepo.GetByQuestion(id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AnswerService) Update(id uint, text string) (*models.Answer, error) {
	if text == "" {
		return nil, fmt.Errorf("text is undefined")
	}

	answer := &models.Answer{
		ID:   id,
		Text: text,
	}

	if err := s.answerRepo.Update(answer); err != nil {
		return nil, err
	}

	return answer, nil
}

// Удалить ответ
func (s *AnswerService) Delete(id uint) error {
	return s.answerRepo.Delete(id)
}
