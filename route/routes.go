package route

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/NYTimes/gziphandler"
	"github.com/go-chi/chi"
)

// APIServerHandler returns a Gzip handler
func APIServerHandler() http.Handler {
	r := newAPIRouter()
	return gziphandler.GzipHandler(r)
}

func newAPIRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "application is healthy")
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return r
}
