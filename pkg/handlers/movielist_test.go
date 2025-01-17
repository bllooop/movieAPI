package handlers

import (
	"bytes"
	"context"
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

// when testing comment line 31 in movielist
func TestHandler_CreateMovie(t *testing.T) {
	type mockBehavior func(s *mock_service.MockMovieList, userRole string, movie movieapi.MovieList)

	testTable := []struct {
		name                string
		inputBody           string
		inputMovie          movieapi.MovieList
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"title":"actor1", "description":"male", "rating":10, "date":"1999-9-9", "actorname": ["a", "b"]}`,
			inputMovie: movieapi.MovieList{
				Title:       "actor1",
				Description: "male",
				Rating:      10,
				Date:        "1999-9-9",
				ActorName:   []string{"a", "b"},
			},
			mockBehavior: func(s *mock_service.MockMovieList, userRole string, movie movieapi.MovieList) {
				s.EXPECT().Create(userRole, movie).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Error during execution in service",
			inputBody: `{"title":"actor1", "description":"male", "rating":10, "date":"1999-9-9", "actorname":["a", "b"]}`,
			inputMovie: movieapi.MovieList{
				Title:       "actor1",
				Description: "male",
				Rating:      10,
				Date:        "1999-9-9",
				ActorName:   []string{"a", "b"},
			},
			mockBehavior: func(s *mock_service.MockMovieList, userRole string, movie movieapi.MovieList) {
				s.EXPECT().Create(userRole, movie).Return(0, errors.New("Internal Server Error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "Internal Server Error",
		},
		{
			name:                "Bad input",
			inputBody:           `{"role":"0"}`,
			inputMovie:          movieapi.MovieList{},
			mockBehavior:        func(s *mock_service.MockMovieList, userRole string, movie movieapi.MovieList) {},
			expectedStatusCode:  400,
			expectedRequestBody: "invalid input body",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			userRole := "1"
			repo := mock_service.NewMockMovieList(c)
			services := &service.Service{MovieList: repo}
			handler := Handler{services}
			req, err := http.NewRequest("POST", "/api/movies/add", bytes.NewBufferString(testCase.inputBody))
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.WithValue(req.Context(), roleCtx, userRole)
			req = req.WithContext(ctx)
			testCase.mockBehavior(repo, userRole, testCase.inputMovie)
			w := httptest.NewRecorder()
			hh := http.HandlerFunc(handler.createMovielist)

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
