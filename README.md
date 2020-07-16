# vinc
`vinc` (`v`ersion `inc`rement) is a simple semantic versioning utility written in [Go](https://golang.org/). It follows the standard set forth in [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html)

We created this because we struggled to find a Maven plugin for Compass that semantically versioned correctly. This utility is meant to function similar to [the `semver` npm module](https://www.npmjs.com/package/semver).

## Prerequisites
- [Go v1.14](https://golang.org/dl/)

## Build

### For local use:
```
   go build vinc.go
```
This will produce a portable, standalone binary, `vinc`, appropriate for use your local terminal.

### For automation use:
```
env GOOS=linux GOARCH=amd64 go build -o vinc vinc.go
```

This will produce a binary, `vinc`, appropriate for use in most automation pipelines (Linux/AMD64-based).

## Usage

`vinc` accepts a three-digit semantic version number, followed by a flag indicating which digit to increment: `--major`, `--minor` or `--patch`. It will return an incremented version number.

### Examples

#### Terminal
```
   $ ./vinc 1.2.3 --major
   2.0.0

   $ ./vinc 1.2.3 --minor
   1.3.0

   $ ./vinc 1.2.3 --patch
   1.2.4
```

#### Shell Script / Set Variable
```
    $ NEW_VERSION=$(./vinc 10.20.1 --minor)
    $ echo $NEW_VERSION
    10.21.0
```

### Other Options

| Flag | Details |
--- | --- |
| `--help, -h` | Displays usage/help |
| `--version, -v` | Displays vinc version |