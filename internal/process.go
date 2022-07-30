package internal

import (
	"bufio"
	"fmt"
	"os"
)

// Loops over a list of files, applies a filter and prints non-deduplicatedlines
func Process(filenames []string, filter Filter) {
	var f *os.File
	var firstline string
	var filtered_firstline string
	var line string
	var filtered_line string
	var cnt uint64

	for _, filename := range filenames {
		if filename == "-" {
			f = os.Stdin
		} else {
			var err error
			f, err = os.Open(filename)
			if err != nil {
				println("failed to open", filename)
				continue
			}
		}

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		if !scanner.Scan() {
			continue
		}
		firstline = scanner.Text()
		cnt = 1
		for scanner.Scan() {
			line = scanner.Text()
			filtered_firstline = filter(firstline)
			filtered_line = filter(line)
			if filtered_line == filtered_firstline {
				cnt++
			} else {
				fmt.Println(cnt, "\t", firstline)
				cnt = 1
				firstline = line
			}
		}
		fmt.Println(cnt, "\t", firstline)

		f.Close()
	}
}
