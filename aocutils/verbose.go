package aocutils

import (
	"flag"
	"log"
	"os"
)

// Verbose is true when the program was run with the -verbose flag.
var Verbose bool

// OutputFile is the file verbose output gets written to when -f is set.
// If empty, verbose output goes to stdout.
var OutputFile string

var logger *log.Logger

// Initiate parses CLI flags (-verbose, -f) and should be called once at
// the top of main().
func Initiate() {
	flag.BoolVar(&Verbose, "verbose", false, "enable verbose output")
	flag.StringVar(&OutputFile, "f", "", "save verbose output to the given file instead of stdout")
	flag.Parse()

	if !Verbose {
		return
	}

	out := os.Stdout
	if OutputFile != "" {
		f, err := os.Create(OutputFile)
		if err != nil {
			log.Fatalf("failed to create output file %s: %v", OutputFile, err)
		}
		out = f
	}

	logger = log.New(out, "", 0)
	logger.Println("===Verbose mode enabled===")
}

// VPrintf writes formatted output when verbose mode is enabled, either to
// stdout or to the file passed via -f.
func VPrintf(format string, args ...any) {
	if Verbose && logger != nil {
		logger.Printf(format, args...)
	}
}
