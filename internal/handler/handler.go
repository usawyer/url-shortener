package handler

import (
	"encoding/json"
	"fmt"
	"github.com/usawyer/url-shortener/internal/models"
	"github.com/usawyer/url-shortener/internal/service"
	"github.com/usawyer/url-shortener/internal/util"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	service *service.Service
	logger  *zap.Logger
}

func New(service *service.Service, logger *zap.Logger) *Handler {
	logger = logger.Named("Handler")
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateAlias(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	h.logger.Info("creating alias")
	var req models.UrlsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(fmt.Sprintf("failed to decode request body err=%s", err.Error()))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateAlias(r.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.JSON(w, r, http.StatusCreated, res)
}

func (h *Handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("getting long url")
	req := models.AliasRequest{
		Alias: r.PathValue("alias"),
	}

	res, err := h.service.GetUrl(r.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.JSON(w, r, http.StatusOK, res)
}
