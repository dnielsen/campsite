package service

import (
	"errors"
)

type MockAPI struct {
	MockGetSpeakersByIds func(ids []string) (*[]Speaker, error)
}

func (api *MockAPI) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	if api.MockGetSpeakersByIds != nil {
		return api.MockGetSpeakersByIds(ids)
	}
	return nil, errors.New("something went wrong")
}
