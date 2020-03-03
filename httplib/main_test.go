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
	xTestStats := map[string]Stat{
		server.URL:                       {url: server.URL, contentLength: 18, responseTime: 2000 * time.Microsecond, returnCode: 200},
		`http://someunexistenturl.wrong`: {url: `http://someunexistenturl.wrong`, contentLength: -1, responseTime: 5000 * time.Microsecond, returnCode: -1},
		`http://verywrong.wrong`:         {url: `http://verywrong.wrong`, contentLength: -1, responseTime: 5000 * time.Microsecond, returnCode: -1},
	}

	// This is the list (slice x) of URLs I would like to test
	var xTestListOfUrls = []string{
		server.URL,
		`http://someunexistenturl.wrong`,
		`http://verywrong.wrong`,
	}
	// let's test the function, not all fields are tested
	cStats := Stats(xTestListOfUrls...)
	for i := 0; i < len(xTestListOfUrls); i++ {
		ts := <-cStats
		if ts.returnCode != xTestStats[ts.url].returnCode {
			t.Errorf("URL: %v\n\tExpected: %v, got: %v", ts.url, xTestStats[ts.url].returnCode, ts.returnCode)
		}
		if ts.responseTime > xTestStats[ts.url].responseTime {
			t.Errorf("URL: %v\n\tExpected response time minor of: %v, got: %v", ts.url, xTestStats[ts.url].responseTime, ts.responseTime)
		}
		if ts.contentLength != xTestStats[ts.url].contentLength {
			t.Errorf("URL: %v\n\tExpected: %v, got: %v", ts.url, xTestStats[ts.url].contentLength, ts.contentLength)
		}
		if ts.url != xTestStats[ts.url].url {
			t.Errorf("Expected: %v, got: %v", xTestStats[ts.url].url, ts.url)
		}
	}
}
