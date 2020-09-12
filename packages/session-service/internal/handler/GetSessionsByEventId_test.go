package handler

import (
	"bytes"
	"dave-web-app/packages/session-service/internal/service"
	"dave-web-app/packages/session-service/internal/testUtil"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetSessionsByIds(t *testing.T) {
	ids := []string{uuid.New().String(), uuid.New().String()}
	sessions := []service.Session{
		{
			ID:          uuid.New().String(),
			Name:        "dddd Name",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(1, 1, 1),
			Description: "ddd of the session",
		},
		{
			ID:          uuid.New().String(),
			Name:        "dd Name",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(1, 1, 1),
			Description: "desddcription of the session",
		},
	}

	jsonSessions, err := json.Marshal(sessions)
	if err != nil {
		t.Fatalf("Failed to marshal sessions: %v", err)
	}

	testCases := []struct {
		name             string
		ids []string
		getSessionsByIds func(ids []string) (*[]service.Session, error)
		wantCode         int
		wantBody         string
	}{
		{
			"api returns an error",
			ids,
			func(ids []string) (*[]service.Session, error) {
				return nil, errors.New("oops")
			},
			http.StatusBadRequest,
			"oops\n",
		},
		{
			"api returns all sessions given all session ids are in request body",
			ids,
			func(ids []string) (*[]service.Session, error) {
				return &sessions, nil
			},
			http.StatusOK,
			string(jsonSessions),
		},
		{
			"api returns one session",
			[]string{ids[1]},
			func(ids []string) (*[]service.Session, error) {
				wantIds := []string{ids[1]}

				if reflect.DeepEqual(ids, wantIds) {
					return &[]service.Session{sessions[1]}, nil
				}
				// TODO
				return &sessions, nil
			},
			http.StatusOK,
			// TODO
			string(jsonSessions),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSessionsByIds != nil {
				api.MockGetSessionsByIds = tc.getSessionsByIds
			}
			res := httptest.NewRecorder()
			b, err := json.Marshal(getSessionsByIdsRequestBody{SessionIds: ids})
			if err != nil {
				t.Fatalf("Failed to marshal body: %v", err)
			}
			req := httptest.NewRequest("GET", "/speakers", bytes.NewBuffer(b))
			h := GetSessionsByEventId(api)
			h(res, req)
			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
