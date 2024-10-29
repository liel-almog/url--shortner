package models

type Url struct {
	ShortUrl    string `json:"shortUrl" db:"short_url"`
	OriginalUrl string `json:"originalUrl" db:"original_url"`
	// ClickCount  int    `json:"clickCount"`
}
