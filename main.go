package main

import (
	"fmt"
	"log"
	"net/http"
	router "rpc/routes"
)

// Struct untuk respons

func main() {
	// Handler untuk endpoint POST '/api'
	router.Router()
	// Mulai server pada port 8080
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
