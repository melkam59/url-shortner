package handlers

import (
	"net/http"

	"github.com/melkam59/url-shortner/internal/services"
	"github.com/labstack/echo/v4"
)

func GetRedirectShortURL(app services.URLShortner) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		shortUrl := c.Param("shortUrl")
		if shortUrl == "" {
			return c.JSON(http.StatusBadRequest, &FailureMessage{
				Reason: "short url can't be empty",
			})
		}

		originalUrl, err := app.RedisService.GetValue(ctx, shortUrl)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &FailureMessage{
				Reason: err.Error(),
			})
		}

		return c.Redirect(http.StatusFound, originalUrl)
	}
}
