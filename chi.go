package bench

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
)

func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(chi.URLParam(r, "id")))
}

type tchi struct{}

func (tchi) Build(routes []Route) http.Handler {
	re := regexp.MustCompile(":([^/]*)")

	h := chiHandleWrite

	mux := chi.NewRouter()
	for _, route := range routes {
		// replace to /path/{a}
		path := re.ReplaceAllString(route.Path, "{$1}")

		switch route.Method {
		case "GET":
			mux.Get(path, h)
		case "POST":
			mux.Post(path, h)
		case "PUT":
			mux.Put(path, h)
		case "PATCH":
			mux.Patch(path, h)
		case "DELETE":
			mux.Delete(path, h)
		default:
			panic("Unknown HTTP method: " + route.Method)
		}
	}
	return mux
}

func (tchi) BuildSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := chi.NewRouter()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}
