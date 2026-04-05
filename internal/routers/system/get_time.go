package handlers

import (
	"net/http"
	"time"

	"github.com/melkam59/url-shortner/internal/services"
	"github.com/labstack/echo/v4"
)

type GetSystemTimeSuccessResponseSpec struct {
	Time time.Time `json:"time" example:"2024-04-10T12:32:01.25295022+03:00"`
}

func GetSystemTime(app services.URLShortner) echo.HandlerFunc {
	return func(c echo.Context) error {

		serverTime := time.Now()

		return c.JSON(http.StatusOK, &GetSystemTimeSuccessResponseSpec{
			Time: serverTime,
		})
	}
}
