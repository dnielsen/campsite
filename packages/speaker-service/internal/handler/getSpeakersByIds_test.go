package handler

import (
	"bytes"
	"dave-web-app/packages/speaker-service/internal/service"
	"dave-web-app/packages/speaker-service/internal/testUtil"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSpeakersByIds(t *testing.T) {
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
		ids []string
		getSpeakersByIds func(ids []string) (*[]service.Speaker, error)
		wantCode         int
		wantBody         string
	}{
		{
			"speakers not found",
			[]string{uuid.New().String(), uuid.New().String()},
			func(ids []string) (*[]service.Speaker, error) {
				var foundSpeakers []service.Speaker
				if len(ids) == 1 && (ids[0] == speakers[0].ID || ids[1] == speakers[0].ID) {
					foundSpeakers = append(foundSpeakers, speakers[0])
				}
				if len (ids) == 2 && (ids[0] == speakers[1].ID || ids[1] == speakers[1].ID) {
					foundSpeakers = append(foundSpeakers, speakers[1])
				}

				if len(foundSpeakers) == 0 {
					return nil, errors.New("speakers not found")
				}

				return &foundSpeakers, nil
			},
			http.StatusBadRequest,
			"speakers not found\n",
		},
		{
			"one speaker found",
			[]string{speakers[1].ID},
			func(ids []string) (*[]service.Speaker, error) {
				var foundSpeakers []service.Speaker
				if len(ids) == 1 {
					if ids[0] == speakers[0].ID {
						foundSpeakers = append(foundSpeakers, speakers[0])
					} else if ids[0] == speakers[1].ID {
						foundSpeakers = append(foundSpeakers, speakers[1])
					}
				}

				if len (ids) == 2 {
					if ids[0] == speakers[0].ID || ids[1] == speakers[0].ID {
						foundSpeakers = append(foundSpeakers, speakers[0])
					} else if ids[0] == speakers[1].ID || ids[1] == speakers[1].ID {
						foundSpeakers = append(foundSpeakers, speakers[1])
					}
				}

				if len(foundSpeakers) == 0 {
					return nil, errors.New("speakers not found")
				}
				return &foundSpeakers, nil
			},
			http.StatusOK,
			func() string {
				jsonSpeakers, _ := json.Marshal([]service.Speaker{speakers[1]})
				return string(jsonSpeakers)
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSpeakersByIds != nil {
				api.MockGetSpeakersByIds = tc.getSpeakersByIds
			}
			res := httptest.NewRecorder()
			b, err := json.Marshal(getSpeakersByIdsRequestBody{SpeakerIds: tc.ids})
			if err != nil {
				t.Fatalf("Failed to marshal body: %v", err)
			}
			req := httptest.NewRequest(http.MethodGet, "/speakers", bytes.NewBuffer(b))
			h := GetSpeakersByIds(api)
			h(res, req)

			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
