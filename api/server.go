package api

import (
	_ "github.com/go-feature-flag/app-api/docs"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

// New creates a new instance of the API server
func New(serverAddress string, flagHandlers handler.Flags, healthHandlers handler.Health) *Server {
	return &Server{
		flagHandlers:   flagHandlers,
		healthHandlers: healthHandlers,
		apiEcho:        echo.New(),
		serverAddress:  serverAddress,
	}
}

// Server is the struct that represents the API server
type Server struct {
	flagHandlers   handler.Flags
	healthHandlers handler.Health
	apiEcho        *echo.Echo
	serverAddress  string
}

// Start starts the API server
func (s *Server) Start() {
	// config echo
	s.apiEcho.HideBanner = true
	s.apiEcho.HidePort = true

	// Middlewares
	s.apiEcho.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// init health routes
	s.apiEcho.POST("/health", s.healthHandlers.Health)

	// init API routes
	groupV1 := s.apiEcho.Group("/v1")
	groupV1.GET("/flags", s.flagHandlers.GetAllFeatureFlags)
	groupV1.GET("/flags/:id", s.flagHandlers.GetFeatureFlagsByID)
	groupV1.POST("/flags", s.flagHandlers.CreateNewFlag)
	groupV1.PUT("/flags/:id", s.flagHandlers.UpdateFlagByID)
	groupV1.DELETE("/flags/:id", s.flagHandlers.DeleteFlagByID)
	groupV1.PATCH("/flags/:id/status", s.flagHandlers.UpdateFeatureFlagStatus)

	// TODO: conditionally enable swagger based on configuration
	s.apiEcho.GET("/swagger/*", echoSwagger.WrapHandler)

	// start the server
	s.apiEcho.Logger.Fatal(s.apiEcho.Start(s.serverAddress))
}

// Stop stops the API server
func (s *Server) Stop() error {
	return s.apiEcho.Close()
}
