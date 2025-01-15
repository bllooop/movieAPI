package handlers

import (
	"bytes"
	"errors"
	"movieapi"
	"movieapi/pkg/service"
	mock_service "movieapi/pkg/service/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user movieapi.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           movieapi.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"test", "password":"12345", "role":"0"}`,
			inputUser: movieapi.User{
				UserName: "test",
				Password: "12345",
				Role:     "0",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user movieapi.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Error during execution in service",
			inputBody: `{"username": "test", "password":"12345", "role":"0"}`,
			inputUser: movieapi.User{
				UserName: "test",
				Password: "12345",
				Role:     "0",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user movieapi.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("Internal Server Error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "Internal Server Error",
		},
		{
			name:                "Bad input",
			inputBody:           `{"role":"0"}`,
			inputUser:           movieapi.User{},
			mockBehavior:        func(s *mock_service.MockAuthorization, user movieapi.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: "invalid input body",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(repo, testCase.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}
			req, err := http.NewRequest("POST", "/api/auth/sign-up", bytes.NewBufferString(testCase.inputBody))
			if err != nil {
				t.Fatal(err)
			}
			w := httptest.NewRecorder()
			hh := http.HandlerFunc(handler.signUp)

			hh.ServeHTTP(w, req)
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			if testCase.name == "OK" {
				assert.JSONEq(t, w.Body.String(), testCase.expectedRequestBody)
			} else {
				assert.Equal(t, strings.TrimSuffix(w.Body.String(), "\n"), testCase.expectedRequestBody)
			}
		})
	}
}
