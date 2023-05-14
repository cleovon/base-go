package models

import "encoding/json"

type HealthCheck struct {
	Status string `json:"status"`
}

func (h HealthCheck) ToBytes() []byte {
	structInBytes, _ := json.Marshal(h)
	return structInBytes
}
