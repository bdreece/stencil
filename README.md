# stencil

A command-line tool for generic code generation using Go template files.

## Installation

```sh
go install github.com/bdreece/stencil@latest
```

or for Go projects:

```sh
go get -tool github.com/bdreece/stencil@latest
```

## Usage

```

      |                  _) |
  __| __|  _ \ __ \   __| | |
\__ \ |    __/ |   | (    | |
____/\__|\___|_|  _|\___|_|_|

stencil is a command-line tool for generic code generation using Go template files.

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

```

## Examples

See the [examples/](./examples/) directory for more examples.

---

MIT License

Copyright (c) 2025 Brian Reece
