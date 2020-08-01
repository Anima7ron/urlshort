package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"urlshort/students/anima7ron/urlshort"
)

func main() {
	mux := defaultMux()

	help := "JSON file in format:\n[ { \"path\": \"<PATH>\", \"url\": \"<URL>\" } ]\n\nOR\n\nYAML file in format:\n - path: <PATH>\n   url: <URL>\n"
	fileName := flag.String("file", "pathURL.json", help)
	flag.Parse()

	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	Handle, err := urlshort.Handle([]byte(file), filepath.Ext(*fileName), mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", Handle)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
