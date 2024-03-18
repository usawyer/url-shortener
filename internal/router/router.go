package router

import (
	"net/http"

	"github.com/usawyer/url-shortener/internal/handler"
)

type Router struct {
	mux *http.ServeMux
}

func New(h *handler.Handler) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.CreateAlias)
	mux.HandleFunc("GET /{alias}", h.GetUrl)

	return &Router{mux: mux}
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
