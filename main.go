package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const version = "0.3.0-dev"

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
	fmt.Fprintln(os.Stderr, "  <value> -> string with timestamps in it")
	fmt.Fprintln(os.Stderr, "  -       -> pipe input with timestamps from stdin")
}

func getFormat(formatString string) string {
	switch formatString {
	case "ANSIC":
		return time.ANSIC
	case "UnixDate":
		return time.UnixDate
	case "RubyDate":
		return time.RubyDate
	case "RFC822":
		return time.RFC822
	case "RFC822Z":
		return time.RFC822Z
	case "RFC850":
		return time.RFC850
	case "RFC1123":
		return time.RFC1123
	case "RFC1123Z":
		return time.RFC1123Z
	case "RFC3339":
		return time.RFC3339
	case "RFC3339Nano":
		return time.RFC3339Nano
	case "Kitchen":
		return time.Kitchen
	case "Stamp":
		return time.Stamp
	case "StampMilli":
		return time.StampMilli
	case "StampMicro":
		return time.StampMicro
	case "StampNano":
		return time.StampNano
	default:
		return formatString
	}
}

func getLocation(locationString string) *time.Location {
	location, err := time.LoadLocation(locationString)
	if err == nil {
		return location
	}
	panic(err.Error())
}

func main() {
	locationPtr := flag.String("location", "Local", "tzdata location to convert to")
	formatPtr := flag.String("format", "Mon 2006 Jan 02 03:04pm MST", "format to use")
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

		switch *typePtr {
		case "unix":
			fmt.Println(NewUnixTimestampReplacer().ReplaceDates(arg, format, loc))
		default:
			fmt.Println(NewIso8601Replacer().ReplaceDates(arg, format, loc))
		}
	}
}
