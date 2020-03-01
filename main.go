package main

import (
	"github.com/stregatto/go_ex2-grab-http-stats_with_channels/file"
	"github.com/stregatto/go_ex2-grab-http-stats_with_channels/httplib"
)

func main() {
	xURL := file.Load("list_of_urls.list")
	chanStats := httplib.Stats(xURL...)
	httplib.Print(chanStats, len(xURL))
}
