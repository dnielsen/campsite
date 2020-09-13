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

func TestGetAllSpeakers(t *testing.T) {
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
		getAllSpeakers func() (*[]service.Speaker, error)
		wantCode         int
		wantBody         string
	}{
		{
			"api returns an error",
			func() (*[]service.Speaker, error) {
				return nil, errors.New("something went wrong")
			},
			http.StatusInternalServerError,
			"something went wrong\n",
		},
		{
			"speakers found",
			func() (*[]service.Speaker, error) {
				return &speakers, nil
			},
			http.StatusOK,
			func() string {
				jsonSpeakers, _ := json.Marshal(speakers)
				return string(jsonSpeakers)
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getAllSpeakers != nil {
				api.MockGetAllSpeakers = tc.getAllSpeakers
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/speakers", nil)
			// We have to set url vars for unit testing, otherwise gorilla mux won't register
			// our vars, so the id would be an empty string.

			h := GetAllSpeakers(api)
			h(res, req)

			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
