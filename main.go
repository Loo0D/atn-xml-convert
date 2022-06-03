package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"hash/crc32"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	tomlFile := flag.String("f", "", "name of TOML ballistic settings file")
	profilesFile := flag.String("o", "profiles.xml", "path and name of output profiles.xml file")
	flag.Parse()

	if *tomlFile == "" {
		_, _ = fmt.Fprintln(os.Stderr, "please provide filename of TOML settings file with -f flag")
	}

	f, err := os.Open(*tomlFile)
	if err != nil {
		log.Fatalln(err)
	}

	tomlContent, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	var tomlSettings TOMLSettings

	_, err = toml.Decode(string(tomlContent), &tomlSettings)
	if err != nil {
		log.Fatalln(err)
	}

	if tomlSettings.VelocityFps == true {
		convertFPS(&tomlSettings)
	}

	if tomlSettings.SightHeightIn == true {
		convertInches(&tomlSettings)
	}

	if tomlSettings.ZeroDistanceYds == true {
		convertYards(&tomlSettings)
	}

	t, err := template.New("xml").Parse(XMLTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	var xml strings.Builder

	err = t.ExecuteTemplate(&xml, "xml", tomlSettings)
	if err != nil {
		log.Fatalln(err)
	}

	xmlString := xml.String()

	xmlString = strings.ReplaceAll(xmlString, "\r\n", "")
	xmlString = strings.ReplaceAll(xmlString, "\n", "")
	xmlString = xmlString + "\n"

	crc := crc32.ChecksumIEEE([]byte(xmlString))

	crcString := strings.ToUpper(fmt.Sprintf("%08x", crc))

	commentLine := "<!--" + crcString + "-->" + "\r\n"

	f, err = os.Create(*profilesFile)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = f.WriteString(commentLine + xmlString)
	if err != nil {
		log.Fatalln(err)
	}

}

func convertYards(settings *TOMLSettings) {
	yards := settings.One.ZeroingDistance
	settings.One.ZeroingDistance = yards * 0.9144
	yards = settings.Two.ZeroingDistance
	settings.Two.ZeroingDistance = yards * 0.9144
	yards = settings.Three.ZeroingDistance
	settings.Three.ZeroingDistance = yards * 0.9144
	yards = settings.Four.ZeroingDistance
	settings.Four.ZeroingDistance = yards * 0.9144
	yards = settings.Five.ZeroingDistance
	settings.Five.ZeroingDistance = yards * 0.9144
	yards = settings.Six.ZeroingDistance
	settings.Six.ZeroingDistance = yards * 0.9144
}

func convertInches(settings *TOMLSettings) {
	inches := settings.One.SightHeight
	settings.One.SightHeight = inches * 25.4
	inches = settings.Two.SightHeight
	settings.Two.SightHeight = inches * 25.4
	inches = settings.Three.SightHeight
	settings.Three.SightHeight = inches * 25.4
	inches = settings.Four.SightHeight
	settings.Four.SightHeight = inches * 25.4
	inches = settings.Five.SightHeight
	settings.Five.SightHeight = inches * 25.4
	inches = settings.Six.SightHeight
	settings.Six.SightHeight = inches * 25.4
}

func convertFPS(settings *TOMLSettings) {
	fps := settings.One.InitVelocity
	settings.One.InitVelocity = fps / 3.28084
	fps = settings.Two.InitVelocity
	settings.Two.InitVelocity = fps / 3.28084
	fps = settings.Three.InitVelocity
	settings.Three.InitVelocity = fps / 3.28084
	fps = settings.Four.InitVelocity
	settings.Four.InitVelocity = fps / 3.28084
	fps = settings.Five.InitVelocity
	settings.Five.InitVelocity = fps / 3.28084
	fps = settings.Six.InitVelocity
	settings.Six.InitVelocity = fps / 3.28084
}
