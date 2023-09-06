package main

import "net/http"

type MyHandler struct {
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, сервер! Привет Мир!"))
}
func main() {
	http.ListenAndServe(":3000", MyHandler{})
}
