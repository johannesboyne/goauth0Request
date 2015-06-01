package goauth0Request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDo(t *testing.T) {
	// auth0 fake server
	auth0_ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var a Auth0Prepost
		_ = decoder.Decode(&a)
		if a.Client_id != "1234567890" {
			t.Errorf("Wrong client id transmitted: %v", a.Client_id)
		}
		fmt.Fprintln(w, `{"id_token": "9876543210"}`)
	}))
	defer auth0_ts.Close()

	// get fake server
	get_ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer 9876543210" {
			t.Errorf("Wrong Auth.-Token: %s", r.Header.Get("Authorization"))
		}
		fmt.Fprintln(w, `return 2`)
	}))
	defer auth0_ts.Close()

	Get{
		auth0_ts.URL,
		"1234567890",
		"username",
		"password",
		"",
		"",
		"",
		get_ts.URL,
	}.Do()
}
