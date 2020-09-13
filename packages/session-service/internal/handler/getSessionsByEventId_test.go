package handler

import (
	"dave-web-app/packages/session-service/internal/service"
	"dave-web-app/packages/session-service/internal/testUtil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetSessionsByEventId(t *testing.T) {
	eventIds := []string{uuid.New().String(), uuid.New().String()}

	sessions := []service.Session{
		{
			ID:          uuid.New().String(),
			Name:        "Test Session",
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(5000),
			Description: "Session description",
			SpeakerIds:  []string{uuid.New().String(), uuid.New().String()},
			EventId:     eventIds[0],
		},
		{
			ID:          uuid.New().String(),
			Name:        "Test Session 2",
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(4000),
			Description: "Session description 2",
			SpeakerIds:  []string{uuid.New().String()},
			EventId:     eventIds[1],
		},
		{
			ID:          uuid.New().String(),
			Name:        "Test Session 2",
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(4000),
			Description: "Session description 2",
			SpeakerIds:  []string{uuid.New().String()},
			EventId:     eventIds[1],
		},
	}

	testCases := []struct {
		name           string
		eventId string
		getSessionsByEventId func(id string) (*[]service.Session, error)
		wantCode       int
		wantBody       string
	}{
		{
			"event id doesn't exist",
			uuid.New().String(),
			func(id string) (*[]service.Session, error) {
				var foundSessions []service.Session
				if id == sessions[0].EventId {
					foundSessions = append(foundSessions, sessions[0])
				}
				if id == sessions[1].EventId {
					foundSessions = append(foundSessions, sessions[1])
				}
				if id == sessions[2].EventId {
					foundSessions = append(foundSessions, sessions[2])
				}

				if len(foundSessions) == 0 {
					return nil, errors.New("sessions with the given event id not found")
				}

				return &foundSessions, nil
			},
			http.StatusBadRequest,
			"sessions with the given event id not found\n",
		},
		{
			"api returns the sessions with the specified event id",
			eventIds[1],
			// It's the same func as above, however, in Go,
			// a little bit of code duplication is ok (common).
			func(id string) (*[]service.Session, error) {
				var foundSessions []service.Session
				if id == sessions[0].EventId {
					foundSessions = append(foundSessions, sessions[0])
				}
				if id == sessions[1].EventId {
					foundSessions = append(foundSessions, sessions[1])
				}
				if id == sessions[2].EventId {
					foundSessions = append(foundSessions, sessions[2])
				}

				if len(foundSessions) == 0 {
					return nil, errors.New("sessions with the given event id not found")
				}

				return &foundSessions, nil
			},
			http.StatusOK,
			func () string {
				wantSessions := []service.Session{sessions[1], sessions[2]}
				jsonSession, _ := json.Marshal(wantSessions)
				return string(jsonSession)
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSessionsByEventId != nil {
				api.MockGetSessionsByEventId = tc.getSessionsByEventId
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/event/%v", tc.eventId), nil)
			// We have to set url vars for unit testing, otherwise gorilla mux won't register
			// our vars, so the id would be an empty string.
			vars := map[string]string{
				EVENT_ID: tc.eventId,
			}
			req = mux.SetURLVars(req, vars)

			h := GetSessionsByEventId(api)
			h(res, req)

			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
