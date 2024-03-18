package util

import (
	"github.com/usawyer/url-shortener/internal/models"
	"strings"
)

func ToUrl(u *models.UrlsRequest) *models.Urls {
	return &models.Urls{
		Url: u.Url,
	}
}

func StandardizeUrl(url string) string {
	if url[:4] != "http" {
		genURL := strings.Replace(url, "www.", "", 1)
		trimGenURL := strings.TrimSuffix(genURL, "/")
		return "http://" + trimGenURL
	}

	if url[:5] == "https" {
		genURL := strings.Replace(url, "www.", "", 1)
		trimGenURL := strings.TrimSuffix(genURL, "/")
		return strings.Replace(trimGenURL, "https", "http", 1)
	}

	genURL := strings.Replace(url, "www.", "", 1)
	trimGenURL := strings.TrimSuffix(genURL, "/")
	return trimGenURL
}
