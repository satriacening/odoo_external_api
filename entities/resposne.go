package entities

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Struct untuk menyimpan data yang diterima dari request
type Data struct {
	Message string `json:"message"`
}
