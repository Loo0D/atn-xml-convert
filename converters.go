package main

import "strings"

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

func convertMetersToYards(settings *TOMLSettings) {
	meters := settings.One.ZeroingDistance
	settings.One.ZeroingDistance = meters * 1.09361
	meters = settings.Two.ZeroingDistance
	settings.Two.ZeroingDistance = meters * 1.09361
	meters = settings.Three.ZeroingDistance
	settings.Three.ZeroingDistance = meters * 1.09361
	meters = settings.Four.ZeroingDistance
	settings.Four.ZeroingDistance = meters * 1.09361
	meters = settings.Five.ZeroingDistance
	settings.Five.ZeroingDistance = meters * 1.09361
	meters = settings.Six.ZeroingDistance
	settings.Six.ZeroingDistance = meters * 1.09361
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

func convertMMtoInches(settings *TOMLSettings) {
	mm := settings.One.SightHeight
	settings.One.SightHeight = mm / 25.4
	mm = settings.Two.SightHeight
	settings.Two.SightHeight = mm / 25.4
	mm = settings.Three.SightHeight
	settings.Three.SightHeight = mm / 25.4
	mm = settings.Four.SightHeight
	settings.Four.SightHeight = mm / 25.4
	mm = settings.Five.SightHeight
	settings.Five.SightHeight = mm / 25.4
	mm = settings.Six.SightHeight
	settings.Six.SightHeight = mm / 25.4
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

func convertMStoFPS(settings *TOMLSettings) {
	ms := settings.One.InitVelocity
	settings.One.InitVelocity = ms * 3.28084
	ms = settings.Two.InitVelocity
	settings.Two.InitVelocity = ms * 3.28084
	ms = settings.Three.InitVelocity
	settings.Three.InitVelocity = ms * 3.28084
	ms = settings.Four.InitVelocity
	settings.Four.InitVelocity = ms * 3.28084
	ms = settings.Five.InitVelocity
	settings.Five.InitVelocity = ms * 3.28084
	ms = settings.Six.InitVelocity
	settings.Six.InitVelocity = ms * 3.28084
}
