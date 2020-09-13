package handler

import (
	"dave-web-app/packages/session-service/internal/service"
	"dave-web-app/packages/session-service/internal/testUtil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetSessionById(t *testing.T) {
	sessions := []service.Session{
		{
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
		},
	}
	testCases := []struct {
		name           string
		id string
		getSessionById func(id string) (*service.Session, error)
		wantCode       int
		wantBody       string
	}{
		{
			"session not found",
			uuid.New().String(),
			func(id string) (*service.Session, error) {
				if id == sessions[0].ID {
					return &sessions[0], nil
				}
				if id == sessions[1].ID {
					return &sessions[1], nil
				}
				return nil, errors.New("session not found")
			},
			http.StatusBadRequest,
			"session not found\n",
		},
		{
			"api returns the specified session",
			sessions[0].ID,
			// It's the same func as above, however, in Go,
			// a little bit of code duplication is ok (common).
			func(id string) (*service.Session, error) {
				log.Println(id, sessions[0].ID)
				if id == sessions[0].ID {
					return &sessions[0], nil
				}
				if id == sessions[1].ID {
					return &sessions[1], nil
				}
				return nil, errors.New("session not found")
			},
			http.StatusOK,
			func () string {
				jsonSession, _ := json.Marshal(sessions[0])
				return string(jsonSession)
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSessionById != nil {
				api.MockGetSessionById = tc.getSessionById
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/%v", tc.id), nil)
			// We have to set url vars for unit testing, otherwise gorilla mux won't register
			// our vars, so the id would be an empty string.
			vars := map[string]string{
				ID: tc.id,
			}
			req = mux.SetURLVars(req, vars)

			h := GetSessionById(api)
			h(res, req)

			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
