package api

import (
	"net/http"

	_ "github.com/go-feature-flag/app-api/docs"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// New creates a new instance of the API server
func New(serverAddress string, handlers handler.Handlers) (*Server, error) {
	if handlers.HealthHandler == nil {
		return nil, handler.ErrMissingHealthHandler
	}
	if handlers.FlagAPIHandler == nil {
		return nil, handler.ErrMissingFlagAPIHandler
	}
	return &Server{
		flagHandlers:   handlers.FlagAPIHandler,
		healthHandlers: handlers.HealthHandler,
		apiEcho:        echo.New(),
		serverAddress:  serverAddress,
	}, nil
}

// Server is the struct that represents the API server
type Server struct {
	flagHandlers   *handler.FlagAPIHandler
	healthHandlers *handler.HealthHandler
	apiEcho        *echo.Echo
	serverAddress  string
}

func (s *Server) configure() {
	// config echo
	s.apiEcho.HideBanner = true
	s.apiEcho.HidePort = true
	s.apiEcho.HTTPErrorHandler = customHTTPErrorHandler

	// Middlewares
	s.apiEcho.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// init health routes
	s.apiEcho.GET("/health", s.healthHandlers.Health)

	// TODO: conditionally enable swagger based on configuration
	s.apiEcho.GET("/swagger/*", echoSwagger.WrapHandler)

	// init API routes
	groupV1 := s.apiEcho.Group("/v1")
	//groupV1.Use(echojwt.WithConfig(echojwt.Config{
	//	SigningKey: []byte("JKapFhI4Srnos8Exdxm7IOQAt7fjgJDU"),
	//}))
	groupV1.GET("/flags", s.flagHandlers.GetAllFeatureFlags)
	groupV1.GET("/flags/:id", s.flagHandlers.GetFeatureFlagByID)
	groupV1.POST("/flags", s.flagHandlers.CreateNewFlag)
	groupV1.PUT("/flags/:id", s.flagHandlers.UpdateFlagByID)
	groupV1.DELETE("/flags/:id", s.flagHandlers.DeleteFlagByID)
	groupV1.PATCH("/flags/:id/status", s.flagHandlers.UpdateFeatureFlagStatus)
}

// Start starts the API server
func (s *Server) Start() {
	s.configure()
	// start the server
	s.apiEcho.Logger.Error(s.apiEcho.Start(s.serverAddress))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.configure()
	s.apiEcho.ServeHTTP(w, r)
}

// Stop stops the API server
func (s *Server) Stop() error {
	return s.apiEcho.Close()
}
