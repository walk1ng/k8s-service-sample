package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	buildTime := time.Now().String()
	commit := "test hash"
	release := "0.0.1"

	h := home(buildTime, commit, release)
	h(w, nil)

	resp := w.Result()
	got := resp.StatusCode
	want := http.StatusOK
	checkStatusCode(t, got, want)

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{}

	err = json.Unmarshal(greeting, &info)
	if err != nil {
		t.Fatal(err)
	}

	if info.Release != release {
		t.Errorf("Release version is wrong. got %q but want %q\n", info.Release, release)
	}

	if info.Commit != commit {
		t.Errorf("Commit is wrong. got %q but want %q\n", info.Commit, commit)
	}
	if info.BuildTime != buildTime {
		t.Errorf("Build time is wrong. got %q but want %q\n", info.BuildTime, buildTime)
	}
}
