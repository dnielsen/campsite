package service

import (
	"errors"
	"github.com/google/uuid"
)

type MockAPI struct {
	MockGetEventById func(id uuid.UUID) (*Event, error)
}

func (api *MockAPI) GetEventById(id uuid.UUID) (*Event, error) {
	if api.MockGetEventById != nil {
		return api.MockGetEventById(id)
	}
	return nil, errors.New("something went wrong")
}
