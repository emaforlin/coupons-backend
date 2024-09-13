package server

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/internal/config"
	"github.com/emaforlin/coupons-app/internal/database"
	"github.com/emaforlin/coupons-app/internal/handlers"
	"github.com/emaforlin/coupons-app/internal/repositories"
	"github.com/emaforlin/coupons-app/internal/usecases"
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

	s.app.Use(middleware.Logger(), middleware.Recover(), middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 1 * time.Second}))
	s.app.Validator = &handlers.CustomValidator{V: validator.New()}

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverURL))
}

func (s *echoServer) initializeHttpHandlers() {
	public := s.app.Group("/" + s.cfg.App.ApiVersion)
	public.GET("/health", func(c echo.Context) error {
		s.app.Logger.Info("Handle /health")
		return c.String(200, "OK")
	})

	// Accounts
	accountsRepo := repositories.NewAccountMysqlRepositoryImpl(s.db)
	accountsUsecase := usecases.NewAccountUsecaseImpl(accountsRepo, s.cfg.Jwt)
	accountsHttpHandler := handlers.NewAccountHttpHandler(accountsUsecase)

	public.POST("/login", accountsHttpHandler.Login)
	public.POST("/signup", accountsHttpHandler.SignupPerson)
	public.POST("/signup/partner", accountsHttpHandler.SignupFoodPlace)

	// Coupons
	couponsRepo := repositories.NewCouponMysqlRepositoryImpl(s.db)
	couponsUsecase := usecases.NewCouponUsecase(couponsRepo)
	couponsHandler := handlers.NewCouponHttpHandler(couponsUsecase)

	public.POST("/coupons", couponsHandler.CreateCoupon)
	public.GET("/coupons", couponsHandler.RetrieveCoupons)

}

func NewEchoServer(cfg *config.Config, db database.Database) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}
