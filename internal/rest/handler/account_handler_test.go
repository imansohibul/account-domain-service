package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"imansohibul.my.id/account-domain-service/entity"
	"imansohibul.my.id/account-domain-service/internal/rest/handler"
	usecasemock "imansohibul.my.id/account-domain-service/internal/rest/handler/mock"
	"imansohibul.my.id/account-domain-service/internal/rest/server"
	"imansohibul.my.id/account-domain-service/util"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*testing.T, *usecasemock.MockCreateAccountUsecase)
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:        "Create Account - Success",
			requestBody: &handler.CreateAccountRequest{Fullname: "John Doe", PhoneNumber: "+6234567890222", IdentityNumber: "3204081901970002"},
			mockSetup: func(t *testing.T, createAccountUsecase *usecasemock.MockCreateAccountUsecase) {
				createAccountUsecase.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Return(&entity.Account{AccountNumber: "123456"}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       "123456",
		},
		{
			name:        "Create Account - Invalid Request",
			requestBody: nil, // Invalid body to trigger error
			mockSetup: func(t *testing.T, createAccountUsecase *usecasemock.MockCreateAccountUsecase) {
				// No need to mock since it's an error test case
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Permintaan tidak valid", // Assuming error message from the handler
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := echo.New()
			e.Validator = server.NewCommonValidator(util.GetValidator())

			bodyBytes, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/daftar", bytes.NewReader(bodyBytes))

			if tt.requestBody != nil {
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			}
			rec := httptest.NewRecorder()

			mockCreateAccountUsecase := usecasemock.NewMockCreateAccountUsecase(ctrl)
			tt.mockSetup(t, mockCreateAccountUsecase)

			handler := handler.NewAccountHandler(mockCreateAccountUsecase, nil, nil, nil)

			c := e.NewContext(req, rec)
			err := handler.CreateAccount(c)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			assert.Equal(t, tt.expectedStatusCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)
		})
	}
}
