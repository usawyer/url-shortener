package router

import (
	"github.com/usawyer/url-shortener/internal/handler"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func New(h *handler.Handler) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("POST /", h.CreateAlias)
	mux.HandleFunc("GET /{alias}", h.GetUrl)

	return &Router{mux: mux}
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
