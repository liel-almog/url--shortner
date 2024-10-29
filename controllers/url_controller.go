package controllers

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/liel-almog/url-shortener/services"
)

type UrlController interface {
	RedirectUrl(c echo.Context) error
	Shorten(c echo.Context) error
}

type urlControllerImpl struct {
	urlService services.UrlService
}

var (
	initUrlController sync.Once
	urlController     *urlControllerImpl
)

func newUrlController() *urlControllerImpl {
	return &urlControllerImpl{
		urlService: services.GetUrlService(),
	}
}

func GetUrlController() UrlController {
	initUrlController.Do(func() {
		urlController = newUrlController()
	})

	return urlController
}

func (u *urlControllerImpl) RedirectUrl(c echo.Context) error {
	return nil
}

func (u *urlControllerImpl) Shorten(c echo.Context) error {
	return nil
}
