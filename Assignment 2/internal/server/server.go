package server

import (
	"net/http"

	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/model"
	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/store"
)

func New(addr string, store *store.Store[string, string], stats *model.Stats) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewMux(store, stats),
	}
}
