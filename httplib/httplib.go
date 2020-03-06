package httplib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

//Stat collects infos from the url queried
type Stat struct {
	url           string
	contentLength int64
	responseTime  time.Duration
	DNSQueryTime  time.Duration
	ConnectTime   time.Duration
	TLSHandshake  time.Duration
	TTFB          time.Duration
	TotalTime     time.Duration
	returnCode    int
	err           error
}

// Print prints all infos collected from Stats
func Print(c <-chan Stat, n int) {
	for i := 0; i < n; i++ {
		s := <-c
		fmt.Printf("URL: %v\n", s.url)
		if s.err == nil {
			fmt.Printf("\tReturn code:\t\t%v\n", s.returnCode)
			fmt.Printf("\tContent length [bytes]:\t%v\n", s.contentLength)
			fmt.Printf("\tResponse time:\t\t%v\n", s.responseTime)
			fmt.Printf("\tDNS query time:\t\t%v\n", s.DNSQueryTime)
			fmt.Printf("\tConnect time:\t\t%v\n", s.ConnectTime)
			fmt.Printf("\tTLS handshake time:\t%v\n", s.TLSHandshake)
			fmt.Printf("\tTime to first bite:\t%v\n", s.TTFB)
		} else {
			fmt.Printf("\tError: %v\n", s.err)
		}
	}
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
		url:           URL,
		contentLength: contentLength,
		responseTime:  responseTime,
		DNSQueryTime:  DNSQueryTime,
		ConnectTime:   ConnectTime,
		TLSHandshake:  TLSHandshake,
		TTFB:          TTFB,
		TotalTime:     TotalTime,
		returnCode:    returnCode,
		err:           err,
	}
}
