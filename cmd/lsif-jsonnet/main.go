package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

func findRepo(dir string) (string, error) {
	inRepo, err := refs.PathExists(filepath.Join(dir, ".git"))
	if err != nil {
		return "", err
	}
	if inRepo {
		return dir, nil
	}

	if len(dir) == 1 && dir[0] == filepath.Separator {
		return "", errors.New("no repo found")
	}

	return findRepo(filepath.Dir(dir))
}

func collectInputFiles(args []string) ([]string, error) {
	var res []string

	if len(args) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		args = []string{cwd}
	}

	for _, arg := range args {
		absArg, err := filepath.Abs(arg)
		if err != nil {
			return nil, err
		}

		fi, err := os.Stat(absArg)
		if err != nil {
			return nil, err
		}

		if fi.IsDir() {
			xs, err := collectInputsInDir(absArg)
			if err != nil {
				return nil, err
			}
			res = append(res, xs...)
		} else {
			res = append(res, absArg)
		}
	}
	return res, nil
}

func collectInputsInDir(dir string) ([]string, error) {
	var res []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".jsonnet") {
			res = append(res, path)
		}
		return nil
	})
	return res, err
}

func mainImpl() error {
	var help = flag.BoolP("help", "h", false, "This message")
	var jps = flag.StringArrayP("jpath", "J", nil, "Optional. Specify an additional library search dir (right-most wins)")
	var outFile = flag.StringP("output-file", "o", "", "Optional. Write to the output file rather than stdout")
	var projectRoot = flag.StringP("project-root", "p", "", "Optional. Specifies which dir is the project root. Defaults to repo dir where sources are located.")

	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "lsif-jsonnet %s\n", versionString)
		fmt.Fprintln(os.Stderr, `Basic usage: lsif-jsonnet -J <libpath> -o <out_file> -p <project-root> <jsonnet_files or directories>

If no directories or jsonnet input files are specified, it will look for files with .jsonnet suffix in current working dir tree.`)
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

	inputs, err := collectInputFiles(flag.Args())
	if err != nil {
		return err
	}

	if len(inputs) == 0 {
		return errors.New("no jsonnet input files found")
	}

	var lls []*refs.Listener

	for _, absInFile := range inputs {
		inDir := filepath.Dir(absInFile)
		err = pathResolver.AddPath(inDir)
		if err != nil {
			return fmt.Errorf("failed to add %s to path resolver: %w\n", inDir, err)
		}

		ll, err := refs.ParseFile(absInFile, pathResolver)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w\n", absInFile, err)
		}

		lls = append(lls, ll)
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

	root := *projectRoot
	if root == "" {
		repoRoot, err := findRepo(filepath.Dir(inputs[0]))
		if err != nil {
			return errors.New("couldn't determine project root")
		}
		root = repoRoot
	}

	dmpr := dumper.NewDumper(root, outWriter)

	return dmpr.DumpProject(toolInfo, lls)
}
