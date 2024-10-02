package main

import (
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	data, err := dao.NewPostgresDao("localhost", 5432, "gofeatureflag", "goff-user", "my-secret-pw")
	if err != nil {
		panic(err)
	}

	handlers := handler.NewFlags(data)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	groupV1 := e.Group("/v1")
	groupV1.GET("/flags", handlers.GetAllFeatureFlags)
	groupV1.GET("/flags/:id", handlers.GetFeatureFlagsByID)
	groupV1.POST("/flags", handlers.CreateNewFlag)
	groupV1.PUT("/flags/:id", handlers.UpdateFlagByID)
	groupV1.DELETE("/flags/:id", handlers.DeleteFlagByID)
	groupV1.PATCH("/flags/:id/status", handlers.UpdateFeatureFlagStatus)
	e.Logger.Fatal(e.Start(":3001"))
}
