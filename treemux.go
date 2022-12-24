package bench

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
)

func treemuxWriter(http.ResponseWriter, *http.Request, map[string]string) {}

type treemux struct{}

func (treemux) Build(routes []Route) http.Handler {
	router := httptreemux.New()
	group := router.NewGroup("/")

	for _, route := range routes {
		group.Handle(route.Method, route.Path, treemuxWriter)
	}
	return router
}

func (treemux) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	return nil
}

func treemuxcWriter(http.ResponseWriter, *http.Request) {}

type treectxmux struct{}

func (treectxmux) Build(routes []Route) http.Handler {
	router := httptreemux.NewContextMux()
	group := router.NewGroup("/")

	for _, route := range routes {
		group.Handle(route.Method, route.Path, treemuxcWriter)
	}
	return router
}

func (treectxmux) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	return nil
}
