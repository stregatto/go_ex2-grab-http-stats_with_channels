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
