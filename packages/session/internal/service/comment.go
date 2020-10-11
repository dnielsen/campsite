package service

import (
	"encoding/json"
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
		tx = tx.Where("created_at <= ?", cursor)
	}

	if err := tx.Find(&commentsPlusOne).Error; err != nil {
		return nil, nil, err
	}

	if len(commentsPlusOne) == 0 {
		return &commentsPlusOne, nil, nil
	}

	comments := commentsPlusOne[:len(commentsPlusOne)-1]

	if len(comments) != limit {
		return &commentsPlusOne, nil, nil
	}

	// `endCursor` is the `createdAt` date of the next comment
	var endCursor *string = nil
	bonusComment := commentsPlusOne[len(commentsPlusOne)-1]
	// Convert Go's `time.Time` into JSON Date. Otherwise our `endCursor` would be of type `time.Time`.
	endCursorBytes, _ := json.Marshal(bonusComment.CreatedAt)
	endCursorString := string(endCursorBytes)
	endCursor = &endCursorString
	return &comments, endCursor, nil
}