package file

import (
	"testing"
)

const testFilename string = `url_test.list`

var testListOfUrls = []string{
	`http://www.google.com`,
	`https://duckduckgo.com/?q=duckduckgo&t=ffab&atb=v205-1&ia=web`,
}

func TestLoad(t *testing.T) {
	xs := Load(testFilename)
	for i, v := range xs {
		if v != testListOfUrls[i] {
			t.Errorf("\nExpected:\t%v\ngot:\t\t%v", testListOfUrls[i], v)
		}
	}
}
