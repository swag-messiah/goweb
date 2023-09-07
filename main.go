package main

import "net/http"

type MyHandler struct {
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home page"))
		return
	}
	if r.URL.Path == "/hello" {
		w.Write([]byte("Hello, user"))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}
func main() {
	http.ListenAndServe(":3000", MyHandler{})
}
