package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract8/internal/utils"
	"net/http"
)

func GetSyncCookie(c echo.Context) error {

	cookie, err := c.Cookie("cookie")
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верные куки")
	}

	value, err := utils.GetCookie("cookie", cookie.Value)
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верные куки")
	}

	return c.String(http.StatusOK, "Cookie value: "+*value)
}
