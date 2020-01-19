package _utils

import (
	"strings"
)

func GetCitiesFromMapFile(mapFile []byte) map[string]map[string]string {
	var cities = map[string]map[string]string{}
	cities = make(map[string]map[string]string)
	mapFileLines := getMapFileLines(string(mapFile))

	for _, line := range mapFileLines {
		if line == "" { continue }
		city, relCities := getCitiesFromMapFileLine(line)
		cities[city] = relCities
	}

	return cities
}

func getMapFileLines(mapFile string) []string {
	return strings.Split(mapFile, "\r\n")
}

func getCitiesFromMapFileLine(mapFileLine string) (string, map[string]string) {
	cities := strings.Split(mapFileLine, " ")
	mainCity := cities[0]

	var relCities map[string]string
	relCities = make(map[string]string)

	for i := 1; i < len(cities); i++ {
		citiesAndDirections := strings.Split(cities[i], "=")
		relCities[strings.ToLower(citiesAndDirections[0])] = citiesAndDirections[1]
	}

	return mainCity, relCities
}
