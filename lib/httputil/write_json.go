package httputil

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteJson(w http.ResponseWriter, v interface{}) error {

	// avoid json vulnerabilities, always wrap v in an object literal
	doc := map[string]interface{}{"d": v}

	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return nil
}
