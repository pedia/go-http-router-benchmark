package bench

import (
	"net/http"
)

type Route struct {
	Method       string
	Path         string
	MultiplePath bool
}

type Benchable interface {
	Build(routes []Route) http.Handler
	BuildSingle(method, path string, handler http.HandlerFunc) http.Handler
}

type DummyWriter struct{}

func (m *DummyWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *DummyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *DummyWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *DummyWriter) WriteHeader(int) {}
