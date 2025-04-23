package config

import (
	"imansohibul.my.id/account-domain-service/internal/repository"
	"imansohibul.my.id/account-domain-service/internal/rest/server"
	"imansohibul.my.id/account-domain-service/internal/usecase"
	"imansohibul.my.id/account-domain-service/util"
)

func NewRestAPI() (*server.RestAPIServer, error) {
	// Load configuration
	serviceConfig, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// Initialize database connection
	db, err := initPostgresDatabase(serviceConfig)
	if err != nil {
		return nil, err
	}

	// Initialize logger
	logger := util.GetZapLogger()

	// Initialize repositories
	var (
		accountRepository          = repository.NewAccountRepository(db)
		transactionRepository      = repository.NewTransactionRepository(db)
		customerRepository         = repository.NewCustomerRepository(db)
		customerIdentityRepository = repository.NewCustomerIdentityRepository(db)
		transactionManager         = repository.NewTransactionManager(db)
	)

	// Create usecases
	var (
		createAccountUsecase = usecase.NewCreateAccountUsecase(
			accountRepository,
			transactionManager,
			customerRepository,
			customerIdentityRepository,
			transactionRepository,
			logger,
		)

		depositUsecase = usecase.NewDepositUsecase(
			accountRepository,
			transactionRepository,
			transactionManager,
			logger,
		)

		withdrawUsecase = usecase.NewWithdrawUsecase(
			accountRepository,
			transactionRepository,
			transactionManager,
			logger,
		)
	)

	// Initialize Rest API server
	return server.NewRestAPIServer(
		createAccountUsecase,
		depositUsecase,
		withdrawUsecase,
		nil,
	), nil
}
