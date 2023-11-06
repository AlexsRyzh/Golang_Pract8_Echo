package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pract8/internal/router"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return
	}

	app := echo.New()

	app.GET("/cookies", router.GetSyncCookie)
	app.POST("/sync/cookies", router.SetSyncCookie)
	app.POST("/async/cookies", router.SetAsyncCookie)

	log.Fatal(app.Start(":3000"))
}
