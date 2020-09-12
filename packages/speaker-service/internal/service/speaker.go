package service


type Speaker struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Name string    `json:"name"`
	Bio string `json:"bio"`
	Headline string `json:"headline"`
	Photo string `json:"photo"`
}

type SpeakerDatastore interface {
	GetSpeakersByIds(ids []string) (*[]Speaker, error)
}

func (api *api) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	var speakers []Speaker
	_ = api.db.Find(&speakers, ids)
	return &speakers, nil
}
