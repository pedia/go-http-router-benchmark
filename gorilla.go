package bench

import (
	"net/http"

	"github.com/gorilla/mux"
)

func gorillaHandleWrite(w http.ResponseWriter, r *http.Request) {}

type tgorilla struct{}

func (tgorilla) Build(routes []Route) http.Handler {
	mux := mux.NewRouter()
	for _, route := range routes {
		if !route.MultiplePath {
			mux.HandleFunc(route.Path, muxHandleWrite)
		}
	}
	return mux
}

func (tgorilla) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	return nil
}
