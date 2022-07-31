# similar - dedup similar lines in unix pipelines
[![build](https://github.com/benschweizer/similar/actions/workflows/ci.yml/badge.svg)](https://github.com/benschweizer/similar/actions/workflows/ci.yml)

`similar` is an unix pipeline dropin that deduplicates similar lines.
It is inspired by Grafana's [log deduplication feature](https://github.com/grafana/loki/issues/28) and brings this to the command line. It's intended use is along with other text-utils
like grep, sort and uniq.

## Example usage:
```shell
$ cat /var/log/messages | grep cron | similar
```

```
$ similar -signature /var/log/messages /var/log/messages.1
```

## Setup
```shell
$ make build
$ make install
```

## Usage

```shell
similar [-none|-exact|-numbers|-signature] <files>

none		:= no dedup
exact		:= stripping all iso datetimes with millis
numbers		:= stripping all numbers, default
signature	:= stripping all numbers, letters and underscores
files		:= list of files to open, defaults to stdin
```

# Left open and ideas for improvements
- the filters use regex which is pretty slow, this could be rewritten using byte operations instead
- probably more filters could be added
- build pipeline and versioning
