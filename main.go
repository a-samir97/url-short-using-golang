package main

import (
	"fmt"
	"net/http"
	"urlshort/handle"
)

func main() {

	mux := defaultMux()

	// build MapHandler
	pathsToUrls := map[string]string{

		"/github/ahmedsamir": "https://www.github.com/a-samir97",
		"/linked/ahmedsamir": "https://www.linkedin.com/in/ahmedsamir97",
	}

	mapHandler := handle.MapHandler(pathsToUrls, mux)

	// build YAMLHandler
	yaml := `
- path: /github
  url: https://github.com/a-samir97
- path: /linkedin
  url: https://linkedin.com/in/ahmedsamir97
`
	yamlHandler, err := handle.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on port : 8080")

	http.ListenAndServe(":8080", yamlHandler)

}

func defaultMux() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, World")
}
