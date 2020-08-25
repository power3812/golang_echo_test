package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username") //プレースホルダusernameの値取り出し
		return c.String(http.StatusOK, "Hello World "+username)
	}
}
