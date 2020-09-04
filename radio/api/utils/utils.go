package utils

import (
	"net/http"

	"github.com/segmentio/encoding/json"
)

type Msg struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func RespondError(w http.ResponseWriter, status int, data interface{}) {
	switch m := data.(type) {
	case error:
		data = Msg{Error: m.Error()}
	case string:
		data = Msg{Error: m}
	}
	Respond(w, status, data)
}

func ServerError(w http.ResponseWriter, data interface{}) {
	RespondError(w, http.StatusInternalServerError, data)
}

func BadRequest(w http.ResponseWriter, data interface{}) {
	RespondError(w, http.StatusBadRequest, data)
}
