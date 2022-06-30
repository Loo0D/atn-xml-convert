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

/*
Replaces special characters in profile names with equivalent XML entities
&amp;  - The ampersand character (&) starts entity markup (the first character of a character entity reference).
&lt;   - The less-than character (<) starts element markup (the first character of a start-tag or an end-tag).
&gt;   - The greater-than character (>) ends a start-tag or an end-tag.
&quot; - The double-quote character (") can be symbolised with this character entity reference when you need to embed a double-quote inside a string which is already double-quoted.
&apos; - The apostrophe or single-quote character (') can be symbolised with this character entity reference when you need to embed a single-quote or apostrophe inside a string which is already single-quoted.
*/
func convertSpecialChars(settings *TOMLSettings) {
	settings.One.ProfileName = replaceChars(settings.One.ProfileName)
	settings.Two.ProfileName = replaceChars(settings.Two.ProfileName)
	settings.Three.ProfileName = replaceChars(settings.Three.ProfileName)
	settings.Four.ProfileName = replaceChars(settings.Four.ProfileName)
	settings.Five.ProfileName = replaceChars(settings.Five.ProfileName)
	settings.Six.ProfileName = replaceChars(settings.Six.ProfileName)
}

func replaceChars(profileName string) string {
	profileName = strings.ReplaceAll(profileName, "&", "&amp;")
	profileName = strings.ReplaceAll(profileName, "<", "&lt;")
	profileName = strings.ReplaceAll(profileName, ">", "&gt;")
	profileName = strings.ReplaceAll(profileName, "\"", "&quot;")
	profileName = strings.ReplaceAll(profileName, "'", "&apos;")
	return profileName
}

func convertYardsToMeters(settings *TOMLSettings) {
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

func convertInchesToMM(settings *TOMLSettings) {
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

func convertFPStoMS(settings *TOMLSettings) {
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
