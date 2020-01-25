package world

import (
	"testing"
)

const numberOfAliens = 4

func TestWorld(t *testing.T) {
	cities := initCities()
	cityMap := GenerateMap(cities)

	validateCitiesCreation(cityMap, t)
	validateCitiesCommunications(cityMap, t)
	validateAliensCreation(cityMap, t)
	validateSimulation(cityMap, t)
}

func validateCitiesCreation(cityMap *CitiesMap, t *testing.T) {
	if len(cityMap.Cities) != 11 {
		t.Error("Expected the map to contain exactly 10 cities")
	}
}

func validateCitiesCommunications(cityMap *CitiesMap, t *testing.T) {
	if len(cityMap.Cities["Madrid"].AdjCities) != 4 {
		t.Error("Expected Madrid to be communicated exactly with 4 cities")
	}

	if len(cityMap.Cities["Bilbao"].AdjCities) != 3 {
		t.Error("Expected Bilbao to be communicated exactly with 2 cities")
	}

	if len(cityMap.Cities["Sevilla"].AdjCities) != 1 {
		t.Error("Expected Sevilla to be communicated exactly with 1 city")
	}

	if len(cityMap.Cities["Lisboa"].AdjCities) != 1 {
		t.Error("Expected Lisboa to be communicated exactly with 1 city")
	}

	if len(cityMap.Cities["Huesca"].AdjCities) != 1 {
		t.Error("Expected Huesca to be communicated exactly with 1 city")
	}

	if len(cityMap.Cities["Valencia"].AdjCities) != 2 {
		t.Error("Expected Valencia to be communicated exactly with 2 cities")
	}

	if len(cityMap.Cities["Zaragoza"].AdjCities) != 3 {
		t.Error("Expected Zaragoza to be communicated exactly with 3 cities")
	}

	if len(cityMap.Cities["Santander"].AdjCities) != 2 {
		t.Error("Expected Santander to be communicated exactly with 2cities")
	}

	if len(cityMap.Cities["Barcelona"].AdjCities) != 2 {
		t.Error("Expected Barcelona to be communicated exactly with 2 cities")
	}

	if len(cityMap.Cities["Oviedo"].AdjCities) != 2 {
		t.Error("Expected Oviedo to be communicated exactly with 2 cities")
	}

	if len(cityMap.Cities["Coruna"].AdjCities) != 1 {
		t.Error("Expected Oviedo to be communicated exactly with 1 city")
	}
}

func validateAliensCreation(cityMap *CitiesMap, t *testing.T) {
	cityMap.GenerateAliens(numberOfAliens)

	alienCount := 0
	for _, city := range cityMap.Cities {
		alienCount += len(city.Aliens)
	}

	if alienCount != numberOfAliens {
		t.Error("Expected the number of aliens generated to be exactly " + string(numberOfAliens))
	}
}

func validateSimulation(cityMap *CitiesMap, t *testing.T) {
	cityMap.StartSimulation()

	// The simulation ends correctly when we reach on of the following states:
	// 1. All aliens destroyed or trapped
	// 2. Each alien moved 10000 times

	simulationOk := true
	for _, city := range cityMap.Cities {
		// If we reached an invalid state, we stop checking
		if simulationOk == false {
			break
		}
		// No aliens in the current city
		if len(city.Aliens) == 0 {
			continue
		}

		for _, alien := range city.Aliens {
			// If the alien is not trapped and didn't move 10000 times, we reached an invalid state
			if !alien.Trapped && alien.Steps != 10000 {
				simulationOk = false
				break
			}
		}
	}

	if simulationOk == false {
		t.Error("Expected the simulation to finish in a valid state")
	}
}

func initCities() map[string]map[string]string {
	var cities = map[string]map[string]string{}
	cities = make(map[string]map[string]string)

	cities["Madrid"] = make(map[string]string)
	cities["Madrid"]["north"] = "Bilbao"
	cities["Madrid"]["south"] = "Sevilla"
	cities["Madrid"]["east"] = "Lisboa"
	cities["Madrid"]["west"] = "Valencia"

	cities["Bilbao"] = make(map[string]string)
	cities["Bilbao"]["east"] = "Zaragoza"
	cities["Bilbao"]["west"] = "Santander"

	cities["Zaragoza"] = make(map[string]string)
	cities["Zaragoza"]["east"] = "Barcelona"
	cities["Zaragoza"]["north"] = "Huesca"

	cities["Santander"] = make(map[string]string)
	cities["Santander"]["west"] = "Oviedo"

	cities["Barcelona"] = make(map[string]string)
	cities["Barcelona"]["south"] = "Valencia"

	cities["Oviedo"] = make(map[string]string)
	cities["Oviedo"]["west"] = "Coruna"

	return cities
}
