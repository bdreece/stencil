/*
# stencil

A command-line tool for generic code generation using Go template files.

Template authoring is assisted by the inclusion of the [github.com/Masterminds/sprig/v3] library.

	Usage: stencil [options] <template-file>
		-data.file string
		      path to JSON data file
		-data.json string
		      inline JSON data
		-include value
		      glob pattern of included templates (repeat for multiple patterns)
		-level int
		      log level (debug -4, info 0, warn 4, error 8) (default 4)
		-out string
		      path to generated output
*/
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

const banner string = `
      |                  _) |
  __| __|  _ \ __ \   __| | |
\__ \ |    __/ |   | (    | |
____/\__|\___|_|  _|\___|_|_|

stencil is a command-line tool for generic code generation using Go template files.

Usage: stencil [options] <template-file>
`

var (
	input string
	opts  options
)

func init() {
	const (
		descLevel   string = "log level (debug -4, info 0, warn 4, error 8)"
		descOut     string = "path to generated output"
		descInline  string = "inline JSON data"
		descFile    string = "path to JSON data file"
		descInclude string = "glob pattern of included templates (repeat for multiple patterns)"
	)

	flag.IntVar(&opts.level, "level", 4, descLevel)
	flag.StringVar(&opts.out, "out", "", descOut)
	flag.StringVar(&opts.inline, "data.json", "", descInline)
	flag.StringVar(&opts.file, "data.file", "", descFile)
	flag.Func("include", descInclude, func(s string) error {
		opts.include = append(opts.include, s)
		return nil
	})

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, banner)
		flag.PrintDefaults()
	}
}

func main() {
	defer exit()
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.Level(opts.level),
		})),
	)

	g := generator{
		input:   flag.Arg(0),
		options: opts,
	}

	if err := g.generate(); err != nil {
		panic(err)
	}
}

func exit() {
	if r := recover(); r != nil {
		slog.Error("unexpected panic occurred", "err", r)
		os.Exit(1)
	}
}
