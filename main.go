package main
import(
	"log"
	"net/http"
	"encoding/json"

)

func main(){
	http.HandleFunc("/health", healthHandler)

	log.Println("pipeline server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{"status":"ok"}

	

	json.NewEncoder(w).Encode(resp)
}
