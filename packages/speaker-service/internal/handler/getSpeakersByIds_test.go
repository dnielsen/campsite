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
	"reflect"
	"testing"
)

func TestGetSpeakersByEventId(t *testing.T) {
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

	jsonSpeakers, err := json.Marshal(speakers)
	if err != nil {
		t.Fatalf("Failed to marshal speakers: %v", err)
	}

	testCases := []struct {
		name             string
		ids []string
		getSpeakersByIds func(ids []string) (*[]service.Speaker, error)
		wantCode         int
		wantBody         string
	}{
		{
			"api returns an error",
			ids,
			func(ids []string) (*[]service.Speaker, error) {
				return nil, errors.New("oops")
			},
			http.StatusInternalServerError,
			"oops\n",
		},
		{
			"api returns all speakers given all speaker ids are in request body",
			ids,
			func(ids []string) (*[]service.Speaker, error) {
				return &speakers, nil
			},
			http.StatusOK,
			string(jsonSpeakers),
		},
		{
			"api returns one speaker",
			[]string{ids[1]},
			func(ids []string) (*[]service.Speaker, error) {
				wantIds := []string{ids[1]}

				if reflect.DeepEqual(ids, wantIds) {
					return &[]service.Speaker{speakers[1]}, nil
				}
				return &speakers, nil
			},
			http.StatusOK,
			string(jsonSpeakers),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getSpeakersByIds != nil {
				api.MockGetSpeakersByIds = tc.getSpeakersByIds
			}
			res := httptest.NewRecorder()
			b, err := json.Marshal(getSpeakersByIdsRequestBody{SpeakerIds: ids})
			if err != nil {
				t.Fatalf("Failed to marshal body: %v", err)
			}
			req := httptest.NewRequest("GET", "/speakers", bytes.NewBuffer(b))
			h := GetSpeakersByIds(api)
			h(res, req)
			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
