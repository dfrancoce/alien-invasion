package world

import (
	"fmt"
	"main/aliens"
	"math/rand"
	"time"
)

type CitiesMap struct {
	Cities map[string]*City
}

type City struct {
	Name      string
	AdjCities map[direction]*City
	Aliens    []aliens.Alien
}

var numberOfAliensAliveAndNotTrapped int

func GenerateMap(cities map[string]map[string]string) *CitiesMap {
	fmt.Println("Generating map... ")

	cityMap := CitiesMap{Cities: map[string]*City{}}

	for cityName, relCities := range cities {
		city := cityMap.addCity(cityName)
		for direction, relCity := range relCities {
			relCity := cityMap.addCity(relCity)
			cityMap.addRoad(city, relCity, direction)
		}
	}

	return &cityMap
}

func (cityMap *CitiesMap) GenerateAliens(numberOfAliens int) {
	fmt.Printf("Generating %d aliens\n", numberOfAliens)

	numberOfAliensAliveAndNotTrapped = numberOfAliens
	cities := cityMap.getMapCities()
	for i := 0; i < numberOfAliens; i++ {
		randomCityName := getRandomCityName(cities)
		cityMap.Cities[randomCityName].Aliens = append(cityMap.Cities[randomCityName].Aliens, aliens.GenerateRandomAlien())
	}
}

func getRandomCityName(source []string) string {
	rand.Seed(time.Now().UnixNano())
	return source[rand.Intn(len(source))]
}

func (cityMap *CitiesMap) getMapCities() []string {
	cities := make([]string, 0, len(cityMap.Cities))
	for city := range cityMap.Cities {
		cities = append(cities, city)
	}

	return cities
}

func (cityMap *CitiesMap) StartSimulation() {
	alienSteps := 1

	for alienSteps <= 100000 && numberOfAliensAliveAndNotTrapped > 0 {
		fmt.Printf("Iteration: %d\n", alienSteps)
		fmt.Printf("numberOfAliensAliveAndNotTrapped: %d\n", numberOfAliensAliveAndNotTrapped)

		for i := range cityMap.Cities {
			currentCity := cityMap.Cities[i]
			currentCityName := currentCity.Name
			aliensInTheCity := len(currentCity.Aliens)

			if aliensInTheCity > 1 {
				cityMap.destroyCity(currentCityName)
			} else if aliensInTheCity == 1 && !currentCity.Aliens[0].Trapped {
				cityMap.moveAlien(currentCityName, alienSteps)
			}
		}

		alienSteps += 1
	}
}

func (cityMap *CitiesMap) addCity(name string) *City {
	if foundCity, ok := cityMap.Cities[name]; ok {
		return foundCity
	}

	city := City{Name: name, AdjCities: map[direction]*City{}, Aliens: []aliens.Alien{}}
	cityMap.Cities[name] = &city

	return &city
}

func (cityMap *CitiesMap) addRoad(fromCity *City, toCity *City, directionString string) {
	if _, ok := cityMap.Cities[fromCity.Name]; !ok {
		panic("No city named " + fromCity.Name)
	}
	if _, ok := cityMap.Cities[toCity.Name]; !ok {
		panic("No city named " + toCity.Name)
	}

	direction := convertToDirection(directionString)
	if _, ok := fromCity.AdjCities[direction]; !ok {
		fromCity.AdjCities[direction] = toCity
	}

	inverseDirection := inverseDirection(direction)
	if _, ok := toCity.AdjCities[inverseDirection]; !ok {
		toCity.AdjCities[inverseDirection] = fromCity
	}
}

func (cityMap *CitiesMap) destroyCity(cityName string) {
	city := cityMap.Cities[cityName]
	fmt.Printf("%s has been destroyed by alien %s and alien %s\n", cityName, city.Aliens[0].Name, city.Aliens[1].Name)

	killAliens(city)
	removeRoadsFromAdjacentCities(city)
	removeRoadsFromDestroyedCity(city)
	delete(cityMap.Cities, cityName)
}

func killAliens(city *City) {
	numberOfAliensAliveAndNotTrapped -= len(city.Aliens)
	city.Aliens = nil
}

func removeRoadsFromDestroyedCity(city *City) {
	// Remove roads from our destroyed city to each adj city
	for direction := range city.AdjCities {
		delete(city.AdjCities, direction)
	}
}

func removeRoadsFromAdjacentCities(city *City) {
	for _, adjCity := range city.AdjCities {
		for direction, adjCityCity := range adjCity.AdjCities {
			// We found our destroyed city
			if adjCityCity.Name == city.Name {
				delete(adjCity.AdjCities, direction)
				break
			}
		}
	}
}

func (cityMap *CitiesMap) moveAlien(cityName string, alienSteps int) {
	city := cityMap.Cities[cityName]
	alien := city.Aliens[0]

	// Return if the alien already moved in the iteration
	if alien.Steps == alienSteps {
		return
	}

	possibleDirections := getPossibleDirections(city)
	if isAlienTrapped(possibleDirections, alien) {
		fmt.Printf("The alien %s is trapped in %s!!!\n", alien.Name, cityName)
		return
	}

	// The alien moves to the new city
	directionToMove := getDirectionToMove(possibleDirections)
	destinationCity := moveAlienToNewCity(alien, directionToMove, city)
	fmt.Printf("The alien %s moved from %s to %s\n", alien.Name, cityName, destinationCity.Name)

	if len(destinationCity.Aliens) > 1 {
		cityMap.destroyCity(destinationCity.Name)
	}
}

func moveAlienToNewCity(alien aliens.Alien, directionToMove direction, city *City) *City {
	alien.Steps += 1
	destinationCity := city.AdjCities[directionToMove]
	destinationCity.Aliens = append(destinationCity.Aliens, alien)
	city.Aliens = nil

	return destinationCity
}

func isAlienTrapped(possibleDirections []direction, alien aliens.Alien) bool {
	if len(possibleDirections) == 0 {
		alien.Trapped = true
		numberOfAliensAliveAndNotTrapped -= 1
		return true
	}

	return false
}

func getDirectionToMove(possibleDirections []direction) direction {
	rand.Seed(time.Now().UnixNano())
	return possibleDirections[rand.Intn(len(possibleDirections))]
}

func getPossibleDirections(city *City) []direction {
	possibleDirections := make([]direction, 0, len(city.AdjCities))

	for direction, _ := range city.AdjCities {
		possibleDirections = append(possibleDirections, direction)
	}

	return possibleDirections
}
