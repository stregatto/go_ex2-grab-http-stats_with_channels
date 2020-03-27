//Package httplib provides core functions for urlsstats like URLs queries
package httplib

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"
)

//Stat collects infos from the url queried
type Stat struct {
	URL           string
	ContentLength int64
	ResponseTime  time.Duration
	DNSQueryTime  time.Duration
	ConnectTime   time.Duration
	TLSHandshake  time.Duration
	TTFB          time.Duration
	TotalTime     time.Duration
	ReturnCode    int
	Err           error
}

//Stats returns a slice of stats type
func Stats(xURL ...string) chan Stat {
	cStats := make(chan Stat)
	for _, URL := range xURL {
		go func(c chan Stat, u string) {
			c <- statFromURL(u)
		}(cStats, URL)
	}
	return cStats
}

//statFromURL collects stats about TTFB (Time to first byte) I'm using
//the below answer from stackoverflow.com
//https://stackoverflow.com/questions/48077098/getting-ttfb-time-to-first-byte-value-in-golang/48077762?r=SearchResults#48077762
//returnCode is -1 if query returns error.
//contentLength is -1 if query returns error or the contentLength is unknown https://golang.org/pkg/net/http/
func statFromURL(URL string) Stat {

	var start, connect, dns, tlsHandshake time.Time
	var DNSQueryTime, ConnectTime, TLSHandshake, TTFB, TotalTime time.Duration

	req, _ := http.NewRequest("GET", URL, nil)
	trace := &httptrace.ClientTrace{
		GetConn:     nil,
		GotConn:     nil,
		PutIdleConn: nil,
		GotFirstResponseByte: func() {
			TTFB = time.Since(start)
		},
		Got100Continue: nil,
		Got1xxResponse: nil,
		DNSStart: func(dsi httptrace.DNSStartInfo) {
			dns = time.Now()
		},
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			DNSQueryTime = time.Since(dns)
		},
		ConnectStart: func(network, addr string) {
			connect = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			ConnectTime = time.Since(connect)
		},
		TLSHandshakeStart: func() {
			tlsHandshake = time.Now()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			TLSHandshake = time.Since(tlsHandshake)
		},
		WroteHeaderField: nil,
		WroteHeaders:     nil,
		Wait100Continue:  nil,
		WroteRequest:     nil,
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	TotalTime = time.Since(start)
	responseTime := time.Since(start)
	var contentLength int64
	var returnCode int

	// fix some response value in case of error
	if err != nil {
		contentLength = -1
		returnCode = -1
	} else {
		contentLength = resp.ContentLength
		returnCode = resp.StatusCode
	}

	return Stat{
		URL:           URL,
		ContentLength: contentLength,
		ResponseTime:  responseTime,
		DNSQueryTime:  DNSQueryTime,
		ConnectTime:   ConnectTime,
		TLSHandshake:  TLSHandshake,
		TTFB:          TTFB,
		TotalTime:     TotalTime,
		ReturnCode:    returnCode,
		Err:           err,
	}
}
