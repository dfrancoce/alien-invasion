package _utils

import "testing"

func TestGetCitiesFromMapFile(t *testing.T) {
	testFile := []byte("Madrid north=Bilbao west=Lisboa south=Sevilla east=Valencia\n" +
		"Bilbao south=Madrid east=Zaragoza west=Santander")

	cities := GetCitiesFromMapFile(testFile)

	if cities["Madrid"] == nil {
		t.Error("Expected 'Madrid' to be among the cities")
	}

	if cities["Madrid"]["north"] != "Bilbao" {
		t.Error("Expected 'Bilbao' to be to the north of 'Madrid'")
	}

	if cities["Madrid"]["west"] != "Lisboa" {
		t.Error("Expected 'Lisboa' to be to the west of 'Madrid'")
	}

	if cities["Madrid"]["south"] != "Sevilla" {
		t.Error("Expected 'Sevilla' to be to the south of 'Madrid'")
	}

	if cities["Madrid"]["east"] != "Valencia" {
		t.Error("Expected 'Valencia' to be to the east of 'Madrid'")
	}

	if cities["Bilbao"] == nil {
		t.Error("Expected 'Bilbao' to be among the cities")
	}

	if cities["Bilbao"]["south"] != "Madrid" {
		t.Error("Expected 'Madrid' to be to the south of 'Bilbao'")
	}

	if cities["Bilbao"]["east"] != "Zaragoza" {
		t.Error("Expected 'Zaragoza' to be to the east of 'Bilbao'")
	}

	if cities["Bilbao"]["west"] != "Santander" {
		t.Error("Expected 'Santander' to be to the west of 'Bilbao'")
	}
}

func TestGetCitiesFromMapFilePassingEmptyLines(t *testing.T) {
	testFile := []byte("Madrid north=Bilbao west=Lisboa south=Sevilla east=Valencia\n" +
		"\n" +
		"Bilbao south=Madrid east=Zaragoza west=Santander")

	cities := GetCitiesFromMapFile(testFile)
	if cities[""] != nil {
		t.Error("Expected an empty string to not be among the cities'")
	}
}

func TestGetCitiesFromMapFileTheFileContainsReturn(t *testing.T) {
	testFile := []byte("Madrid north=Bilbao west=Lisboa south=Sevilla east=Valencia\r\n" +
		"Bilbao south=Madrid east=Zaragoza west=Santander")

	cities := GetCitiesFromMapFile(testFile)

	if len(cities) != 2 {
		t.Error("Expected the number of cities to be 2: Madrid and Bilbao")
	}
}

func TestGetCitiesFromMapFileNoSpaceSeparator(t *testing.T) {
	testFile := []byte("Madridnorth=Bilbao west=Lisboa south=Sevilla east=Valencia")

	cities := GetCitiesFromMapFile(testFile)
	if len(cities) != 0 {
		t.Error("Expected an error indicating the format of the line provided is not correct")
	}
}

func TestGetCitiesFromMapFileNumbers(t *testing.T) {
	testFile := []byte("M4dr1d n0rth=B1lb40 west=L1sb04 s0uth=Sev1ll4 e4st=V4lenc14")

	cities := GetCitiesFromMapFile(testFile)
	if len(cities) != 0 {
		t.Error("Expected an error indicating the format of the line provided is not correct")
	}
}

func TestGetCitiesFromMapFileIncludeNoLetters(t *testing.T) {
	testFile := []byte("Madrid_north=Bilbao-west=Lisboa-south=Sevilla-east=Valencia")

	cities := GetCitiesFromMapFile(testFile)
	if len(cities) != 0 {
		t.Error("Expected an error indicating the format of the line provided is not correct")
	}
}

func TestGetCitiesFromMapFileDirectionsContainsUppercase(t *testing.T) {
	testFile := []byte("Madrid nOrth=Bilbao West=Lisboa South=Sevilla easT=Valencia")

	cities := GetCitiesFromMapFile(testFile)
	if len(cities) != 0 {
		t.Error("Expected an error indicating the format of the line provided is not correct")
	}
}

func TestGetCitiesFromMapFileWrongDirections(t *testing.T) {
	testFile := []byte("Madrid Norte=Bilbao Oeste=Lisboa Sur=Sevilla Este=Valencia")

	cities := GetCitiesFromMapFile(testFile)
	if len(cities) != 0 {
		t.Error("Expected an error indicating the format of the line provided is not correct")
	}
}

func TestGetCitiesFromMapFileSameDirectionTwice(t *testing.T) {
	testFile := []byte("Madrid north=Bilbao north=Lisboa south=Sevilla east=Valencia")

	cities := GetCitiesFromMapFile(testFile)
	if len(cities) != 0 {
		t.Error("Expected an error indicating the format of the line provided is not correct")
	}
}
