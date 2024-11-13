package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/liel-almog/url-shortener/controllers"
)

func NewUrlRoute(router *echo.Group) {
	group := router.Group("/url")

	controller := controllers.GetUrlController()
	group.GET("/:shortUrl", controller.RedirectToOriginalUrl)
	group.POST("/shorten", controller.Shorten)
}
