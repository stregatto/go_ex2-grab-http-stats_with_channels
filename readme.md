# Go exercise, urlsstats

This is a simple exercise now trying to use channels.

Given a list of url in a file named `list_of_urls.list` the program returns some statistics.

Like:
```
URL: http://www.wikipedia.com
        Return code:            301
        Content length [bytes]: 185
        Response time:          65.918934ms
        DNS query time:         4.574314ms
        Connect time:           28.192658ms
        TLS handshake time:     0s
        Time to first bite:     65.868282ms
```

This program checks all URL in concurrent mode using _goroutines_

## Usage

`# ./urlsstats -h`
```
Usage of ./urlsstats:
  -f string
        The name of the file containing the list of urls you want to test (default "list_of_urls.list")
  -o string
        [STDOUT|json] prints the output.
        stdout: DEFAULT pretty print on stdout
        json: print output in json format on stdout
         (default "stdout")
```

## Binary

The binary `urlsstats` is build for OSX Catalina.

A very usefull guide to build golang software for other platforms https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

