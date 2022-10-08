package apiserver

import (
	"bytes"
	"context"
	"encoding/json"
	"example.com/prj/model"
	"example.com/prj/store/test"
	"fmt"
	"github.com/gorilla/securecookie"
	sessions2 "github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(test.NewStore(), sessions2.NewCookieStore([]byte("some-session-key")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.org",
				"pwd":   "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "not-a-json",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "no password",
			payload: map[string]string{
				"email": "user@example.org",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "not an email",
			payload: map[string]string{
				"email": "amongus",
				"pwd":   "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
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
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	ctx := context.Background()
	u := model.TestUser(t)
	store := test.NewStore()
	store.User().Create(ctx, u)
	s := NewServer(store, sessions2.NewCookieStore([]byte("some-session-key")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": u.Email,
				"pwd":   u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "different email",
			payload: map[string]string{
				"email": "prowler@swag.inc",
				"pwd":   u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "different password",
			payload: map[string]string{
				"email": u.Email,
				"pwd":   "amongus",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "invalid payload",
			payload:      "not-a-json",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_AuthenticateUser(t *testing.T) {

	store := test.NewStore()
	u := model.TestUser(t)
	store.User().Create(context.Background(), u)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}
	secretKey := []byte("secret")
	s := NewServer(store, sessions2.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.AuthenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
