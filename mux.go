package bench

import (
	"net/http"
)

func muxHandleWrite(w http.ResponseWriter, r *http.Request) {}

type tmux struct{}

func (tmux) Build(routes []Route) http.Handler {
	mux := http.ServeMux{}
	for _, route := range routes {
		if !route.MultiplePath {
			mux.HandleFunc(route.Path, muxHandleWrite)
		}
	}
	return &mux
}

func (tmux) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := http.ServeMux{}
	mux.HandleFunc(path, handler)
	return &mux
}
