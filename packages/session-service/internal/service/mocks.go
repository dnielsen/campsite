package service

import (
	"errors"
)

type MockAPI struct {
	MockGetAllSessions func() (*[]Session, error)
	MockGetSessionsByEventId func(id string) (*[]Session, error)
	MockGetSessionById func(id string) (*Session, error)
}

func (api *MockAPI) GetSessionsByEventId(id string) (*[]Session, error) {
	if api.MockGetSessionsByEventId != nil {
		return api.MockGetSessionsByEventId(id)
	}
	return nil, errors.New("something went wrong")
}

func (api *MockAPI) GetSessionById(id string) (*Session, error) {
	if api.MockGetSessionById != nil {
		return api.MockGetSessionById(id)
	}
	return nil, errors.New("something went wrong")
}

func (api *MockAPI) GetAllSessions() (*[]Session, error) {
	if api.MockGetAllSessions != nil {
		return api.MockGetAllSessions()
	}
	return nil, errors.New("something went wrong")
}

