package handlers

import (
	"github.com/melkam59/url-shortner/internal/services"
	"github.com/labstack/echo/v4"
)

type SuccessMessage struct {
	Message string `json:"message"`
}

type FailureMessage struct {
	Reason string `json:"reason"`
}

func NewURLShortnerHandlers(router *echo.Group, app services.URLShortner) {
	router.GET("/:shortUrl", GetRedirectShortURL(app))
	router.POST("", PostShortenURL(app))
}
