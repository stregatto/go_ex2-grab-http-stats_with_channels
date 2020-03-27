package main

import (
	"flag"
	"github.com/stregatto/urlsstats/file"
	"github.com/stregatto/urlsstats/httplib"
	"github.com/stregatto/urlsstats/output"
)

func main() {
	inputFile := flag.String("f", "list_of_urls.list", "The name of the file containing the list of urls you want to test")
	outputFlag := flag.String("o", "stdout", "[STDOUT|json] prints the output.\n"+
		"stdout: DEFAULT pretty print on stdout\n"+
		"json: print output in json format on stdout\n")
	flag.Parse()
	xURL := file.Load(*inputFile)
	chanStats := httplib.Stats(xURL...)
	switch *outputFlag {
	case "json":
		output.Jprint(chanStats, len(xURL))
	default:
		output.Print(chanStats, len(xURL))
	}
}
