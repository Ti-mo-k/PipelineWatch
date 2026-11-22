package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("api/v1/health", healthHandler)
	mux.HandleFunc("/", rootHandler )

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

	writeJSON(w,resp)

}
func healthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{"status": "ok"}

	writeJSON(w,resp)
}

func writeJSON(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
