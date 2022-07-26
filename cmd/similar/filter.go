package main

import (
	"regexp"
)

// Current filters implement the Grafana log dedup logic, see
// https://github.com/grafana/grafana/blob/main/public/app/core/logsModel.ts
// The implementation is pretty slow due to massive regex usage, this could be
// rewritten more efficiently with byte operations instead.
type Filter func(string) string

// Filter none
func noneFilter(s string) string {
	return s
}

// Filter ISO8601 with millis
func exactFilter(s string) string {
	// FIXME: filter isodate
	isoDateRegexp := `\d{4}-[01]\d-[0-3]\dT[0-2]\d:[0-5]\d:[0-6]\d[,\.]\d+([+-][0-2]\d:[0-5]\d|Z)`
	m := regexp.MustCompile(isoDateRegexp)
	ret := string(m.ReplaceAllString(s, ""))
	return ret
}

// Filter numbers
func numbersFilter(s string) string {
	m := regexp.MustCompile(`[0-9]`)
	ret := string(m.ReplaceAllString(s, ""))
	return ret
}

// Filter signature
func signatureFilter(s string) string {
	m := regexp.MustCompile(`\w`)
	ret := string(m.ReplaceAllString(s, ""))
	return ret
}
