package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //What does w mean?
		w.Write([]byte("Hello, World!"))
		//Why Write uses byte slice instead of string?
	})
	http.ListenAndServe(":8080", nil)
}
