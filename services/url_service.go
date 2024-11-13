package services

import (
	"sync"

	"github.com/liel-almog/url-shortener/repositories"
)

type UrlService interface {
	GetOriginalUrl(shortURL string) (string, error)
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
