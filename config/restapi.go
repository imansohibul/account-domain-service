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
	db := initPostgresDatabase(serviceConfig)

	// Initialize logger
	logger := util.NewZapLogger()

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
	)

	// Initialize Rest API server
	return server.NewRestAPIServer(createAccountUsecase, nil, nil, nil), nil
}
