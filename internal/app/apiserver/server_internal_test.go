package apiserver

import (
	"bytes"
	"encoding/json"
	"example.com/prj/store/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(test.NewStore())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.org",
				"pwd":   "pwd",
			},
			expectedCode: http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

	//rec := httptest.NewRecorder()
	//req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	//
	//s.ServeHTTP(rec, req)
	//assert.Equal(t, rec.Code, http.StatusOK)
}
