package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

// GhostMoving implements EnemyMover for ghost-type enemies that teleport randomly.
type GhostMoving struct{}

// Move implements random teleportation movement for Ghost enemies.
// Makes the enemy invisible half the time and teleports to a random valid position.
func (s GhostMoving) Move(e *Enemy, d dungeon.Dungeon) {
	r := FindRoomByCoords(e.GetCoords(), d.Rooms[:])
	if r == nil {
		return
	}

	e.Visibility = common.RandomBool()
	maxAttempts := 10
	for attempt := 0; attempt < maxAttempts; attempt++ {
		newX := common.RandomInRange(r.X, r.X+r.Width-1)
		newY := common.RandomInRange(r.Y, r.Y+r.Height-1)
		if isPossibleEnemyMove(common.Coords{X: newX, Y: newY}, *r, d) {
			e.Coords.X = newX
			e.Coords.Y = newY
			return
		}
	}
}
