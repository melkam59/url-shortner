package main

import (
	"net/http"

	shortner_handlers "github.com/melkam59/url-shortner/internal/routers/shortner"
	system_handlers "github.com/melkam59/url-shortner/internal/routers/system"
	"github.com/melkam59/url-shortner/internal/services"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	urlShortner, err := services.NewURLShortner()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	api := e.Group("/api")
	apiV1 := api.Group("/v1")

	v1Shortner := apiV1.Group("/shortner")
	shortner_handlers.NewURLShortnerHandlers(v1Shortner, urlShortner)

	v1System := apiV1.Group("/system")
	system_handlers.NewSystemHandlers(v1System, urlShortner)

	e.Logger.Fatal(e.Start(":8000"))
}
