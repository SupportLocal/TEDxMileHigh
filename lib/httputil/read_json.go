package httputil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ReadJson will parses the JSON-encoded data in the http request and store the result in v
func ReadJson(r *http.Request, v interface{}) (err error) {
	defer r.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, v); err != nil {
		return
	}

	return
}
