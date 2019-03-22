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
