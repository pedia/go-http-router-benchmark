package bench

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func httprouterHandleWrite(http.ResponseWriter, *http.Request, httprouter.Params) {}

type thttprouter struct{}

func (thttprouter) Build(routes []Route) http.Handler {
	mux := httprouter.New()
	for _, route := range routes {
		if !route.MultiplePath {
			mux.Handle(route.Method, route.Path, httprouterHandleWrite)
		}
	}
	return mux
}

func (thttprouter) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	return nil
}
