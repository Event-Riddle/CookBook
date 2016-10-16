package server

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/api/v1/config/(filter|user)")

type InvalidConfigType struct{ error }

type MiddlewareHandler func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

type Middleware struct {
	Middles []http.HandlerFunc
}

func NewMiddleware(router http.Handler) *Middleware {
	vw := &Middleware{}
	vw.Middles = []http.HandlerFunc{router.ServeHTTP}
	return vw
}

func (self *Middleware) Use(middle MiddlewareHandler) {
	m := self.Middles[len(self.Middles)-1]
	sh := func(w http.ResponseWriter, r *http.Request) {
		middle(w, r, m)
	}
	self.Middles = append(self.Middles, sh)
}

func (self *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.Middles[len(self.Middles)-1].ServeHTTP(w, r)
}

func (self *Middleware) PreProcessing() {
	self.Use(validatePath)
	self.Use(cors)
}

func validatePath(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	url := r.URL.Path
	m := validPath.FindStringSubmatch(url)
	if m == nil {
		err := InvalidConfigType{errors.New("Invalid Path")}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	next(w, r)
}

func cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Request-Method", "GET, POST")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")

	next(w, r)
}

func middlewareErrorHandler(f func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) error) MiddlewareHandler {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		err := f(w, r, next)
		if err == nil {
			return
		}
		switch err.(type) {
		case InvalidConfigType:
			fmt.Println("Not goooodd")
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			fmt.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

/* USAGE
muxy := mux.NewRouter()
muxy.HandleFunc().Get()
m := NewValidatorWare(muxy)
m.Use(validate)
return m
*/
