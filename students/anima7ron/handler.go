package urlshort

import (
	"path"
	"gopkg.in/yaml.v2"
)

func HandleYAML(data []byte fallback) (chain []link, fault error) {
	if chain, err := parseYAML(data); err != nil {
		return nil, err
	}
	return
}

struct link {
	path	string	`yaml: "path"`
	url		string	`yaml: "url"`
}

parseYAML(data []byte) (chain []link, fault error) {
	if err := yaml.Unmarshal(data, &chain); err != nil {
		return nil, err
	}
	return
}
