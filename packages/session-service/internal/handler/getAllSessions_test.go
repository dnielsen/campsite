package handler

import (
	"dave-web-app/packages/session-service/internal/service"
	"dave-web-app/packages/session-service/internal/testUtil"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAllSessions(t *testing.T) {
	sessions := []service.Session{{
		ID:          uuid.New().String(),
		Name:        "Test Session",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(5000),
		Description: "Session description",
		SpeakerIds:  []string{uuid.New().String(), uuid.New().String()},
		EventId:     uuid.New().String(),
	},
		{
			ID:          uuid.New().String(),
			Name:        "Test Session 2",
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(4000),
			Description: "Session description 2",
			SpeakerIds:  []string{uuid.New().String()},
			EventId:     uuid.New().String(),
		}}

	jsonSessions, err := json.Marshal(sessions)
	if err != nil {
		t.Fatalf("Failed to marshal event: %v", err)
	}

	testCases := []struct {
		name           string
		getAllSessions func() (*[]service.Session, error)
		wantCode       int
		wantBody       string
	}{
		{
			"api returns an error",
			func() (*[]service.Session, error) {
				return nil, errors.New("something went wrong")
			},
			http.StatusInternalServerError,
			"something went wrong\n",
		},
		{
			"api returns all sessions",
			// It's the same func as above, however, in Go, a little bit of code duplication is ok.
			func() (*[]service.Session, error) {
				return &sessions, nil
			},
			http.StatusOK,
			string(jsonSessions),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getAllSessions != nil {
				api.MockGetAllSessions = tc.getAllSessions
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/sessions", nil)

			h := GetAllSessions(api)
			h(res, req)

			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
