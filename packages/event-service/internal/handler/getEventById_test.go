package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"dave-web-app/packages/event-service/internal/testUtil"
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

func TestGetSpeakers(t *testing.T) {
	eventId := uuid.New()
	event := service.Event{
		ID:            eventId,
		Name:          "Great Event",
		StartDate:     time.Now(),
		EndDate:       time.Date(2022, time.November, 10, 23, 0, 0, 0, time.UTC),
		Photo:         "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
		OrganizerName: "John Tim",
		Address:       "San Francisco, California",
	}

	jsonEvent, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal event: %v", err)
	}

	testCases := []struct {
		id           uuid.UUID
		name         string
		getEventById func(id uuid.UUID) (*service.Event, error)
		wantCode     int
		wantBody     string
	}{
		{
			eventId,
			"api returns the event when passing id of an existing event",
			func(id uuid.UUID) (*service.Event, error) {
				if id == eventId {
					return &event, nil
				}
				return nil, errors.New("event not found")
			},
			http.StatusOK,
			string(jsonEvent),
		},
		{
			uuid.New(),
			"api returns an error when event not found",
			// It's the same func as above, however, in Go, a little bit of code duplication is ok.
			func(id uuid.UUID) (*service.Event, error) {
				if id == eventId {
					return &event, nil
				}
				return nil, errors.New("event not found")
			},
			http.StatusBadRequest,
			fmt.Sprintf("%v\n", errors.New("event not found").Error()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getEventById != nil {
				api.MockGetEventById = tc.getEventById
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/events/%v", tc.id), nil)
			// We have to set url vars for unit testing, otherwise gorilla mux won't register
			// our vars, so the email would be an empty string.
			vars := map[string]string{
				ID: tc.id.String(),
			}
			req = mux.SetURLVars(req, vars)

			h := GetEventById(api)
			h(res, req)
			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}
