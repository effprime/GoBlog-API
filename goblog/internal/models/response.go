package models

import "encoding/json"

type HttpResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Payload json.RawMessage `json:"payload"`
}
