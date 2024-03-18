package models

type Urls struct {
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

type UrlsRequest struct {
	Url string `json:"url"`
}

type UrlsResponse struct {
	Alias string `json:"alias,omitempty"`
	Error string `json:"error,omitempty"`
}

type AliasRequest struct {
	Alias string
}

type AliasResponse struct {
	Url   string `json:"url,omitempty"`
	Error string `json:"error,omitempty"`
}
