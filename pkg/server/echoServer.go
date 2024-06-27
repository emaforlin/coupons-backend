package server

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/pkg/config"
	"github.com/emaforlin/coupons-app/pkg/database"
	"github.com/emaforlin/coupons-app/pkg/handlers"
	"github.com/emaforlin/coupons-app/pkg/repositories"
	"github.com/emaforlin/coupons-app/pkg/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoServer struct {
	app *echo.Echo
	db  database.Database
	cfg *config.Config
}

func (s *echoServer) Start() {
	s.initializeHttpHandlers()

	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 5 * time.Second}))

	s.app.Validator = &handlers.CustomValidator{V: validator.New()}

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Ports["web"])
	s.app.Logger.Fatal(s.app.Start(serverURL))
}

func (s *echoServer) initializeHttpHandlers() {
	// Initialize repositories, usescases, handlers here...
	repository := repositories.NewAccountMysqlRepositoryImpl(s.db)
	usecase := usecases.NewAccountUsecaseImpl(repository)
	accountsHttpHandler := handlers.NewAccountHttpHandler(usecase)

	// 	######## public routers ########

	s.app.GET("/health", func(c echo.Context) error {
		s.app.Logger.Info("Handle /health")
		return c.String(200, "OK")
	})

	accountsRouter := s.app.Group(s.cfg.App.ApiVersion + "/accounts")
	accountsRouter.POST("/signup", accountsHttpHandler.SignupPerson)
	accountsRouter.POST("/signup/partner", accountsHttpHandler.SignupFoodPlace)

}

func NewEchoServer(cfg *config.Config, db database.Database) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}
