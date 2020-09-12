package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"dave-web-app/packages/speaker-service/internal/testUtil"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSpeakersByEventId(t *testing.T) {
	speakers := []service.Speaker{
		{
			ID:   uuid.New(),
			Name: "Warren Smith",
		},
		{
			ID:   uuid.New(),
			Name: "Thomas Smith",
		},
	}

	jsonSpeakers, err := json.Marshal(speakers)
	if err != nil {
		t.Fatalf("Failed to marshal speakers: %v", err)
	}

	testCases := []struct {
		name                 string
		getSpeakersByEventId func(id uuid.UUID) (*[]service.Speaker, error)
		wantCode             int
		wantBody             string
	}{
		{
			"api returns an error",
			func(id uuid.UUID) (*[]service.Speaker, error) {
				return nil, errors.New("oops")
			},
			http.StatusInternalServerError,
			"oops\n",
		},
		{
			"api returns speakers",
			func(id uuid.UUID) (*[]service.Speaker, error) {
				return &speakers, nil
			},
			http.StatusOK,
			//fmt.Sprintf("%v", jsonSpeakers),
			string(jsonSpeakers),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSpeakersByEventId != nil {
				api.MockGetSpeakersByEventId = tc.getSpeakersByEventId
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/speakers", nil)
			h := GetSpeakersByEventId(api)

			h(res, req)
			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
