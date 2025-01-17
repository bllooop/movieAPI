package handlers

/*import (
	"errors"
	"movieapi/pkg/service"
	mock_service "movieapi/pkg/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_userIdentity(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Wrongs Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Wrong Header Value",
			headerName:           "Authorization",
			headerValue:          "Ber token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Token Parse Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(0, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.token)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			mux := http.NewServeMux()
			mux.HandleFunc("/authorization", handler.AuthMiddleware(next))
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/authorization", nil)
			req.Header.Set(test.headerName, test.headerValue)
			mux.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
func next(w http.ResponseWriter, r *http.Request) {
} */

/*
package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthorizationService struct {
	mock.Mock
}

func (m *MockAuthorizationService) ParseToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func TestAuthMiddleware(t *testing.T) {
	mockAuth := new(MockAuthorizationService)
	handler := &Handler{services: &Services{Authorization: mockAuth}}

	// Mock the "next" handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value(roleCtx).(string)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userRole))
	})

	// Middleware under test
	middleware := handler.AuthMiddleware(nextHandler)

	t.Run("Missing Authorization Header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		middleware.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Authorization token is required")
	})

	t.Run("Malformed Authorization Header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(authorizationHeader, "Bearer") // Missing token
		rec := httptest.NewRecorder()

		middleware.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid auth header")
	})

	t.Run("Invalid Token", func(t *testing.T) {
		mockAuth.On("ParseToken", "invalid-token").Return("", errors.New("invalid token"))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(authorizationHeader, "Bearer invalid-token")
		rec := httptest.NewRecorder()

		middleware.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid auth header")
		mockAuth.AssertCalled(t, "ParseToken", "invalid-token")
	})

	t.Run("Valid Token", func(t *testing.T) {
		mockAuth.On("ParseToken", "valid-token").Return("admin", nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(authorizationHeader, "Bearer valid-token")
		rec := httptest.NewRecorder()

		middleware.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "admin", strings.TrimSpace(rec.Body.String()))
		mockAuth.AssertCalled(t, "ParseToken", "valid-token")
	})
}
*/
