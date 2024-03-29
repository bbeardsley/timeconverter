package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const version = "0.3.6"

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
	for key := range TimeFormats {
		keys = append(keys, key)
	}
	return keys
}

func getLocationAliasKeys() []string {
	keys := make([]string, 0)
	for key := range LocationAliases {
		keys = append(keys, key)
	}
	return keys
}

func getFormat(formatString string) string {
	builtInFormat, ok := TimeFormats[formatString]
	if ok {
		return builtInFormat
	}
	return formatString
}

func getLocation(locationString string) *time.Location {
	locationFromAlias, ok := LocationAliases[locationString]
	if ok {
		location, err := time.LoadLocation(locationFromAlias)

		if err == nil {
			return location
		}
		panic(err.Error())
	} else {
		location, err := time.LoadLocation(locationString)
		if err == nil {
			return location
		}
		panic(err.Error())
	}
}

func main() {
	builtinFormatKeys := getBuiltInFormatKeys()
	sort.Strings(builtinFormatKeys)
	locationPtr := flag.String("location", "Local", "tzdata location to convert to (aliases \""+strings.Join(getLocationAliasKeys(), "\", \"")+"\")")
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
			if format == "UnixSeconds" {
				fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
			} else {
				fmt.Println(time.Now().In(loc).Format(format))
			}
		case *typePtr == "unix":
			fmt.Println(NewUnixTimestampReplacer().ReplaceDates(arg, format, loc))
		default:
			fmt.Println(NewIso8601Replacer().ReplaceDates(arg, format, loc))
		}
	}
}
