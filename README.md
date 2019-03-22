# timeconverter
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
```
$ timeconverter 2019-03-17T00:00:00Z
```

```
$ timeconverter -location="America/Chicago" 2019-03-17T00:00:00Z
```

```
$ timeconverter -format="ANSIC" 2019-03-17T00:00:00Z
```
### bash aliases
```
  alias lt='timeconverter'
  alias cst='timeconverter -location="America/Chicago"'
```

