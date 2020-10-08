package service

import (
	"github.com/google/uuid"
	"time"
)

func (api *API) CreateComment(sessionId string, i CommentInput) (*Comment, error) {
	c := Comment{
		ID:        uuid.New().String(),
		Content:   i.Content,
		CreatedAt: time.Now(),
		SessionID: sessionId,
	}
	if err := api.db.Create(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}