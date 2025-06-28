package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

// OgrMoving implements EnemyMover for Ogr enemies that move for 2 tiles.
type OgrMoving struct{}

// Move implements 2-step movement for Ogr enemies.
// Attempts random 2-step moves until a valid one is found.
func (s OgrMoving) Move(e *Enemy, d dungeon.Dungeon) {
	r := FindRoomByCoords(e.GetCoords(), d.Rooms[:])
	if r == nil {
		return
	}
	maxAttempts := 10
	for attempt := 0; attempt < maxAttempts; attempt++ {
		dir := common.RandomInRange(int(Down), int(Right))
		var moved bool
		switch dir {
		case int(DownLeft):
			moved = ogrMoveD(e, *r, d)
		case int(UpLeft):
			moved = ogrMoveL(e, *r, d)
		case int(UpRight):
			moved = ogrMoveU(e, *r, d)
		case int(DownRight):
			moved = ogrMoveR(e, *r, d)
		}
		if moved {
			return
		}
	}
}

// ogrMoveDL attempts to move the enemy 2-step down.
// Returns true if movement was successful.
func ogrMoveD(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X
	newY := e.Coords.Y + 2
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}

// ogrMoveUL attempts to move the enemy 2-step left.
// Returns true if movement was successful.
func ogrMoveL(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X - 2
	newY := e.Coords.Y
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}

// ogrMoveUR attempts to move the enemy 2-step up.
// Returns true if movement was successful.
func ogrMoveU(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X
	newY := e.Coords.Y - 2
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}

// ogrMoveDR attempts to move the enemy 2-step right.
// Returns true if movement was successful.
func ogrMoveR(e *Enemy, r dungeon.Room, d dungeon.Dungeon) bool {
	newX := e.Coords.X + 2
	newY := e.Coords.Y
	if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, r, d) {
		e.Coords.X = newX
		e.Coords.Y = newY
		return true
	}
	return false
}
