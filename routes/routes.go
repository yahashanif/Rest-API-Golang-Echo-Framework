package routes

import (
	"net/http"
	"rest-api-go-echo/controllers"
	"rest-api-go-echo/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "haaaa")
	})

	authRoutes := e.Group("/api/auth")
	{
		authRoutes.POST("/login", controllers.CheckLogin)
		authRoutes.POST("/register", controllers.Register)

	}

	userRoutes := e.Group("/api/user", middleware.IsAuth)
	{
		userRoutes.GET("/pegawai", controllers.FetchAllPegawai)
		userRoutes.POST("/pegawai", controllers.StorePegawai)
		userRoutes.PUT("/pegawai", controllers.UpdatePegawai)
		userRoutes.DELETE("/pegawai", controllers.DeletePegawai)

		userRoutes.GET("/generate-hash/:password", controllers.GenerateHashPassword)
		userRoutes.POST("/upload", controllers.HandleUpload)
		userRoutes.Static("/static", "files")
	}

	return e
}
