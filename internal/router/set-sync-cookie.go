package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract8/internal/schema"
	"github.com/pract8/internal/utils"
	"net/http"
	"time"
)

func SetSyncCookie(c echo.Context) error {

	start := time.Now()

	cookieParam := c.QueryParam("cookie")
	c.SetCookie(utils.CreateCookie("cookie", cookieParam))

	utils.SyncSleep()
	duration := time.Now().Sub(start)

	return c.JSON(http.StatusOK, schema.RequestSchema{
		Status: "Куки успешно установлены",
		Time:   int64(duration.Seconds()),
	})
}
