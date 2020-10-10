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

func (api *API) GetCommentsBySessionId(sessionId string, limit int, cursor string) (*[]Comment, *string, error) {
	var commentsPlusOne []Comment

	tx := api.db.Where("session_id = ?", sessionId).Order("created_at DESC").Limit(limit+1)
	if cursor != "" {
		tx = tx.Where("id >= ?", cursor)
	}

	if err := tx.Find(&commentsPlusOne).Error; err != nil {
		return nil, nil, err
	}

	comments := commentsPlusOne[:len(commentsPlusOne)-1]

	var endCursor *string = nil
	if len(comments) == limit && len(commentsPlusOne) > 1 {
		bonusComment := commentsPlusOne[len(commentsPlusOne)-1]
		endCursor = &bonusComment.ID
	}

	return &comments, endCursor, nil
}