package models

type Urls struct {
	Alias string `json:"alias" gorm:"unique;not null"`
	Url   string `json:"url"`
}

type UrlsRequest struct {
	Url string `json:"url"`
}

type UrlsResponse struct {
	Alias string `json:"alias"`
}

type AliasRequest struct {
	Alias string
}

type AliasResponse struct {
	Url string `json:"url"`
}
