package urlshort

import (
	"net/http"

	yamlV2 "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yaml, err := parseYAML(yml)
	return MapHandler(yaml, fallback), err
}

func parseYAML(yaml []byte) (res map[string]string, err error) {
	out := []map[string]string{}
	err = yamlV2.Unmarshal(yaml, &out)
	if err != nil {
		panic(err)
	}

	res = make(map[string]string)
	for _, v := range out {
		res[v["path"]] = v["url"]
	}

	return res, err
}
