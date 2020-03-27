package httplib

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStats(t *testing.T) {

	// starting a test server, with content type html and bothering answer
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", `text/html; charset=UTF-8`)
		_, _ = io.WriteString(w, `hello, I'm a test'`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// this is the slice (x) containing what I'm expecting to test
	xTestStats := map[string]Stat{
		server.URL:                       {Url: server.URL, ContentLength: 18, ResponseTime: 2000 * time.Microsecond, ReturnCode: 200},
		`http://someunexistenturl.wrong`: {Url: `http://someunexistenturl.wrong`, ContentLength: -1, ResponseTime: 50000 * time.Microsecond, ReturnCode: -1},
		`verywrong.wrong`:                {Url: `verywrong.wrong`, ContentLength: -1, ResponseTime: 5000 * time.Microsecond, ReturnCode: -1},
	}

	// This is the list (slice x) of URLs I would like to test
	var xTestListOfUrls = []string{
		server.URL,
		`http://someunexistenturl.wrong`,
		`verywrong.wrong`,
	}
	// let's test the function, not all fields are tested
	cStats := Stats(xTestListOfUrls...)
	for i := 0; i < len(xTestListOfUrls); i++ {
		ts := <-cStats
		if ts.ReturnCode != xTestStats[ts.Url].ReturnCode {
			t.Errorf("URL: %v\n\tExpected: %v, got: %v", ts.Url, xTestStats[ts.Url].ReturnCode, ts.ReturnCode)
		}
		if ts.ResponseTime > xTestStats[ts.Url].ResponseTime {
			t.Errorf("URL: %v\n\tExpected response time minor of: %v, got: %v", ts.Url, xTestStats[ts.Url].ResponseTime, ts.ResponseTime)
		}
		if ts.ContentLength != xTestStats[ts.Url].ContentLength {
			t.Errorf("URL: %v\n\tExpected: %v, got: %v", ts.Url, xTestStats[ts.Url].ContentLength, ts.ContentLength)
		}
		if ts.Url != xTestStats[ts.Url].Url {
			t.Errorf("Expected: %v, got: %v", xTestStats[ts.Url].Url, ts.Url)
		}
	}
}

func BenchmarkStats(b *testing.B) {
	// starting a test server, with content type html and bothering answer
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", `text/html; charset=UTF-8`)
		_, _ = io.WriteString(w, `hello, I'm a test'`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	// just a local URL to be tested, this is a benchmark.
	var xTestListOfUrls = []string{
		server.URL,
	}
	for i := 0; i < b.N; i++ {
		cStats := Stats(xTestListOfUrls...)
		// we can throw out everything is returned.
		_ = <-cStats
	}

}
