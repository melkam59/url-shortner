package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/melkam59/url-shortner/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostShortenURLRequestSpec struct {
	URL string `json:"url" validate:"required"`
}

func PostShortenURL(app services.URLShortner) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		requestBody := new(PostShortenURLRequestSpec)

		if err := c.Bind(&requestBody); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, &FailureMessage{
				Reason: "invalid request/JSON",
			})
		}

		if err := c.Validate(requestBody); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, &FailureMessage{
				Reason: "invalid request/JSON",
			})
		}

		userId := uuid.New()
		hashInput := fmt.Sprintf("%s-%s", requestBody.URL, userId.String())
		hash := app.ShortnerService.Sha256(hashInput)

		generatedNumber := new(big.Int).SetBytes(hash).Uint64()
		finalString := app.ShortnerService.Encode([]byte(fmt.Sprintf("%d", generatedNumber)))

		err := app.RedisService.SetKey(ctx, requestBody.URL, finalString)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &FailureMessage{
				Reason: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, &SuccessMessage{
			Message: hashInput,
		})
	}
}
