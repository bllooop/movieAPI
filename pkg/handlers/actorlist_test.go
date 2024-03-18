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

func TestHandler_CreateActor(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActorList, userRole string, actor movieapi.ActorList)

	testTable := []struct {
		name                string
		inputBody           string
		inputActor          movieapi.ActorList
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"actor1", "gender":"male", "date":"1999-9-9"}`,
			inputActor: movieapi.ActorList{
				Name:      "actor1",
				Gender:    "male",
				Birthdate: "1999-9-9",
			},
			mockBehavior: func(s *mock_service.MockActorList, userRole string, actor movieapi.ActorList) {
				s.EXPECT().CreateActor(userRole, actor).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Error during execution in service",
			inputBody: `{"name":"actor1", "gender":"male", "date":"1999-9-9"}`,
			inputActor: movieapi.ActorList{
				Name:      "actor1",
				Gender:    "male",
				Birthdate: "1999-9-9",
			},
			mockBehavior: func(s *mock_service.MockActorList, userRole string, actor movieapi.ActorList) {
				s.EXPECT().CreateActor(userRole, actor).Return(0, errors.New("Internal Server Error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "Internal Server Error",
		},
		{
			name:                "Bad input",
			inputBody:           `{"role":"0"}`,
			inputActor:          movieapi.ActorList{},
			mockBehavior:        func(s *mock_service.MockActorList, userRole string, actor movieapi.ActorList) {},
			expectedStatusCode:  400,
			expectedRequestBody: "invalid input body",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			userRole := "1"
			repo := mock_service.NewMockActorList(c)
			services := &service.Service{ActorList: repo}
			handler := Handler{services}
			req, err := http.NewRequest("POST", "/api/actors/add", bytes.NewBufferString(testCase.inputBody))
			if err != nil {
				t.Fatal(err)
			}
			testCase.mockBehavior(repo, userRole, testCase.inputActor)
			w := httptest.NewRecorder()
			hh := http.HandlerFunc(handler.createActorlist)

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
