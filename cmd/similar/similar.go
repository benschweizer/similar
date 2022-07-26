package main

import (
	"flag"
	"fmt"
)

const usage = `similar is an unix pipeline dropin that deduplicates similar lines. It is inspired
by Grafana's log deduplication feature and brings this to the command line. It's
intended use is along with other text-utils like grep, sort and uniq.

usage:

	similar [-none|-exact|-numbers|-signature] <files>

	none		:= no dedup
	exact		:= stripping all iso datetimes with millis
	numbers		:= stripping all numbers, default
	signature	:= stripping all numbers, letters and underscores
	files		:= list of files to open, defaults to stdin

example:

	cat /var/log/messages | grep cron | similar
	similar -signature /var/log/messages /var/log/messages.1
`

var verbose bool
var no_counter bool

func main() {
	var filter Filter = numbersFilter

	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&verbose, "v", false, "verbose output")

	var noneFlag = flag.Bool("none", false, "No de-duplication")
	var exactFlag = flag.Bool("exact", false, "De-duplication of successive lines that are identical, ignoring ISO datetimes.")
	var numbersFlag = flag.Bool("numbers", false, "De-duplication of successive lines that are identical when ignoring numbers, e.g., IP addresses, latencies.")
	var signatureFlag = flag.Bool("signature", false, "De-duplication of successive lines that have identical punctuation and whitespace.")

	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	if *noneFlag {
		filter = noneFilter
	} else if *exactFlag {
		filter = exactFilter
	} else if *numbersFlag {
		filter = numbersFilter
	} else if *signatureFlag {
		filter = signatureFilter
	}

	filenames := flag.Args()
	if len(filenames) == 0 {
		filenames = []string{"-"}
	}

	process(filenames, filter)
}
