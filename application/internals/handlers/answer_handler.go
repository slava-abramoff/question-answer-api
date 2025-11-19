package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/slava-abramoff/question-answer-api/application/internals/services"
)

type AnswerHandler struct {
	answerService *services.AnswerService
}

func NewAnswerHandler(s *services.AnswerService) *AnswerHandler {
	return &AnswerHandler{answerService: s}
}

func (h *AnswerHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	type Request struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}

	var body Request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questionIDStr := ps.ByName("id")
	qID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		http.Error(w, "invalid question id", http.StatusBadRequest)
		return
	}

	if _, err := uuid.Parse(body.UserID); err != nil {
		http.Error(w, "invalid user_id (must be uuid)", http.StatusBadRequest)
		return
	}

	a, err := h.answerService.Create(uint(qID), body.UserID, body.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AnswerHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	a, err := h.answerService.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AnswerHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	type Request struct {
		Text string `json:"text"`
	}

	idStr := ps.ByName("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var body Request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := h.answerService.Update(uint(id), body.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AnswerHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.answerService.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
