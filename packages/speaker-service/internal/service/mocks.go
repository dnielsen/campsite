package service

import (
	"errors"
)

type MockAPI struct {
	MockGetAllSpeakers func() (*[]Speaker, error)
}

func (api *MockAPI) GetAllSpeakers() (*[]Speaker, error) {
	if api.MockGetAllSpeakers != nil {
		return api.MockGetAllSpeakers()
	}
	return nil, errors.New("something went wrong")
}