package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"dave-web-app/packages/speaker-service/internal/testUtil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSpeakerById(t *testing.T) {
	ids := []string{uuid.New().String(), uuid.New().String()}
	speakers := []service.Speaker{
		{
			ID:       ids[0],
			Name:     "John Doe",
			Bio:      "Hello world",
			Headline: "CEO of Hello",
			Photo:    "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
		},
		{
			ID:       ids[1],
			Name:     "Tom Unsplash",
			Bio:      "Hello Earth",
			Headline: "CEO of World",
			Photo:    "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
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
