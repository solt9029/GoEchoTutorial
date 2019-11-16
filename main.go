package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/nlopes/slack"
	"github.com/solt9029/GoEchoTutorial/routes"
)

func main() {
	godotenv.Load()
	api := slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_ACCESS_TOKEN"))

	e := echo.New()
	routes.Init(e)
	e.GET("/popper", func(c echo.Context) error {
		api.PostMessage("popper", slack.MsgOptionText("🎉", true))
		return c.String(http.StatusOK, "popper")
	})

	e.Logger.Fatal(e.Start(":8091"))
}
