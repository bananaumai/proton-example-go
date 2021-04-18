package main

import "net/http"

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	_ = http.ListenAndServe("0.0.0.0:3000", nil)
}
