package service

type Speaker struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type SpeakerDatastore interface {
	GetAllSpeakers() (*[]Speaker, error)
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	_ = api.db.Find(&speakers)
	return &speakers, nil
}