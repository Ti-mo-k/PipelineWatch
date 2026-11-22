package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Ti-mo-k/PipelineWatch/internal/domain"
	"github.com/Ti-mo-k/PipelineWatch/internal/usecases"
)

type PipelineRunHandler struct {
	Usecase *usecases.PipelineRunUsecase
}

func NewPipelineRunHandler(
	usecase *usecases.PipelineRunUsecase,
) *PipelineRunHandler {
	return &PipelineRunHandler{
		Usecase: usecase,
	}
}

func (h *PipelineRunHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePipelineRunRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid JSON request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	run, err := h.Usecase.CreatePipelineRun(ctx, req)
	if err != nil {
		writeError(w, "could not create pipeline run", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, run)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string, status int) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
