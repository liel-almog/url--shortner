package services

import (
	"fmt"
	"strings"
	"sync"

	"github.com/liel-almog/url-shortener/consts"
	"github.com/liel-almog/url-shortener/models"
	"github.com/liel-almog/url-shortener/repositories"
)

type UrlService interface {
	GetOriginalUrl(shortURL string) (string, error)
	Shorten(originalURL string) (string, error)
}

type urlServiceImpl struct {
	urlRepository repositories.UrlRepository
}

var (
	initUrlService sync.Once
	urlService     *urlServiceImpl
)

func newUrlService() *urlServiceImpl {
	return &urlServiceImpl{
		urlRepository: repositories.GetUrlRepository(),
	}
}

func GetUrlService() UrlService {
	initUrlService.Do(func() {
		urlService = newUrlService()
	})

	return urlService
}

func (service *urlServiceImpl) GetOriginalUrl(shortURL string) (string, error) {
	originalUrl, err := service.urlRepository.FindOriginalUrl(shortURL)

	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (service *urlServiceImpl) Shorten(originalUrl string) (string, error) {
	parts := strings.Split(originalUrl, "https://")
	domain := parts[1]
	if domain == "" {
		return "", nil
	}

	// Join two stings to create a short URL and add / between them using fmt package
	shortUrl := fmt.Sprintf("%s/%s", consts.BaseShortUrl, RandStringBytesMaskImprSrcSB(6))

	url := models.Url{
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
	}

	err := service.urlRepository.InsertShortenUrl(url)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}
