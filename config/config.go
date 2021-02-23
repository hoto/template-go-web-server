package config

import (
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

var (
	Version     string
	ShortCommit string
	BuildDate   string
	Debug       bool
	Port	 	int
)

func ParseArgsAndFlags() {
	flag.Usage = overrideUsage()

	flag.BoolVar(&Debug, "debug", false, "Show verbose debug information")
	showVersion := flag.Bool("version", false, "Show version")
	flag.IntVar(&Port, "port", 5000, "Set server port")

	flag.Parse()

	if *showVersion {
		fmt.Printf("template-go-web-server version %s, commit %s, build %s\n",
			Version, ShortCommit, BuildDate)
		os.Exit(0)
	}

	if Debug {
		debugLog()
	}
}

func overrideUsage() func() {
	return func() {
		_, _ = fmt.Fprintf(
			os.Stdout,
			"Usage:"+
				"\n\t"+
				"template-go-web-server [flags]"+
				"\n\n"+
				"Flags:"+
				"\n")
		flag.PrintDefaults()
	}
}

func debugLog() {
	fmt.Println("Args:")
	fmt.Printf("  args=%s\n", aurora.Cyan(flag.Args()))
}
