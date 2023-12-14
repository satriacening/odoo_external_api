package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	_entities "rpc/entities"
	_pers "rpc/persistence"
)

func HandleAPIRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err, respons := _pers.HandleAPIRequest(w, r)

	fmt.Println(err, respons)

	response := _entities.Response{
		Status:  "success",
		Message: "Data processed successfully",
		Data:    respons,
	}

	// Marshal response ke format JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header ke application/json
	w.Header().Set("Content-Type", "application/json")

	// Kirim respons JSON ke client
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
