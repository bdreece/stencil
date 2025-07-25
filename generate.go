package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

type options struct {
	level   int
	out     string
	inline  string
	file    string
	include []string
}

type generator struct {
	options

	input string
}

var errInfer = errors.New(`failed to infer output filename: expected input filename convention "<name>.<ext>.*"`)

func (g *generator) generate() error {
	var (
		data any
		err  error
	)

	if g.input, err = filepath.Abs(g.input); err != nil {
		return fmt.Errorf("failed to resolve input file path: %w", err)
	}

	if g.out, err = g.resolveOutput(); err != nil {
		return fmt.Errorf("failed to resolve output file path: %w", err)
	}

	if err := g.readFile(&data); err != nil {
		return err
	}

	if err := g.readInline(&data); err != nil {
		return err
	}

	slog.Info("prepared input data", "data", data)

	srcs, err := g.resolveGlobs()
	if err != nil {
		return err
	}

	render := g.renderDefault
	if filepath.Ext(g.out) == ".go" {
		render = g.renderGo
	}

	if err = render(srcs, data, g.out); err != nil {
		return err
	}

	slog.Info("done!")
	return nil
}

func (g *generator) readFile(data any) error {
	if g.file == "" {
		return nil
	}

	g.file = filepath.Clean(g.file)

	slog.Debug("reading input file...", "file", g.file)
	bytes, err := os.ReadFile(g.file)
	if err != nil {
		return errOpenFile(g.file, err)
	}

	slog.Debug("unmarshaling JSON input file...")
	if err = json.Unmarshal(bytes, &data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON input file: %w", err)
	}

	return nil
}

func (g *generator) readInline(data any) error {
	if g.inline == "" {
		return nil
	}

	slog.Debug("unmarshaling inline JSON input data...", "text", g.inline)
	if err := json.Unmarshal([]byte(g.inline), &data); err != nil {
		return fmt.Errorf("failed to unmarshal inline JSON input data: %w", err)
	}

	return nil
}

func (g *generator) resolveOutput() (string, error) {
	if g.out != "" {
		abs, err := filepath.Abs(g.out)
		if err != nil {
			return "", err
		}

		finfo, err := os.Stat(abs)
		if err != nil {
			return "", err
		}

		if !finfo.IsDir() {
			return abs, nil
		}

		return g.inferOutput(abs)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return g.inferOutput(cwd)
}

func (g *generator) inferOutput(cwd string) (string, error) {
	slog.Debug("inferring output filename by convention", "input", g.input)
	filename := filepath.Base(g.input)
	ext := filepath.Ext(filename)
	trimmed := strings.TrimSuffix(filename, ext)
	if ext = filepath.Ext(trimmed); len(ext) < 1 {
		return "", errInfer
	}

	output := filepath.Join(cwd, trimmed)
	slog.Info("inferred output filename", "output", trimmed)
	return output, nil
}

func (g *generator) resolveGlobs() ([]string, error) {
	slog.Debug("resolving include glob patterns...", "patterns", g.include)
	srcs := []string{g.input}
	for _, inc := range g.include {
		paths, err := filepath.Glob(inc)
		if err != nil {
			return nil, fmt.Errorf("failed to parse include pattern: %w", err)
		}

		srcs = append(srcs, paths...)
	}

	return srcs, nil
}

func (g *generator) renderDefault(srcs []string, data any, out string) error {
	tmpl, err := g.loadSources(srcs)
	if err != nil {
		return err
	}

	slog.Debug("creating output file...", "output", g.out)
	f, err := os.Create(out)
	if err != nil {
		return errOpenFile(g.out, err)
	}

	defer f.Close()

	slog.Debug("rendering template...")
	if err = tmpl.ExecuteTemplate(f, g.input, data); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}

func (g *generator) renderGo(srcs []string, data any, out string) error {
	tmpl, err := g.loadSources(srcs)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	slog.Debug("rendering template...")
	if err = tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	slog.Debug("formatting Go code...")
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format generated Go code: %w", err)
	}

	slog.Debug("creating output file...", "output", g.out)
	f, err := os.Create(out)
	if err != nil {
		return errOpenFile(g.out, err)
	}

	defer f.Close()

	slog.Debug("writing output file...")
	if _, err = f.Write(formatted); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}

func (g *generator) loadSources(srcs []string) (*template.Template, error) {
	slog.Debug("parsing template sources...", "srcs", srcs)
	tmpl, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(srcs...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	slog.Info("loaded templates", "srcs", srcs)
	return tmpl.Lookup(filepath.Base(g.input)), nil
}

func errOpenFile(name string, err error) error {
	return fmt.Errorf("failed to open file %q: %w", name, err)
}
