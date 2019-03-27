# timeconverter
[![Go Report Card](https://goreportcard.com/badge/github.com/bbeardsley/timeconverter)](https://goreportcard.com/report/github.com/bbeardsley/timeconverter)
[![GoDoc](https://godoc.org/github.com/bbeardsley/timeconverter?status.svg)](https://godoc.org/github.com/bbeardsley/timeconverter)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

convert timestamps to local time or specified timezone

## install
Direct downloads are available through the [releases page](https://github.com/bbeardsley/timeconverter/releases/latest).

If you have Go installed on your computer just run `go get`.

    go get github.com/bbeardsley/timeconverter

## usage
```
timeconverter [options] <command>

Options
  -format string
        format to use (default "Mon 2006 Jan 02 03:04pm MST")
  -location string
        tzdata location to convert to (default "Local")
  -type string
        what type of timestamps in the input (options "iso8601", "unix") (default "iso8601")
  -version
        print version number and exit
Commands
  help    -> show this help
  version -> print version number and exit
  <value> -> string with timestamps in it
  -       -> pipe input with timestamps from stdin
```
## format

The format is specified using Golang formatting string.  See [docs](https://yourbasic.org/golang/format-parse-string-time-date-example/) for more info and some examples.

## location

The location is specified using the IANA time zone [database](https://www.iana.org/time-zones).
## examples
### command line
#### convert all dates in server log to local time
```
$ timeconverter - < server.log
```

#### convert timestamp to local time
```
$ timeconverter 2019-03-17T00:00:00Z
```

#### convert timestamp to CST/CDT
```
$ timeconverter -location="America/Chicago" 2019-03-17T00:00:00Z
```

#### convert time to local time using built-in ANSIC format
```
$ timeconverter -format="ANSIC" 2019-03-17T00:00:00Z
```
#### convert unix timestamp to local time
```
$ timeconverter -type=unix 1553534903
```

### bash aliases
```
  alias lt='timeconverter'
  alias cst='timeconverter -location="America/Chicago"'
  alias uts='timeconverter -type=unix'
```

