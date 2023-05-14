package main

import (
	"log"

	"github.com/interrrp/uwuttp/uwu"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/uwu", func(c echo.Context) error {
		cfg := cfgFromCtx(c)

		text := c.QueryParam("text")

		var result string

		cacheKey := encodeUwUKey(text, cfg)
		log.Println(cacheKey)
		cached, err := cacheGet(cacheKey)
		if err != nil {
			result = uwu.UwUify(text, cfg)
			cacheSet(cacheKey, result)
		} else {
			log.Println("cache hit")
			result = cached
		}

		return c.JSON(
			200,
			map[string]any{
				"text":   result,
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
	var cfg uwu.Config
	c.Bind(&cfg)

	return cfg
}
