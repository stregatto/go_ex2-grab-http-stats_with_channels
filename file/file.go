//Package file collect all functions are useful to manage files
package file

import (
	"bufio"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Load returns a slice of string from a given file.
func Load(s string) []string {
	var line []string
	f, err := os.Open(s)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	return line
}
