package bench

import (
	"net/http"

	"github.com/naoina/denco"
)

func dencoHandleWrite(w http.ResponseWriter, r *http.Request, params denco.Params) {}

type tdenco struct{}

func (tdenco) Build(routes []Route) http.Handler {
	hs := []denco.Handler{}
	for _, route := range routes {
		hs = append(hs, denco.Handler{
			Method: route.Method,
			Path:   route.Path,
			Func:   dencoHandleWrite,
		})
	}
	mux := denco.NewMux()
	h, _ := mux.Build(hs)
	return h
}

func (tdenco) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	return nil
}
