package handlers

import (
	"supportlocal/TEDxMileHigh/lib/httputil"
)

var (
	mustWriteHtml = httputil.MustWriteHtml
	mustWriteJson = httputil.MustWriteJson
	readJson      = httputil.ReadJson
	scriptIsland  = httputil.ScriptIsland
)
