package main

import (
	"flag"
	"fmt"

	"github.com/benschweizer/similar/internal"
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

func main() {
	var filter internal.Filter = internal.NumbersFilter

	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&verbose, "v", false, "verbose output")

	var noneFlag = flag.Bool("none", false, "No de-duplication")
	var exactFlag = flag.Bool("exact", false, "De-duplication of successive lines that are identical, ignoring ISO datetimes.")
	var numbersFlag = flag.Bool("numbers", false, "De-duplication of successive lines that are identical when ignoring numbers, e.g., IP addresses, latencies.")
	var signatureFlag = flag.Bool("signature", false, "De-duplication of successive lines that have identical punctuation and whitespace.")

	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	if *noneFlag {
		filter = internal.NoneFilter
	} else if *exactFlag {
		filter = internal.ExactFilter
	} else if *numbersFlag {
		filter = internal.NumbersFilter
	} else if *signatureFlag {
		filter = internal.SignatureFilter
	}

	filenames := flag.Args()
	if len(filenames) == 0 {
		filenames = []string{"-"}
	}

	internal.Process(filenames, filter)
}
