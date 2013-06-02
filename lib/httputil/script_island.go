package httputil

import (
	"fmt"
	"supportlocal/TEDxMileHigh/lib/json"
)

func ScriptIsland(name string, v interface{}) string {
	return fmt.Sprintf(`<script id="%s" type="application/json">%s</script>`, name, json.MustMarshal(v))
}
