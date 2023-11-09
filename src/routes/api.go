package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yaza-putu/online-test-dot/src/app/auth/handler"
	cat "github.com/yaza-putu/online-test-dot/src/app/category/handler"
	"github.com/yaza-putu/online-test-dot/src/config"
)

var authhandler = handler.NewAuthHandler()
var categoryHandler = cat.NewCategoryHandler()

func Api(r *echo.Echo) {
	route := r.Group("api")
	{
		route.POST("/token", authhandler.Create)
		route.PUT("/token", authhandler.Refresh)

		auth := route.Group("")
		{
			auth.Use(middleware.JWT([]byte(config.Key().Token)))

			category := auth.Group("/categories")
			{
				category.GET("", categoryHandler.All)
				category.POST("", categoryHandler.Create)
				category.GET("/:id", categoryHandler.FindById)
				category.PUT("/:id", categoryHandler.Update)
				category.DELETE("/:id", categoryHandler.Delete)
			}
		}
	}
}
