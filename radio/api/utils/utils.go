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

func RespondError(w http.ResponseWriter, r *http.Request, status int, err interface{}) {
	if requestCtx, _ := r.Context().Value(RequestCtxKey).(*RequestContext); requestCtx != nil {
		requestCtx.Error = err
	}
	switch m := err.(type) {
	case error:
		err = Msg{Error: m.Error()}
	case string:
		err = Msg{Error: m}
	}
	Respond(w, status, err)
}

func ServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	RespondError(w, r, http.StatusInternalServerError, err)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err interface{}) {
	RespondError(w, r, http.StatusBadRequest, err)
}
