package aliens

import "testing"

func TestGenerateRandomAlien(t *testing.T) {
	alien := GenerateRandomAlien()
	expectedRandomSteps := 0

	if alien.Name == "" {
		t.Errorf("Expected the Name field %s to be not empty", alien.Name)
	}

	if alien.Trapped != false {
		t.Errorf("Expected the trapped field to be %t but instead got %t", false, alien.Trapped)
	}

	if alien.Steps != expectedRandomSteps {
		t.Errorf("Expected the steps field to be %d but instead got %d", expectedRandomSteps, alien.Steps)
	}
}