package routes

import (
	"fmt"
	"net/http"
	_presentation "rpc/presentation"
)

func Router() {
	http.HandleFunc("/api", _presentation.HandleAPIRequest)
	fmt.Println()
}
