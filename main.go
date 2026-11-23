package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ti-mo-k/PipelineWatch/internal/database"
	"github.com/Ti-mo-k/PipelineWatch/internal/handlers"
	"github.com/Ti-mo-k/PipelineWatch/internal/repository"
	"github.com/Ti-mo-k/PipelineWatch/internal/usecases"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Successfully connected to PostgreSQL")

	pipelineRepo := repository.NewPipelineRunRepository(db)

	pipelineUsecase := usecases.NewPipelineRunUsecase(pipelineRepo)

	pipelineHandler := handlers.NewPipelineRunHandler(pipelineUsecase)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/health", healthHandler)
	mux.HandleFunc("POST /api/v1/pipeline-runs", pipelineHandler.Create)
	mux.HandleFunc("GET /api/v1/pipeline-runs", pipelineHandler.List)
	mux.HandleFunc("/", rootHandler)

	log.Println("pipeline server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{
		"message": "Welcome to PipelineWatch",
		"version": "1.0.0",
	}

	writeJSON(w, resp)

}
func healthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{"status": "ok"}

	writeJSON(w, resp)
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
