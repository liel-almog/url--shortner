package models

type Url struct {
	ShortUrl    string `json:"shortUrl" db:"short_url"`
	OriginalUrl string `json:"originalUrl" db:"original_url"`
}

type RedirectToOriginalUrlModel struct {
	ShortUrl string `param:"shortUrl"`
}

type ShortenUrlModel struct {
	OriginalUrl string `json:"originalUrl"`
}
