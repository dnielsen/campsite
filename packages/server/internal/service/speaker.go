package service

type Speaker struct {
	ID         string    `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name       string    `json:"name,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Headline   string    `json:"headline,omitempty"`
	Photo      string    `json:"photo,omitempty"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;"`
}

func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	var speaker Speaker
	_ = api.db.Preload("Sessions").Where("id = ?", id).First(&speaker)
	return &speaker, nil
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	_ = api.db.Find(&speakers)
	return &speakers, nil
}

