package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/slava-abramoff/question-answer-api/application/internals/handlers"
)

func Setup(q *handlers.QuestionHandler, a *handlers.AnswerHandler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/questions", q.Create)
	router.GET("/questions/:id", q.GetByID)
	router.GET("/questions", q.GetAll)
	router.PUT("/questions/:id", q.Update)
	router.DELETE("/questions/:id", q.Delete)

	router.POST("/questions/:id/answers", a.Create)
	router.GET("/answers/:id", a.GetByID)
	router.DELETE("/answers/:id", a.Delete)

	return router
}

// Вопросы (Questions):
// GET /questions/ — список всех вопросов
// POST /questions/ — создать новый вопрос
// GET /questions/{id} — получить вопрос и все ответы на него
// DELETE /questions/{id} — удалить вопрос (вместе с ответами)

// Ответы (Answers):
// POST /questions/{id}/answers/ — добавить ответ к вопросу
// GET /answers/{id} — получить конкретный ответ
// DELETE /answers/{id} — удалить ответ
