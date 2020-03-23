package main

import (
	"flag"
	"github.com/stregatto/urlsstats/file"
	"github.com/stregatto/urlsstats/httplib"
)

func main() {
	inputFile := flag.String("f", "list_of_urls.list", "The name of the file containing the list of urls you want to test")
	flag.Parse()
	xURL := file.Load(*inputFile)
	chanStats := httplib.Stats(xURL...)
	httplib.Print(chanStats, len(xURL))
}
