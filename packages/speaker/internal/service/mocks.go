package service

import (
	"errors"
)

type MockAPI struct {
	MockGetSpeakersByIds func(ids []string) (*[]Speaker, error)
	MockGetAllSpeakers func() (*[]Speaker, error)
	MockGetSpeakerById func(id string) (*Speaker, error)
}

func (api *MockAPI) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	if api.MockGetSpeakersByIds != nil {
		return api.MockGetSpeakersByIds(ids)
	}
	return nil, errors.New("something went wrong")
}

func (api *MockAPI) GetSpeakerById(id string) (*Speaker, error) {
	if api.MockGetSpeakerById != nil {
		return api.MockGetSpeakerById(id)
	}
	return nil, errors.New("something went wrong")
}

func (api *MockAPI) GetAllSpeakers() (*[]Speaker, error) {
	if api.MockGetAllSpeakers != nil {
		return api.MockGetAllSpeakers()
	}
	return nil, errors.New("something went wrong")
}
