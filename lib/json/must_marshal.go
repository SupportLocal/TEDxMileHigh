package json

import (
	"encoding/json"
)

func MustMarshal(v interface{}) []byte {
	b, e := json.Marshal(v)

	if e != nil {
		panic(e)
	}

	return b
}
