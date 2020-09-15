package service

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
)

type MockAPI struct {
	MockGetEventById func(id uuid.UUID) (*Event, error)
}

// We leave the mock function implementation to the test.
// By default it's gonna return an error.
func (api *MockAPI) GetEventById(id uuid.UUID) (*Event, error) {
	if api.MockGetEventById != nil {
		return api.MockGetEventById(id)
	}
	return nil, errors.New("something went wrong")
}

type mockHttpClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

// We leave the mock function implementation to the test.
// By default it's gonna return an error.
func (c mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if c.MockDo != nil {
		return c.MockDo(req)
	}
	return nil, errors.New("something went wrong")
}
