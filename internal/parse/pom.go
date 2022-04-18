/*
Copyright © 2022 Sumeet Patil sumeet.patil@sap.com

*/

package parse

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Project struct {
	XMLName      xml.Name     `xml:"project"`
	Dependencies []Dependency `xml:"dependencies>dependency"`
	Properties   Entry        `xml:"properties"`
}

type Entry struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

type Dependency struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func ParsePomDependencies(pomPath string) []Library {
	pomDependencies := []Library{}
	xmlFile, err := os.Open(pomPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var project Project

	xml.Unmarshal(byteValue, &project)

	for _, dependency := range project.Dependencies {
		var lib Library
		lib.LibraryName = dependency.GroupID + ":" + dependency.ArtifactID
		lib.LibraryType = "Maven"
		if strings.Contains(dependency.Version, "${") {
			dependency.Version = strings.ReplaceAll(dependency.Version, "${", "")
			dependency.Version = strings.ReplaceAll(dependency.Version, "}", "")
			depVersion, isGetStrTrue := getStringInBetweenTwoString(project.Properties.Value, "<"+dependency.Version+">", "</"+dependency.Version+">")
			if isGetStrTrue {
				dependency.Version = depVersion
			} else {
				continue
			}
		}
		lib.Version = dependency.Version
		pomDependencies = append(pomDependencies, lib)
	}

	return pomDependencies
}

func getStringInBetweenTwoString(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}
