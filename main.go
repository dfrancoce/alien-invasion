package main

import (
	utils "main/utils"
	"main/validators"
	"main/world"
)

func main() {
	mapFile, numberOfAliens := validators.ValidateAndGetArgs()
	cities := utils.GetCitiesFromMapFile(mapFile)

	citiesMap := world.GenerateMap(cities)
	citiesMap.GenerateAliens(numberOfAliens)
	citiesMap.StartSimulation()
}
