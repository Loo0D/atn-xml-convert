package main

import (
	"flag"
	"fmt"
	"os"
)

var tomlZeroDistanceYards *bool
var tomlVelocityFps *bool
var tomlSightHeightInches *bool

func main() {
	inputTomlFileName := flag.String("t", "", "name of TOML ballistic settings file")
	outputFileName := flag.String("o", "", "path and name of output profiles.xml file")
	inputXmlFileName := flag.String("x", "", "path and name of input profiles.xml file")
	tomlZeroDistanceYards = flag.Bool("yards", false, "TOML output zeroingDistance in yards instead of meters")
	tomlVelocityFps = flag.Bool("fps", false, "TOML output initVelocity in fps instead of m/s")
	tomlSightHeightInches = flag.Bool("inches", false, "TOML output sightHeight in inches instead of mm")

	flag.Parse()

	if *inputTomlFileName == "" && *inputXmlFileName == "" {
		_, _ = fmt.Fprintln(os.Stderr, "please provide filename of TOML settings file (-t flag) or input XML file (-x flag)")
	}

	if *inputTomlFileName != "" {
		if *outputFileName == "" {
			*outputFileName = "profiles.xml"
		}
		tomlToXml(inputTomlFileName, outputFileName)
	}

	if *inputXmlFileName != "" {
		if *outputFileName == "" {
			*outputFileName = "profiles.toml"
		}
		xmlToToml(inputXmlFileName, outputFileName)
	}

}
