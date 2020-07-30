package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathURL map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathURL[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	chain, err := parseYAML(data)
	if err != nil {
		return nil, err
	}
	pathURL := mapChain(chain)
	return MapHandler(pathURL, fallback), nil
}

type link struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(data []byte) (chain []link, fault error) {
	err := yaml.Unmarshal(data, &chain)
	if err != nil {
		return nil, err
	}
	return
}

func mapChain(chain []link) map[string]string {
	pathURL := make(map[string]string)
	for _, pu := range chain {
		pathURL[pu.Path] = pu.URL
	}
	return pathURL
}
