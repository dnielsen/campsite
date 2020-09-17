package service

type Speaker struct {
	ID         string    `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name       string    `json:"name,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Headline   string    `json:"headline,omitempty"`
	Photo      string    `json:"photo,omitempty"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;"`
}