package _utils

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func GetCitiesFromMapFile(mapFile []byte) map[string]map[string]string {
	var cities = map[string]map[string]string{}
	cities = make(map[string]map[string]string)
	mapFileLines := getMapFileLines(string(mapFile))

	for _, line := range mapFileLines {
		if line == "" {
			continue
		}

		city, relCities, err := getCitiesFromMapFileLine(line)
		if err == nil {
			cities[city] = relCities
		}
	}

	return cities
}

func getMapFileLines(mapFile string) []string {
	mapFileReplaceReturn := strings.ReplaceAll(mapFile, "\r\n", "\n")
	return strings.Split(mapFileReplaceReturn, "\n")
}

func getCitiesFromMapFileLine(mapFileLine string) (string, map[string]string, error) {
	matched, err := regexp.MatchString(`(\A[a-zA-Z]+)(\s(north|south|west|east)=[a-zA-Z]+)+`, mapFileLine)

	if matched == false || err != nil {
		return "", nil, getInputFileLineError("The format of the line " + mapFileLine + " is not correct. It won't be included.")
	}

	cities := strings.Split(mapFileLine, " ")
	mainCity, relCities := getMainCityAndRelCities(cities)

	for i := 1; i < len(cities); i++ {
		citiesAndDirections := strings.Split(cities[i], "=")

		if _, alreadyContainsDirection := relCities[citiesAndDirections[0]]; alreadyContainsDirection {
			return "", nil, getInputFileLineError("The direction " + citiesAndDirections[0] + " is already in use")
		}

		relCities[citiesAndDirections[0]] = citiesAndDirections[1]
	}

	return mainCity, relCities, nil
}

func getMainCityAndRelCities(cities []string) (string, map[string]string) {
	mainCity := cities[0]
	var relCities map[string]string
	relCities = make(map[string]string)

	return mainCity, relCities
}

func getInputFileLineError(errorMessage string) error {
	err := errors.New(errorMessage)
	log.Printf("An error occurred while parsing the input file: %s\n", err)

	return err
}
