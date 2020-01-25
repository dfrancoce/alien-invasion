package validators

import (
	"log"
	"strconv"
)

func ValidateAndGetArgs(reader func(string) ([]byte, error), inputFile string, aliens string) ([]byte, int) {
	file := validateAndGetInputFile(reader, inputFile)
	numberOfAliens := validateAndGetNumberOfAliens(aliens)

	return file, numberOfAliens
}

func validateAndGetNumberOfAliens(aliens string) int {
	numberOfAliens, err := strconv.Atoi(aliens)

	if err != nil {
		log.Fatalf("Error reading number of aliens parameter: %s", err.Error())
	}

	if numberOfAliens <= 0 {
		log.Fatal("The number of aliens must be greater than zero")
	}

	return numberOfAliens
}

func validateAndGetInputFile(reader func(string) ([]byte, error), inputFile string) []byte {
	file, err := reader(inputFile)

	if err != nil {
		log.Fatal("Error reading the file containing the world")
	}

	return file
}
