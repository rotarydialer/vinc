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

	if (len(args) < 2) {
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
	}

	srcVer := strings.Split(os.Args[1], ".")
	retVer := srcVer
	incFlag := os.Args[2]

	switch incFlag {
	case "--major":
		retVer = bumpMajor(srcVer)
		// fmt.Println("Major version: " + strings.Join(retVer, "."))
		break
	case "--minor":
		retVer = bumpMinor(srcVer)
		// fmt.Println("Minor version: " + strings.Join(retVer, "."))
		break
	case "--patch":
		retVer = bumpPatch(srcVer)
		// fmt.Println("Patch version: " + strings.Join(retVer, "."))
		break
	}

	fmt.Println(strings.Join(retVer, "."))
}

func bumpMajor(incVer []string) []string {
	var retVer []string

	incVal, _ := strconv.Atoi(incVer[0])
	retVer = append(retVer, strconv.Itoa(incVal + 1))
	retVer = append(retVer, "0")
	retVer = append(retVer, "0")

	return retVer
}

func bumpMinor(incVer []string) []string {
	var retVer []string

	incVal, _ := strconv.Atoi(incVer[1])
	retVer = append(retVer, incVer[0])
	retVer = append(retVer, strconv.Itoa(incVal + 1))
	retVer = append(retVer, "0")

	return retVer
}

func bumpPatch(incVer []string) []string {
	var retVer []string

	incVal, _ := strconv.Atoi(incVer[2])
	retVer = append(retVer, incVer[0])
	retVer = append(retVer, incVer[1])
	retVer = append(retVer, strconv.Itoa(incVal + 1))

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