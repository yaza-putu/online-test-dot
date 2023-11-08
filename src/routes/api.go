package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yaza-putu/online-test-dot/src/app/auth/handler"
)

var authhandler = handler.NewAuthHandler()

func Api(r *echo.Echo) {
	r.POST("/token", authhandler.Create)
	r.PUT("/token", authhandler.Refresh)
}
