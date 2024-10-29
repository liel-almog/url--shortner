package services

import (
	"sync"

	"github.com/liel-almog/url-shortener/repositories"
)

type UrlService interface {
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
