package go_kit

import "github.com/labstack/echo/v4"

type AppContext struct {
	e    *echo.Echo
	name string
	age  uint16
}
