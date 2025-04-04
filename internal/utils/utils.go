package utils

import (
	"encoding/json"
	"net/http"
)

type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	json, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	json = append(json, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
	return nil
}
