package main

import (
	"os"
	"fmt"
)

var version string

func main() {
	version = "0.1.0"

	args := os.Args[1:]

	if (len(args) < 1) {
		showUsage()
	}

	for _, arg := range args {
		if (arg == "--help" || arg == "-h") {
			showUsage()
			break
		}

		if (arg == "--version" || arg == "-v") {
			showVersion()
			break
		}

		fmt.Println(arg)
	}
}

func showVersion() {
	fmt.Println(version)
}

func showUsage() {
	fmt.Println("usage: ./vinc <semantic version number> [--major | --minor | --patch]")
	fmt.Println("              [--version]  [--help | -h]")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("   ./vinc 1.5.0 --patch")
	fmt.Println("   1.5.1")
	fmt.Println()
	fmt.Println("   ./vinc 10.5.6 --minor")
	fmt.Println("   10.6.0")
	fmt.Println()
	fmt.Println("   ./vinc 8.7.24 --major")
	fmt.Println("   9.0.0")
}