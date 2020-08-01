package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Handle parses provided YAML or JSON and attempts to map paths to
// corresponding URLs. It returns an http.HandlerFunc. Any path not
// provided in the file  will call fallback http.Handler instead.
//
// JSON is expected to be in the format:
//
//     [{  "path": "<PATH>",
//         "url": "<URL>"    }]
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// Returned errors concern invalid data formatting. ext fmt ".json"
func Handle(data []byte, ext string, fallback http.Handler) (http.HandlerFunc, error) {
	chain, err := parse(data, ext)
	if err != nil {
		return nil, err
	}
	pathURL := mapChain(chain)
	return mapHandler(pathURL, fallback), nil
}

func parse(data []byte, ext string) (chain []link, fault error) {
	switch ext {
	case ".yaml":
		if err := yaml.Unmarshal(data, &chain); err != nil {
			return nil, err
		}
	case ".json":
		if err := json.Unmarshal(data, &chain); err != nil {
			return nil, err
		}
	}
	return
}

type link struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func mapChain(chain []link) map[string]string {
	pathURL := make(map[string]string)
	for _, pu := range chain {
		pathURL[pu.Path] = pu.URL
	}
	return pathURL
}

func mapHandler(pathURL map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathURL[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
