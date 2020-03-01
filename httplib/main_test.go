package httplib

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*type Stats struct {
	url          string
	contentLength int
	responseTime int
	returnCode   int8
}*/

func TestStats(t *testing.T) {

	// starting a test server, with content type html and bothering answer
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", `text/html; charset=UTF-8`)
		io.WriteString(w, `hello, I'm a test'`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// this is the slice (x) containing what I'm expecting to test
	xTestStats := []Stat{
		{url: server.URL, contentLength: 18, responseTime: 2000 * time.Microsecond, returnCode: 200},
	}

	// This is the list (slice x) of URLs I would like to test
	var xTestListOfUrls = []string{
		server.URL,
	}
	// let's test the function, not all fields are tested
	cStats := Stats(xTestListOfUrls...)
	for i := 0; i < len(xTestListOfUrls); i++ {
		ts := <-cStats
		if ts.returnCode != xTestStats[i].returnCode {
			t.Errorf("Expected: %v, got: %v", xTestStats[i].returnCode, ts.returnCode)
		}
		if ts.responseTime > xTestStats[i].responseTime {
			t.Errorf("Expected response time minor of: %v, got: %v", xTestStats[i].responseTime, ts.responseTime)
		}
		if ts.contentLength != xTestStats[i].contentLength {
			t.Errorf("Expected: %v, got: %v", xTestStats[i].contentLength, ts.contentLength)
		}
		if ts.url != xTestStats[i].url {
			t.Errorf("Expected: %v, got: %v", xTestStats[i].url, ts.url)
		}
	}
}
