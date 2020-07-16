package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

var version string

func main() {
	version = "0.1.0"

	args := os.Args[1:]

	if (len(args) < 1) {
		showUsage()
		os.Exit(1)
	}

	for _, arg := range args {
		if (arg == "--help" || arg == "-h") {
			showUsage()
			os.Exit(1)
		}

		if (arg == "--version" || arg == "-v") {
			showVersion()
			os.Exit(1)
		}

		fmt.Println(arg)
	}

	srcVer := strings.Split(os.Args[1], ".")
	retVer := srcVer
	incFlag := os.Args[2]

	fmt.Println("srcVer: " + srcVer[0] + " . " + srcVer[1] + " . " + srcVer[2])
	fmt.Println("incFlag: " + incFlag)

	retVer = bumpMajor(srcVer)

	fmt.Println(strings.Join(retVer, "."))
}

func bumpMajor(incVer []string) []string {
	fmt.Println(incVer)
	var retVer []string

	incVal, _ := strconv.Atoi(incVer[0])
	retVer = append(retVer, strconv.Itoa(incVal + 1))
	retVer = append(retVer, "0")
	retVer = append(retVer, "0")

	return retVer
}

func showVersion() {
	fmt.Println("https://github.inbcu.com/d2c-integrations/vinc")
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