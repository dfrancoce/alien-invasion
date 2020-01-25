package validators

import (
	"testing"
)

func TestValidateAndGetArgs(t *testing.T) {
	inputFileString := "TestValidateAndGetArgsNilArgs"
	aliens := "100"

	file, numberOfAliens := ValidateAndGetArgs(func(string) ([]byte, error) {
		return []byte(inputFileString), nil
	}, inputFileString, aliens)

	if string(file) != inputFileString {
		t.Errorf("Expected the the file to be %s", inputFileString)
	}

	if numberOfAliens != 100 {
		t.Errorf("Expected the the number of aliens to be %s", aliens)
	}
}
