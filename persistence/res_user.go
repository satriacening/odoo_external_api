package persistence

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	_entities "rpc/entities"
)

func HandleAPIRequest(w http.ResponseWriter, r *http.Request) (error, string) {

	var receivedData _entities.Data
	err := json.NewDecoder(r.Body).Decode(&receivedData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return err, ""
	}
	fmt.Printf("Message received: %s\n", receivedData.Message)
	response := receivedData.Message
	// response := prepareResponse()
	fmt.Println("type ", reflect.TypeOf(err), reflect.TypeOf(response))

	// sendJSONResponse(w, response)
	return err, response
}
