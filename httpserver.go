package main

import "net/http"

func main() {
	serv := http.NewServeMux()
	serv.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
		serv.UnhandleFunc("/hello")
	})
	http.ListenAndServe(":8080", serv)
}
