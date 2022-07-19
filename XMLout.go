package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"hash/crc32"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func tomlToXml(tomlFile *string, profilesFile *string) {
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

	convertSpecialChars(&tomlSettings)

	// profiles.xml file requires meters/second
	if tomlSettings.VelocityFps == true {
		convertFPStoMS(&tomlSettings)
	}

	// profiles.xml file requires mm
	if tomlSettings.SightHeightIn == true {
		convertInchesToMM(&tomlSettings)
	}

	// profiles.xml file requires meters
	if tomlSettings.ZeroDistanceYds == true {
		convertYardsToMeters(&tomlSettings)
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
