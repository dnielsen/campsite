package service

import (
	"errors"
)

type MockAPI struct {
	MockGetSessionsByIds func(ids []string) (*[]Session, error)
}

func (api *MockAPI) GetSessionsByIds(ids []string) (*[]Session, error) {
	if api.MockGetSessionsByIds != nil {
		return api.MockGetSessionsByIds(ids)
	}
	return nil, errors.New("something went wrong")
}
