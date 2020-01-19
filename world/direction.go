package world

type direction int

const (
	North direction = iota
	South
	East
	West
)

func convertToDirection(stringDirection string) direction {
	switch stringDirection {
	case "north":
		return North
	case "south":
		return South
	case "east":
		return East
	case "west":
		return West
	default:
		panic("The direction provided is not correct!!")
	}
}

func inverseDirection(direction direction) direction {
	switch direction {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("The direction provided is not correct!!")
	}
}