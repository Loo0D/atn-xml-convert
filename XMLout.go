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

	insertDefaultsIfEmpty(&tomlSettings)

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

// If "unknown" settings are not present, insert defaults
func insertDefaultsIfEmpty(t *TOMLSettings) {

	var msp_acc_border_value = 5000
	var acc_border_up_cross_counter_min = 4
	var msp_acc_border_up_cross_counter_max = 30
	var msp_acc_border_down_cross_counter_min = 1

	if t.One.MspAccBorderValue == 0 {
		t.One.MspAccBorderValue = msp_acc_border_value
	}

	if t.One.AccBorderUpCrossCounterMin == 0 {
		t.One.AccBorderUpCrossCounterMin = acc_border_up_cross_counter_min
	}

	if t.One.MspAccBorderUpCrossCounterMax == 0 {
		t.One.MspAccBorderUpCrossCounterMax = msp_acc_border_up_cross_counter_max
	}

	if t.One.MspAccBorderDownCrossCounterMin == 0 {
		t.One.MspAccBorderDownCrossCounterMin = msp_acc_border_down_cross_counter_min
	}

	if t.Two.MspAccBorderValue == 0 {
		t.Two.MspAccBorderValue = msp_acc_border_value
	}

	if t.Two.AccBorderUpCrossCounterMin == 0 {
		t.Two.AccBorderUpCrossCounterMin = acc_border_up_cross_counter_min
	}

	if t.Two.MspAccBorderUpCrossCounterMax == 0 {
		t.Two.MspAccBorderUpCrossCounterMax = msp_acc_border_up_cross_counter_max
	}

	if t.Two.MspAccBorderDownCrossCounterMin == 0 {
		t.Two.MspAccBorderDownCrossCounterMin = msp_acc_border_down_cross_counter_min
	}

	if t.Three.MspAccBorderValue == 0 {
		t.Three.MspAccBorderValue = msp_acc_border_value
	}

	if t.Three.AccBorderUpCrossCounterMin == 0 {
		t.Three.AccBorderUpCrossCounterMin = acc_border_up_cross_counter_min
	}

	if t.Three.MspAccBorderUpCrossCounterMax == 0 {
		t.Three.MspAccBorderUpCrossCounterMax = msp_acc_border_up_cross_counter_max
	}

	if t.Three.MspAccBorderDownCrossCounterMin == 0 {
		t.Three.MspAccBorderDownCrossCounterMin = msp_acc_border_down_cross_counter_min
	}

	if t.Four.MspAccBorderValue == 0 {
		t.Four.MspAccBorderValue = msp_acc_border_value
	}

	if t.Four.AccBorderUpCrossCounterMin == 0 {
		t.Four.AccBorderUpCrossCounterMin = acc_border_up_cross_counter_min
	}

	if t.Four.MspAccBorderUpCrossCounterMax == 0 {
		t.Four.MspAccBorderUpCrossCounterMax = msp_acc_border_up_cross_counter_max
	}

	if t.Four.MspAccBorderDownCrossCounterMin == 0 {
		t.Four.MspAccBorderDownCrossCounterMin = msp_acc_border_down_cross_counter_min
	}

	if t.Five.MspAccBorderValue == 0 {
		t.Five.MspAccBorderValue = msp_acc_border_value
	}

	if t.Five.AccBorderUpCrossCounterMin == 0 {
		t.Five.AccBorderUpCrossCounterMin = acc_border_up_cross_counter_min
	}

	if t.Five.MspAccBorderUpCrossCounterMax == 0 {
		t.Five.MspAccBorderUpCrossCounterMax = msp_acc_border_up_cross_counter_max
	}

	if t.Five.MspAccBorderDownCrossCounterMin == 0 {
		t.Five.MspAccBorderDownCrossCounterMin = msp_acc_border_down_cross_counter_min
	}

	if t.Six.MspAccBorderValue == 0 {
		t.Six.MspAccBorderValue = msp_acc_border_value
	}

	if t.Six.AccBorderUpCrossCounterMin == 0 {
		t.Six.AccBorderUpCrossCounterMin = acc_border_up_cross_counter_min
	}

	if t.Six.MspAccBorderUpCrossCounterMax == 0 {
		t.Six.MspAccBorderUpCrossCounterMax = msp_acc_border_up_cross_counter_max
	}

	if t.Six.MspAccBorderDownCrossCounterMin == 0 {
		t.Six.MspAccBorderDownCrossCounterMin = msp_acc_border_down_cross_counter_min
	}

}
