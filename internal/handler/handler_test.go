package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"workshop/internal/api"
	"workshop/internal/api/mocks"
)

func TestHandler_Hello(t *testing.T) {

	tests := []struct {
		name     string
		joke     *api.JokeResponse
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "Simple test",
			joke:     &api.JokeResponse{Joke: "Test joke..."},
			codeWant: 200,
			bodyWant: "Test joke...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := &mocks.Client{}
			apiMock.On("GetJoke").Return(tt.joke, tt.err)

			h := NewHandler(apiMock)
			req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
			rec := httptest.NewRecorder()

			h.Hello(rec, req)
			// handler := http.HandlerFunc(h.Hello)
			// handler.ServeHTTP(rec, req)

			gotRaw, _ := ioutil.ReadAll(rec.Body)
			got := string(gotRaw)

			if got != tt.bodyWant {
				t.Errorf("Wrong response body %s want %s", got, tt.bodyWant)
			}
			if statusGot := rec.Result().StatusCode; statusGot != tt.codeWant {
				t.Errorf("Wrong response status code %v want %v", statusGot, tt.codeWant)
			}

		})
	}
}
