package server

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/pkg/config"
	"github.com/emaforlin/coupons-app/pkg/database"
	"github.com/emaforlin/coupons-app/pkg/handlers"
	"github.com/emaforlin/coupons-app/pkg/helpers"
	"github.com/emaforlin/coupons-app/pkg/repositories"
	"github.com/emaforlin/coupons-app/pkg/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 1 * time.Second}))
	s.app.Validator = &handlers.CustomValidator{V: validator.New()}

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Ports["web"])
	s.app.Logger.Fatal(s.app.Start(serverURL))
}

func (s *echoServer) initializeHttpHandlers() {
	// Initialize repositories, usescases, handlers here...
	repository := repositories.NewAccountMysqlRepositoryImpl(s.db)
	usecase := usecases.NewAccountUsecaseImpl(repository, s.cfg.Jwt)
	accountsHttpHandler := handlers.NewAccountHttpHandler(usecase)

	public := s.app.Group("/" + s.cfg.App.ApiVersion)
	public.GET("/health", func(c echo.Context) error {
		s.app.Logger.Info("Handle /health")
		return c.String(200, "OK")
	})

	public.POST("/login", accountsHttpHandler.Login)
	public.POST("/signup", accountsHttpHandler.SignupPerson)
	public.POST("/signup/partner", accountsHttpHandler.SignupFoodPlace)

	authorized := public.Group("/priv")
	authorized.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: s.cfg.Jwt.Secret,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.CustomJWTClaims)
		},
	}))

	authorized.POST("/verify", accountsHttpHandler.VerifyFoodPlace, handlers.CheckRole("Adm"))
}

func NewEchoServer(cfg *config.Config, db database.Database) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}
