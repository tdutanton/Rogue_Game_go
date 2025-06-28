package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

// SnakeWizardMoving implements EnemyMover for SnakeWizard enemies that move diagonally.
type SnakeWizardMoving struct{}

// Move implements diagonal movement for SnakeWizard enemies.
// Attempts random diagonal moves until a valid one is found.
func (s SnakeWizardMoving) Move(e *Enemy, d dungeon.Dungeon) {
	r := FindRoomByCoords(e.GetCoords(), d.Rooms[:])
	if r == nil {
		return
	}
	maxAttempts := 10
	for attempt := 0; attempt < maxAttempts; attempt++ {
		dir := common.RandomInRange(int(DownLeft), int(DownRight))
		var moved bool
		switch dir {
		case int(DownLeft):
			moved = wizardMoveDL(e, *r, d)
		case int(UpLeft):
			moved = wizardMoveUL(e, *r, d)
		case int(UpRight):
			moved = wizardMoveUR(e, *r, d)
		case int(DownRight):
			moved = wizardMoveDR(e, *r, d)
		}
		if moved {
			return
		}
	}
}

// wizardMoveDL attempts to move the enemy diagonally down-left.
// Returns true if movement was successful.
func wizardMoveDL(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X - 1
	newY := e.Coords.Y + 1
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}

// wizardMoveUL attempts to move the enemy diagonally up-left.
// Returns true if movement was successful.
func wizardMoveUL(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X - 1
	newY := e.Coords.Y - 1
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}

// wizardMoveUR attempts to move the enemy diagonally up-right.
// Returns true if movement was successful.
func wizardMoveUR(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X + 1
	newY := e.Coords.Y - 1
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}

// wizardMoveDR attempts to move the enemy diagonally down-right.
// Returns true if movement was successful.
func wizardMoveDR(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X + 1
	newY := e.Coords.Y + 1
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}
