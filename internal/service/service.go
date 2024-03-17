package service

import "github.com/usawyer/url-shortener/internal/storage"

type Service struct {
	Storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{Storage: storage}
}

//func (s *Service) CreateAlias() (, error){
//
//}
//
//func (s *Service) GetUrls() (, error) {
//
//}
