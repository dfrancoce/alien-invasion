package aliens

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

type Alien struct {
	Name string
	Trapped bool
	Steps int
}

func GenerateRandomAlien() Alien {
	var randomAlien Alien

	randomAlien.Name = randomdata.SillyName()
	randomAlien.Trapped = false
	randomAlien.Steps = 0

	fmt.Printf("A new alien named %s has been generated\n", randomAlien.Name)
	return randomAlien
}