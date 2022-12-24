package bench

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var targets = []struct {
	Name   string
	Target Benchable
}{
	{"Chi", tchi{}},
	{"ServeMux", tmux{}},
	{"Denco", tdenco{}},
	{"httprouter", thttprouter{}},
	{"Gorilla", tgorilla{}},
	{"treemux", treemux{}},
	{"treectxmux", treectxmux{}},
}

func TestHandle(t *testing.T) {
	for _, target := range targets {
		h := target.Target.Build(githubAPI)

		r, _ := http.NewRequest("GET", "/authorizations", nil)
		w := httptest.NewRecorder()

		h.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Fatalf("/ not match %d", w.Code)
		}

		if w.Body.String() != "" {
			t.Fatalf("/ not match %s", w.Body.String())
		}

		// r2, _ := http.NewRequest("POST", "/authorizations", nil)
		// w2 := httptest.NewRecorder()

		// h.ServeHTTP(w2, r2)
		// if w2.Code != 200 {
		// 	t.Fatalf("/ not match %d", w2.Code)
		// }

		// if w2.Body.String() != "" {
		// 	t.Fatalf("/ not match %s", w2.Body.String())
		// }
	}
}

func BenchmarkHandle(b *testing.B) {
	r404, _ := http.NewRequest("GET", "/", nil)
	rp, _ := http.NewRequest("POST", "/authorizations", nil)
	rg, _ := http.NewRequest("GET", "/authorizations", nil)
	rg2, _ := http.NewRequest("GET", "/applications/client_id/tokens/access_token", nil)
	w := &DummyWriter{}

	for _, t := range targets {
		h := t.Target.Build(githubAPI)
		b.Run(fmt.Sprintf("%s404", t.Name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				h.ServeHTTP(w, r404)
			}
		})
	}
	for _, t := range targets {
		h := t.Target.Build(githubAPI)
		b.Run(fmt.Sprintf("%sPost", t.Name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				h.ServeHTTP(w, rp)
			}
		})
	}
	for _, t := range targets {
		h := t.Target.Build(githubAPI)
		b.Run(fmt.Sprintf("%sGet", t.Name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				h.ServeHTTP(w, rg)
			}
		})
	}
	for _, t := range targets {
		h := t.Target.Build(githubAPI)
		b.Run(fmt.Sprintf("%sGet2Arg", t.Name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				h.ServeHTTP(w, rg2)
			}
		})
	}
}
