package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"dave-web-app/packages/speaker-service/testUtil"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSpeakers(t *testing.T) {
	testCases := []struct {
		name               string
		getAllSpeakers 	   func() (*[]service.Speaker, error)
		wantCode           int
		wantBody           string
	}{
		// API can't return an error, so there's no test for that
		{
			"api returns an error",
			func() (*[]service.Speaker, error) {
				return nil, errors.New("oops")
			},
			http.StatusInternalServerError,
			"oops\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api := &service.MockAPI{}
			if tc.getAllSpeakers != nil {
				api.MockGetAllSpeakers = tc.getAllSpeakers
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/speakers", nil)
			h := GetSpeakers(api)

			h(res, req)
			gotCode := res.Code
			gotBody := res.Body.String()

			testUtil.Cmp(t, tc.wantCode, gotCode)
			testUtil.Cmp(t, tc.wantBody, gotBody)
		})
	}
}