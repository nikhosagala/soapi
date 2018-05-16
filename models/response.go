package models

// Response json model
type Response struct {
	Code    int         `json:"code"`
	Error   interface{} `json:"error"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
