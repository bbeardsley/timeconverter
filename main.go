package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

const version = "0.3.2"
const unixSeconds = "UnixSeconds"

var builtInFormats = map[string]string{
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
	unixSeconds:   unixSeconds,
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage")
	fmt.Fprintln(os.Stderr, "    timeconverter [options] <command>")
	fmt.Fprintln(os.Stderr, "Version")
	fmt.Fprintln(os.Stderr, "    "+version)
	fmt.Fprintln(os.Stderr, "Options")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "Commands")
	fmt.Fprintln(os.Stderr, "  help    -> show this help")
	fmt.Fprintln(os.Stderr, "  version -> print version number and exit")
	fmt.Fprintln(os.Stderr, "  now     -> print current timestamp for location and format")
	fmt.Fprintln(os.Stderr, "  <value> -> string with timestamps in it")
	fmt.Fprintln(os.Stderr, "  -       -> pipe input with timestamps from stdin")
}

func getBuiltInFormatKeys() []string {
	keys := make([]string, 0)
	for key := range builtInFormats {
		keys = append(keys, key)
	}
	return keys
}

func getFormat(formatString string) string {
	builtInFormat, ok := builtInFormats[formatString]
	if ok {
		return builtInFormat
	}
	return formatString
}

func getLocation(locationString string) *time.Location {
	location, err := time.LoadLocation(locationString)
	if err == nil {
		return location
	}
	panic(err.Error())
}

func main() {
	builtinFormatKeys := getBuiltInFormatKeys()
	sort.Strings(builtinFormatKeys)
	locationPtr := flag.String("location", "Local", "tzdata location to convert to")
	formatPtr := flag.String("format", "Mon 2006 Jan 02 03:04pm MST", "format to use (options \""+strings.Join(builtinFormatKeys, "\", \"")+"\"")
	versionPtr := flag.Bool("version", false, "print version number and exit")
	typePtr := flag.String("type", "iso8601", "what type of timestamps in the input (options \"iso8601\", \"unix\")")
	flag.Parse()

	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	arg := flag.Arg(0)
	switch arg {
	case "", "h", "-h", "--h", "/h", "/?", "help", "-help", "--help", "/help":
		printUsage()
		os.Exit(1)
	case "version", "-version", "--version", "/version":
		fmt.Println(version)
		os.Exit(0)
	case "-":
		// pipe
		loc := getLocation(*locationPtr)
		format := getFormat(*formatPtr)

		scanner := bufio.NewScanner(os.Stdin)

		switch *typePtr {
		case "unix":
			replacer := NewUnixTimestampReplacer()
			for scanner.Scan() {
				fmt.Println(replacer.ReplaceDates(scanner.Text(), format, loc))
			}
		default:
			replacer := NewIso8601Replacer()
			for scanner.Scan() {
				fmt.Println(replacer.ReplaceDates(scanner.Text(), format, loc))
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	default:
		loc := getLocation(*locationPtr)
		format := getFormat(*formatPtr)

		switch {
		case arg == "now":
			fmt.Println(time.Now().In(loc).Format(format))
		case *typePtr == "unix":
			fmt.Println(NewUnixTimestampReplacer().ReplaceDates(arg, format, loc))
		default:
			fmt.Println(NewIso8601Replacer().ReplaceDates(arg, format, loc))
		}
	}
}
