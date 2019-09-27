package go_kit

import (
	"context"
	"errors"

	"github.com/urfave/cli"
	"golang-example/cmd"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ErrLimitExceed = errors.New("request limit exceed")
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "kit",
		Aliases: []string{"kit"},

		Usage:    "Start app with go-kit",
		Action:   kitAgent,
		Category: "app",
	})
}

func kitAgent(c *cli.Context) error {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())

	app := &AppContext{
		e:    e,
		name: "ping",
		age:  40,
	}

	e.Use(bindApp(app))

	var s LocalService
	makeHTTPHandler(e, s)

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}

func bindApp(app *AppContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			newCtx := context.WithValue(req.Context(), "app", app)
			newReq := req.WithContext(newCtx)
			c.SetRequest(newReq)

			return next(c)
		}
	}
}
