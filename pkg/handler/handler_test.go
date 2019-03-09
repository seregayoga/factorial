package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	testCases := []struct {
		url        string
		body       string
		statusCode int
	}{
		{
			url:        "http://localhost/calculate?n=12",
			body:       "479001600",
			statusCode: http.StatusOK,
		},
		{
			url:        "http://localhost/calculate?n=not-a-number",
			body:       "Parameter \"n\" should be a positive number\n",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest("GET", testCase.url, nil)
		w := httptest.NewRecorder()
		CalculateHandler(w, req)

		resp := w.Result()

		if resp.StatusCode != testCase.statusCode {
			t.Errorf("Expected status code %d, got %d for url %s", testCase.statusCode, resp.StatusCode, testCase.url)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		if string(body) != testCase.body {
			t.Errorf("Expected body %s, got %s for url %s", testCase.body, string(body), testCase.url)
		}
	}
}

func TestGetNumber(t *testing.T) {
	testCases := []struct {
		url    string
		number int64
		err    error
	}{
		{
			url:    "/calculate?n=123",
			number: 123,
			err:    nil,
		},
		{
			url:    "/calculate?n=-123",
			number: 0,
			err:    errors.New("Parameter \"n\" should be a positive number"),
		},
		{
			url:    "/calculate?n=",
			number: 0,
			err:    errors.New("Missing \"n\" parameter in query"),
		},
	}

	for _, testCase := range testCases {
		req, err := http.NewRequest(http.MethodGet, testCase.url, nil)
		if err != nil {
			t.Error(err)
		}

		number, err := getNumber(req)

		if number != testCase.number {
			t.Errorf("Wrong number from url %s. Got %d, expected %d", testCase.url, number, testCase.number)
		}
		if testCase.err == nil && err != nil {
			t.Errorf("Unexpected error %s from url %s.", err, testCase.url)
		} else if testCase.err != nil && testCase.err.Error() != err.Error() {
			t.Errorf("Wrong error from url %s. Got %s, expected %s.", testCase.url, err, testCase.err)
		}
	}
}
