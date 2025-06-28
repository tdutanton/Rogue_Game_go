package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

// PursuingMoving implements EnemyMover for enemies that chase the player.
type PursuingMoving struct{}

// Move attempts to move the enemy along the shortest path toward the player.
// If no path exists, falls back to the enemy's default movement strategy.
func (p PursuingMoving) Move(e *Enemy, d dungeon.Dungeon) {
	path, _ := e.FindPathToPlayer(d)
	if e.EnemyType == Ghost {
		e.Visibility = common.RandomBool()
	}
	if len(path) > 0 {
		nextPos := path[0]
		if d.PlayerCoords() != nextPos {
			e.Coords.X = nextPos.X
			e.Coords.Y = nextPos.Y
		}
		return
	}
	// If no path found, use default movement
	if e.DefaultMover != nil {
		e.DefaultMover.Move(e, d)
	}
}
