# Diablo 2 Save editor CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/vitalick/d2editor-cli)](https://goreportcard.com/report/github.com/vitalick/d2editor-cli)
[![GoDoc](https://godoc.org/github.com/vitalick/d2editor-cli?status.svg)](https://godoc.org/github.com/vitalick/d2editor-cli)

CLI for convert .d2s to JSON and vice-versa

## Installation

To install command line program, use the following:

```bash
go install github.com/vitalick/d2s/d2editor-cli@latest
```

## Usage
### CLI

For convert JSON to .d2s, use the following:
```bash
d2editor-cli -tod2s <input files>
```

For convert .d2s to JSON, use the following:
```bash
d2editor-cli -tojson <input files>
```

For create empty .d2s or JSON, use the following:
```bash
d2editor-cli -e -tod2s <input names>
d2editor-cli -e -tojson <input names>
```

To specify the path to the output folder, we use the following:
```bash
d2editor-cli -tod2s -o <output folder> <input files>
d2editor-cli -tojson -o <output folder> <input files>
```

To specify the version of save, we use the following:
```bash
d2editor-cli -tod2s -v <output version> <input files>
```

## Links

- https://github.com/vitalick/go-d2editor
