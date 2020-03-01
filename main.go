package main

import (
	"github.com/stregatto/go_ex1-grab-http-stats/file"
	"github.com/stregatto/go_ex1-grab-http-stats/httplib"
)

func main() {
	xURL := file.Load("list_of_urls.list")
	chanStats := httplib.Stats(xURL...)
	httplib.Print(chanStats, len(xURL))
}
