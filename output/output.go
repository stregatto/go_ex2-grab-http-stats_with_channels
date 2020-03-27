// Package output provides all function to print the data.
package output

import (
	"encoding/json"
	"fmt"
	"github.com/stregatto/urlsstats/httplib"
	"log"
)

// Print pretty prints all infos collected from Stats to standard output
func Print(c <-chan httplib.Stat, n int) {
	for i := 0; i < n; i++ {
		s := <-c
		fmt.Printf("URL: %v\n", s.URL)
		if s.Err == nil {
			fmt.Printf("\tReturn code:\t\t%v\n", s.ReturnCode)
			fmt.Printf("\tContent length [bytes]:\t%v\n", s.ContentLength)
			fmt.Printf("\tResponse time:\t\t%v\n", s.ResponseTime)
			fmt.Printf("\tDNS query time:\t\t%v\n", s.DNSQueryTime)
			fmt.Printf("\tConnect time:\t\t%v\n", s.ConnectTime)
			fmt.Printf("\tTLS handshake time:\t%v\n", s.TLSHandshake)
			fmt.Printf("\tTime to first bite:\t%v\n", s.TTFB)
		} else {
			fmt.Printf("\tError: %v\n", s.Err)
		}
	}
}

// Jprint pretty prints all infos collected from Stats to standard output in json format
func Jprint(c <-chan httplib.Stat, n int) {
	stats := make([]httplib.Stat, n)
	for i := 0; i < n; i++ {
		s := <-c
		stats[i] = s
	}
	joutput, err := json.Marshal(stats)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(joutput))
}
