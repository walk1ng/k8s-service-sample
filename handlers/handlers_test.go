package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	router := Router("", "", "")
	ts := httptest.NewServer(router)
	defer ts.Close()

	t.Run("GET /home", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/home")
		if err != nil {
			t.Fatal(err)
		}
		got := res.StatusCode
		want := http.StatusOK
		checkStatusCode(t, got, want)
	})

	t.Run("POST /home", func(t *testing.T) {
		res, err := http.Post(ts.URL+"/home", "text/plain", nil)
		if err != nil {
			t.Fatal(err)
		}
		got := res.StatusCode
		want := http.StatusMethodNotAllowed
		checkStatusCode(t, got, want)
	})

	t.Run("GET /not-exist", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/not-exist")
		if err != nil {
			t.Fatal(err)
		}
		got := res.StatusCode
		want := http.StatusNotFound
		checkStatusCode(t, got, want)
	})

}

func checkStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("status code is wrong. want %d but got %d\n", got, want)
	}
}
