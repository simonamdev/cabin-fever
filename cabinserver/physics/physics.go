package physics

//Position describes a point in positive co-ordinate space
type Position struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

type Direction string

const (
	Up    Direction = "UP"
	Down  Direction = "DOWN"
	Left  Direction = "LEFT"
	Right Direction = "RIGHT"
)

func CalculateNextPosition(position Position, direction *Direction) Position {
	// No direction, so stay in the same place
	if direction == nil {
		return position
	}
	newPosition := Position{X: position.X, Y: position.Y}
	// According to the direction, change the position
	if *direction == Up {
		newPosition.Y--
	} else if *direction == Down {
		newPosition.Y++
	} else if *direction == Left {
		newPosition.X--
	} else if *direction == Right {
		newPosition.X++
	}
	return newPosition
}
