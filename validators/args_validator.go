package validators

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ValidateAndGetArgs() ([]byte, int) {
	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error reading the file containing the world")
		os.Exit(1)
	}

	numberOfAliens, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Error reading number of aliens parameter")
		os.Exit(1)
	}

	return file, numberOfAliens
}