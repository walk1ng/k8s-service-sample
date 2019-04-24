package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

// Router will create router obj
func Router(buildTime, commit, release string) *mux.Router {

	isReady := &atomic.Value{}
	isReady.Store(false)
	r := mux.NewRouter()

	go func() {
		log.Println("readyz probe is negative by default...")
		// mock time cost by ready everything
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Println("readyz probe is positive.")
	}()

	r.HandleFunc("/home", home(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
