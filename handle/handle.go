package handle

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler wil return an http.HandleFunc (which also
// implements http.Handler) that will attempt to map any paths
// (keys in the map) to their corresponding URL
// (values in the map), if the path is not provided in the map
// then the fallback http.Handler will be called instead.

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// if the path is matched
		// redirect to the url (value in the map)
		path := r.URL.Path
		if destination, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, destination, http.StatusFound)
			return
		}
		// else
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HanderFunc which also implements http.Handler
// that will attempt to map any paths to their corresponding URL
// if the path is not provided in the YAML, then the fallback http.Handler will be called instead

// YAML is expected to be in this format:
// - path : /path
//   url : http:../url

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {

	// 1- parse the yaml
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	// 2- convert YAML array into Map
	pathsToUrls := make(map[string]string)
	for _, path := range pathUrls {
		pathsToUrls[path.Path] = path.URL
	}

	// 3- return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
