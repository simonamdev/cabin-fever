package physics

import "testing"

func TestPositionDoesNotChangeWithoutDirection(t *testing.T) {
	x := uint(1)
	y := uint(2)
	position := Position{X: x, Y: y}
	newPosition := CalculateNextPosition(position, nil)
	if newPosition.X != position.X || newPosition.Y != position.Y {
		t.Errorf("New Position incorrect. X: Expected %d Got %d. Y: Expected %d Got %d", x, newPosition.X, y, newPosition.Y)
	}
}
