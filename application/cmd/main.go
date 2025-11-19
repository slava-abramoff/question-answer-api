package main

import (
	"log"
	"net/http"

	"github.com/slava-abramoff/question-answer-api/application/internals/db"
	"github.com/slava-abramoff/question-answer-api/application/internals/handlers"
	"github.com/slava-abramoff/question-answer-api/application/internals/repostiories"
	"github.com/slava-abramoff/question-answer-api/application/internals/router"
	"github.com/slava-abramoff/question-answer-api/application/internals/services"
)

func main() {
	db := db.Connect()
	if db == nil {
		log.Fatal("DB is nil, cannot create repository")
	}

	qRepo := repostiories.NewQuestionRepository(db)
	aRepo := repostiories.NewAnswerRepository(db)

	qService := services.NewQuestionService(qRepo, aRepo)
	aService := services.NewAnswerService(aRepo, qRepo)

	qHandler := handlers.NewQuestionHandler(qService)
	aHandler := handlers.NewAnswerHandler(aService)

	router := router.Setup(qHandler, aHandler)

	log.Fatal(http.ListenAndServe(":8080", router))

}
