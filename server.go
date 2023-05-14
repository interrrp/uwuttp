package main

import (
	"github.com/labstack/echo/v4"

	"github.com/interrrp/uwuttp/uwu"
)

func main() {
	app := echo.New()

	app.GET("/uwu", func(c echo.Context) error {
		cfg := cfgFromCtx(c)

		text := c.QueryParam("text")

		return c.JSON(
			200,
			map[string]any{
				"text":   uwu.UwUify(text, cfg),
				"config": cfg,
			},
		)
	})

	port := envOr("ADDR", ":8080")
	if err := app.Start(port); err != nil {
		app.Logger.Fatal(err)
	}
}

// cfgFromCtx returns an uwu.Config from a Context.
func cfgFromCtx(c echo.Context) uwu.Config {
	cfg := uwu.NewConfig()
	c.Bind(&cfg)

	return cfg
}
