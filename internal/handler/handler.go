package handler

import (
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

	var req models.UrlsRequest
	if ok := util.BindJSON(w, r, &req); !ok {
		return
	}

	res, err := h.service.CreateAlias(r.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())
		// какая ошибка?

		util.InternalServerError(w, r)
		return
	}

	util.JSON(w, r, http.StatusOK, res)
}

func (h *Handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	req := models.AliasRequest{
		Alias: r.PathValue("alias"),
	}

	res, err := h.service.GetUrls(r.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())

		//log.Printf("ERROR: failed to get actor err=%s\n", err.Error())
		//if errors.Is(err, ErrIdInvalid) || errors.Is(err, ErrActorNotExist) {
		//	util.NotFound(w, r)
		//	return
		//}

		util.InternalServerError(w, r)
		return
	}

	util.JSON(w, r, http.StatusOK, res)
}
