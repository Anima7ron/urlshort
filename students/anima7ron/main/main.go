package main

import "net/http"

func main() {
	mux := defaultMux()
	yaml := `
         - path: /urlshort-godoc
           url: https://godoc.org/github.com/gophercises/urlshort
         - path: /yaml-godoc
           url: https://godoc.org/gopkg.in/yaml.v2
         - path: /urlshort
           url: https://github.com/gophercises/urlshort
         - path: /urlshort-final
           url: https://github.com/gophercises/urlshort/tree/solution
        `
	
	fmt.Println("Starting the server on port:8080")
	http.ListenAndServe(":8080, handleYAML")
}

func defaultMux() http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/" hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}