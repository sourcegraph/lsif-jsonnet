package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"lsif-jsonnet/dumper"
	"lsif-jsonnet/protocol"
	"lsif-jsonnet/refs"
)

const version = "1.0.0"
const versionString = version + ", protocol version " + protocol.Version

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func mainImpl() error {
	var help *bool = flag.BoolP("help", "h", false, "This message")
	var jps *[]string = flag.StringArrayP("jpath", "J", nil, "Specify an additional library search dir (right-most wins)")
	var outFile *string = flag.StringP("output-file", "o", "", "Write to the output file rather than stdout")

	flag.Parse()

	if *help || flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "lsif-jsonnet %s\n", versionString)
		fmt.Fprintln(os.Stderr, "Basic usage: lsif-jsonnet -J <libpath> -o <out_file> <jsonnet_file>")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var lps []string

	for _, p := range filepath.SplitList(os.Getenv("JSONNET_PATH")) {
		lps = append(lps, p)
	}

	for _, pp := range *jps {
		for _, p := range filepath.SplitList(pp) {
			lps = append(lps, p)
		}
	}

	pathResolver, err := refs.NewPathResolver(lps)
	if err != nil {
		return fmt.Errorf("failed to instantiate path resolver: %w\n", err)
	}

	inFile := flag.Arg(0)

	absInFile, err := filepath.Abs(inFile)
	if err != nil {
		return fmt.Errorf("failed to make pathh %s absolute: %w\n", inFile, err)
	}

	inDir := filepath.Dir(absInFile)
	err = pathResolver.AddPath(inDir)
	if err != nil {
		return fmt.Errorf("failed to add %s to path resolver: %w\n", inDir, err)
	}

	ll, err := refs.ParseFile(absInFile, pathResolver)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w\n", inFile, err)
	}

	var outWriter io.Writer

	if len(*outFile) > 0 {
		f, err := os.Create(*outFile)
		if err != nil {
			return fmt.Errorf("failed to create %s: %w\n", *outFile, err)
		}
		defer f.Close()

		bw := bufio.NewWriter(f)
		defer bw.Flush()

		outWriter = bw
	} else {
		outWriter = os.Stdout
	}

	toolInfo := &protocol.ToolInfo{
		Name:    "lsif-jsonnet",
		Version: version,
		Args:    os.Args[1:],
	}

	dmpr := dumper.NewDumper(inDir, outWriter)

	return dmpr.DumpProject(toolInfo, ll)
}
