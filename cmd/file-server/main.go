package main

import (
	"log"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir(`D:\final\Projects\flutter_comic\build\web`))
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, OPTIONS")

		h.ServeHTTP(rw, r)
	})))
}
