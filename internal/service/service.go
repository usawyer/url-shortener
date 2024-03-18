package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/usawyer/url-shortener/internal/models"
	"github.com/usawyer/url-shortener/internal/storage"
	"github.com/usawyer/url-shortener/internal/util"
	"go.uber.org/zap"
)

const aliasLenght = 8

type Service struct {
	storage storage.Storage
	logger  *zap.Logger
}

func New(storage storage.Storage, logger *zap.Logger) *Service {
	logger = logger.Named("Service")
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) CreateAlias(ctx context.Context, req *models.UrlsRequest) (*models.UrlsResponse, error) {
	longUrl := req.Url

	if !govalidator.IsURL(longUrl) {
		return nil, errors.New("invalid URL")
	}

	longUrl = util.StandardizeUrl(longUrl)
	hash := sha256.Sum256([]byte(longUrl))
	alias := hex.EncodeToString(hash[:])[:aliasLenght]

	url := models.Urls{
		Alias: alias,
		Url:   longUrl,
	}

	err := s.storage.AddUrl(ctx, url)

	if err != nil {
		s.logger.Warn("failed to set value in storage")
		return nil, err
	}

	res := &models.UrlsResponse{
		Alias: alias,
	}

	return res, nil
}

func (s *Service) GetUrls(ctx context.Context, req *models.AliasRequest) (*models.AliasResponse, error) {
	if len(req.Alias) != aliasLenght {
		return nil, errors.New("invalid alias input")
	}

	longUrl, err := s.storage.GetUrl(ctx, req.Alias)

	if err != nil {
		s.logger.Warn("failed to get value from storage")
		return nil, err
	}

	res := &models.AliasResponse{
		Url: longUrl,
	}

	return res, nil
}
