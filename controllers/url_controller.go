package controllers

import (
	"errors"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/liel-almog/url-shortener/errors/apperrors"
	"github.com/liel-almog/url-shortener/models"
	"github.com/liel-almog/url-shortener/services"
)

type UrlController interface {
	RedirectToOriginalUrl(c echo.Context) error
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

// This is another way to bind!

// var model models.RedirectToOriginalUrlModel

// if err := (&echo.DefaultBinder{}).BindPathParams(c, &model); err != nil {
// 	return echo.ErrBadRequest
// }

func (u *urlControllerImpl) RedirectToOriginalUrl(c echo.Context) error {
	model := new(models.RedirectToOriginalUrlModel)

	if err := c.Bind(model); err != nil {
		return echo.ErrBadRequest
	}
	// Validate the URL max length

	originalUrl, err := u.urlService.GetOriginalUrl(model.ShortUrl)
	if err != nil {
		if errors.Is(err, apperrors.ErrUrlNotFound) {
			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.Redirect(http.StatusFound, originalUrl)
}

func (u *urlControllerImpl) Shorten(c echo.Context) error {
	model := new(models.ShortenUrlModel)

	if err := c.Bind(model); err != nil {
		return echo.ErrBadRequest
	}

	// Validate that the URL starts with https://
	// Validate the URL max length

	shortUrl, err := u.urlService.Shorten(model.OriginalUrl)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"shortUrl": shortUrl,
	})
}
