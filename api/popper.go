package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/nlopes/slack"
	"github.com/valyala/fasthttp"
)

func Popper() echo.HandlerFunc {
	return func(c echo.Context) error {
		api := slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_ACCESS_TOKEN"))
		api.PostMessage("popper", slack.MsgOptionText("🎉", true))

		return c.String(fasthttp.StatusOK, "popper")
	}
}
