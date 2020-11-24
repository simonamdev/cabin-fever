package physics

import "testing"

func TestPositionDoesNotChangeWithoutDirection(t *testing.T) {
	tables := []Position{
		{1, 2},
		{2, 1},
		{0, 0},
	}
	for _, position := range tables {
		newPosition := CalculateNextPosition(position, nil)
		if newPosition.X != position.X || newPosition.Y != position.Y {
			t.Errorf("New Position incorrect. X: Expected %d Got %d. Y: Expected %d Got %d", position.X, newPosition.X, position.Y, newPosition.Y)
		}
	}
}

type TestVector struct {
	Position
	Direction        Direction
	ExpectedPosition Position
}

func TestPositionDoesChangeWithDirection(t *testing.T) {
	tables := []TestVector{
		{Position: Position{X: 5, Y: 5}, Direction: Up, ExpectedPosition: Position{X: 5, Y: 4}},
		{Position: Position{X: 5, Y: 5}, Direction: Down, ExpectedPosition: Position{X: 5, Y: 6}},
		{Position: Position{X: 5, Y: 5}, Direction: Left, ExpectedPosition: Position{X: 4, Y: 5}},
		{Position: Position{X: 5, Y: 5}, Direction: Right, ExpectedPosition: Position{X: 6, Y: 5}},
	}
	for _, vector := range tables {
		newPosition := CalculateNextPosition(vector.Position, &vector.Direction)
		if newPosition.X != vector.ExpectedPosition.X || newPosition.Y != vector.ExpectedPosition.Y {
			t.Errorf("New Position incorrect. X: Expected %d Got %d. Y: Expected %d Got %d", vector.ExpectedPosition.X, newPosition.X, vector.ExpectedPosition.Y, newPosition.Y)
		}
	}
}
