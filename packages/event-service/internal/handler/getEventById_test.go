package handler

import (
	"dave-web-app/packages/event-service/internal/service"
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

func TestGetEventById(t *testing.T) {
	ids := []string{uuid.New().String(), uuid.New().String()}
	events := []service.Event{
		{
			ID:            "",
			Name:          "",
			Description:   "",
			StartDate:     time.Time{},
			EndDate:       time.Time{},
			Photo:         "",
			OrganizerName: "",
			Address:       "",
			Sessions:      nil,
			SpeakerIds:    nil,
			Speakers:      nil,
		},
	}

	testCases := []struct {
		name             string
		id string
		getSpeakerById func(id string) (*service.Speaker, error)
		wantCode         int
		wantBody         string
	}{
		{
			"speaker found",
			uuid.New().String(),
			func(id string) (*service.Speaker, error) {
				if id == speakers[0].ID {
					return &speakers[0], nil
				}
				if id == speakers[1].ID {
					return &speakers[1], nil
				}

				return nil, errors.New("speaker not found")
			},
			http.StatusBadRequest,
			"speaker not found\n",
		},
		{
			"speaker not found",
			ids[1],
			func(id string) (*service.Speaker, error) {
				if id == speakers[0].ID {
					return &speakers[0], nil
				}
				if id == speakers[1].ID {
					return &speakers[1], nil
				}

				return nil, errors.New("speaker not found")
			},
			http.StatusOK,
			func() string {
				jsonSpeaker, _ := json.Marshal(speakers[1])
				return string(jsonSpeaker)
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSpeakerById != nil {
				api.MockGetSpeakerById = tc.getSpeakerById
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/speakers/%v",tc.id), nil)
			// We have to set url vars for unit testing, otherwise gorilla mux won't register
			// our vars, so the id would be an empty string.
			vars := map[string]string{
				ID: tc.id,
			}
			req = mux.SetURLVars(req, vars)

			h := GetSpeakerById(api)
			h(res, req)

			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
