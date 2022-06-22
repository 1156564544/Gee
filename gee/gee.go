package gee

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	route map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{route: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, pattern string, hfn HandlerFunc) {
	e.route[genKey(method, pattern)] = hfn
}

func (e *Engine) Get(pattern string, hfn HandlerFunc) {
	e.addRoute("GET", pattern, hfn)
}

func (e *Engine) Post(pattern string, hfn HandlerFunc) {
	e.addRoute("POST", pattern, hfn)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := genKey(r.Method, r.URL.Path)
	if fn, ok := e.route[key]; ok {
		fn(w, r)
	} else {
		w.WriteHeader(404)
	}
}

func genKey(method, url string) string {
	return method + "_" + url
}
