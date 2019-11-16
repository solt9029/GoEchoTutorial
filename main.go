package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/solt9029/GoEchoTutorial/routes"
)

func main() {
	godotenv.Load()

	e := echo.New()
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8091"))
}
