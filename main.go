package main

import (
	"io/ioutil"
	"log"
	utils "main/utils"
	"main/validators"
	"main/world"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Missing command line arguments")
	}

	mapFile, numberOfAliens := validators.ValidateAndGetArgs(ioutil.ReadFile, os.Args[1], os.Args[2])
	cities := utils.GetCitiesFromMapFile(mapFile)

	citiesMap := world.GenerateMap(cities)
	citiesMap.GenerateAliens(numberOfAliens)
	citiesMap.StartSimulation()
}
