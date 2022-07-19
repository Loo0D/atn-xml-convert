package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/beevik/etree"
	"log"
	"os"
	"strconv"
)

func xmlToToml(xmlFileName *string, tomlFileName *string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(*xmlFileName); err != nil {
		log.Fatalln(err)
	}

	// Determine number of profiles in profiles.xml file
	count, err := strconv.Atoi(doc.SelectElement("profiles_count").ChildElements()[0].Text())
	if err != nil {
		log.Fatalf("error parsing profiles_count: %v", err)
	}

	var profiles []Profile
	for i := 0; i < count; i++ {
		profiles = append(profiles, parseXmlRecord(doc, i))
	}

	var settings TOMLSettings
	populateTomlSettings(&settings, profiles)

	settings.ZeroDistanceYds = *tomlZeroDistanceYards
	settings.SightHeightIn = *tomlSightHeightInches
	settings.VelocityFps = *tomlVelocityFps

	if settings.VelocityFps == true {
		convertMStoFPS(&settings)
	}
	if settings.SightHeightIn == true {
		convertMMtoInches(&settings)
	}
	if settings.ZeroDistanceYds == true {
		convertMetersToYards(&settings)
	}

	writeTomlFile(&settings, *tomlFileName)

	fmt.Printf("TOML file created: %v\n[!] ENSURE THE UNITS ARE WHAT YOU EXPECT (fps / meters / yards / etc)\n", *tomlFileName)
}

func populateTomlSettings(settings *TOMLSettings, profiles []Profile) {
	if len(profiles) > 0 {
		settings.One = profiles[0]
	}

	if len(profiles) > 1 {
		settings.Two = profiles[1]
	}

	if len(profiles) > 2 {
		settings.Three = profiles[2]
	}

	if len(profiles) > 3 {
		settings.Four = profiles[3]
	}

	if len(profiles) > 4 {
		settings.Five = profiles[4]
	}

	if len(profiles) > 5 {
		settings.Six = profiles[5]
	}
}

func parseXmlRecord(doc *etree.Document, id int) Profile {
	var profile Profile
	var err error

	profileName := doc.FindElement("//storage_profile_name/value_node_" + strconv.Itoa(id))
	profile.ProfileName = profileName.Text()

	dragFunction := doc.FindElement("//storage_drag_function/value_node_" + strconv.Itoa(id))
	profile.DragFunction, err = strconv.Atoi(dragFunction.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (drag_function): %v\n", id, err)
	}

	bc := doc.FindElement("//storage_ballistic_coeff/value_node_" + strconv.Itoa(id))
	profile.Bc, err = strconv.ParseFloat(bc.Text(), 64)
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (ballistic_coeff): %v\n", id, err)
	}

	bulletWeight := doc.FindElement("//storage_bullet_weight/value_node_" + strconv.Itoa(id))
	profile.BulletWeight, err = strconv.ParseFloat(bulletWeight.Text(), 64)
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (bullet_weight): %v\n", id, err)
	}

	initVelocity := doc.FindElement("//storage_init_velocity/value_node_" + strconv.Itoa(id))
	profile.InitVelocity, err = strconv.ParseFloat(initVelocity.Text(), 64)
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (init_velocity): %v\n", id, err)
	}

	sightHeight := doc.FindElement("//storage_sight_height/value_node_" + strconv.Itoa(id))
	profile.SightHeight, err = strconv.ParseFloat(sightHeight.Text(), 64)
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (sight_height): %v\n", id, err)
	}

	zeroingDistance := doc.FindElement("//storage_zeroing_distance/value_node_" + strconv.Itoa(id))
	profile.ZeroingDistance, err = strconv.ParseFloat(zeroingDistance.Text(), 64)
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (zeroing_distance): %v\n", id, err)
	}

	reticleOffsetX := doc.FindElement("//storage_reticle_offset_x/value_node_" + strconv.Itoa(id))
	profile.ReticleOffsetX, err = strconv.Atoi(reticleOffsetX.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (reticle_offset_x): %v\n", id, err)
	}

	reticleOffsetY := doc.FindElement("//storage_reticle_offset_y/value_node_" + strconv.Itoa(id))
	profile.ReticleOffsetY, err = strconv.Atoi(reticleOffsetY.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (reticle_offset_y): %v\n", id, err)
	}

	mspAccBorderValue := doc.FindElement("//storage_msp_acc_border_value/value_node_" + strconv.Itoa(id))
	profile.MspAccBorderValue, err = strconv.Atoi(mspAccBorderValue.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (msp_acc_border_value): %v\n", id, err)
	}

	mspAccBorderUpCrossCounterMin := doc.FindElement("//storage_msp_acc_border_up_cross_counter_min/value_node_" + strconv.Itoa(id))
	profile.AccBorderUpCrossCounterMin, err = strconv.Atoi(mspAccBorderUpCrossCounterMin.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (msp_acc_border_up_cross_counter_min): %v\n", id, err)
	}

	mspAccBorderUpCrossCounterMax := doc.FindElement("//storage_msp_acc_border_up_cross_counter_max/value_node_" + strconv.Itoa(id))
	profile.MspAccBorderUpCrossCounterMax, err = strconv.Atoi(mspAccBorderUpCrossCounterMax.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (msp_acc_border_up_cross_counter_max): %v\n", id, err)
	}

	mspAccBorderDownCrossCounterMin := doc.FindElement("//storage_msp_acc_border_down_cross_counter_min/value_node_" + strconv.Itoa(id))
	profile.MspAccBorderDownCrossCounterMin, err = strconv.Atoi(mspAccBorderDownCrossCounterMin.Text())
	if err != nil {
		log.Fatalf("Failed to parse XML value for profile %d (msp_acc_border_down_cross_counter_min): %v\n", id, err)
	}

	return profile
}

func writeTomlFile(settings *TOMLSettings, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v\n", fileName, err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(settings)
	if err != nil {
		log.Fatalf("Failed to encode TOML file: %v\n", err)
	}
}
