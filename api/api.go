package api

import (
	"github.com/gorilla/mux"
	"github.com/invly/invly/status"
)

func AddV1Handler(r *mux.Router) {
	r.HandleFunc("/api/v1/status", status.GetStatus).Methods("GET")
}
