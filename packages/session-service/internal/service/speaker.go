package service

type Speaker struct {
	Name string `json:"name"`
}

type SpeakerDatastore interface {
	GetAllSpeakers() ([]*Speaker, error)
}

func (api *api) GetAllSpeakers() ([]*Speaker, error) {
	panic("implement me")
}