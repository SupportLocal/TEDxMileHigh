package form

import (
	"supportlocal/TEDxMileHigh/lib/httputil"
)

var (
	readJson      = httputil.ReadJson
	writeJson     = httputil.WriteJson
	mustWriteJson = httputil.MustWriteJson
)
