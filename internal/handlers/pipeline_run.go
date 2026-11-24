package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
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

func (h *PipelineRunHandler) List(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	status := query.Get("status")

	repository := query.Get("repository")

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil && query.Has("limit") {
		writeError(w, "limit must be a valid integer", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil && query.Has("offset") {
		writeError(w, "offset must be a valid integer", http.StatusBadRequest)
		return
	}

	filter := domain.ListPipelineRunsFilter{
		Status:     status,
		Repository: repository,
		Limit:      limit,
		Offset:     offset,
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	runs, err := h.Usecase.ListPipelineRuns(ctx, filter)
	if err != nil {
		if errors.Is(err, usecases.ErrInvalidStatus) {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeError(w, "could not list pipeline runs", http.StatusInternalServerError)
		log.Printf("List pipeline runs error: %v", err)
		return
	}

	writeJSON(w, http.StatusOK, runs)
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
