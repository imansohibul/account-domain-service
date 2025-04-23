package server

import (
	"context"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"imansohibul.my.id/account-domain-service/internal/rest/handler"
	"imansohibul.my.id/account-domain-service/util"
)

// RestServer encapsulates the Echo instance and usecases
type RestAPIServer struct {
	echo                 *echo.Echo
	createAccountUsecase handler.CreateAccountUsecase
	depositUsecase       handler.DepositUsecase
	withdrawUsecase      handler.WithdrawUsecase
	getBalanceUsecase    handler.GetBalanceUsecase
}

// NewRestAPIServer constructs the server with injected usecases
func NewRestAPIServer(
	createAccountUsecase handler.CreateAccountUsecase,
	depositUsecase handler.DepositUsecase,
	withdrawUsecase handler.WithdrawUsecase,
	getBalanceUsecase handler.GetBalanceUsecase,
) *RestAPIServer {
	e := echo.New()

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(echoprometheus.NewMiddleware("account-domain-service")) // adds middleware to gather metrics

	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	return &RestAPIServer{
		echo:                 e,
		createAccountUsecase: createAccountUsecase,
		depositUsecase:       depositUsecase,
		withdrawUsecase:      withdrawUsecase,
		getBalanceUsecase:    getBalanceUsecase,
	}
}

// setupAccountRoutes sets up the routes for account operations
// It binds the HTTP methods to the corresponding handler functions
func (s *RestAPIServer) setupAccountRoutes() {
	accountHandler := handler.NewAccountHandler(
		s.createAccountUsecase,
		s.depositUsecase,
		s.withdrawUsecase,
		s.getBalanceUsecase,
	)

	s.echo.POST("/daftar", accountHandler.CreateAccount)
	s.echo.POST("/tabung", accountHandler.Deposit)
	s.echo.POST("/tarik", accountHandler.Withdraw)
	s.echo.GET("/saldo/:account_number", accountHandler.GetBalance)
}

// Start launches the Echo HTTP server
func (s *RestAPIServer) Start(address string) error {
	s.registerValidator()
	s.setupAccountRoutes()
	return s.echo.Start(address)
}

// Shutdown gracefully shuts down the server
// It waits for all active connections to finish before closing
func (s *RestAPIServer) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *RestAPIServer) registerValidator() {
	// Register custom validation rules
	validate := util.GetValidator()
	s.echo.Validator = NewCommonValidator(validate)
}
