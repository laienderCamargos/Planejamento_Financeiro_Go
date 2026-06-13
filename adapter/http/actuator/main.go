package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody Struct do Status do serviço

type HealthBody struct {
	Status string `json:"status"`
}

// HealthBody  Status do serviço
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthBody{
		"alive",
	}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
