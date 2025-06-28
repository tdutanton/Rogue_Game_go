package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

const (
	// Pursuing Enemy pursuit states
	// Pursuing Enemy is actively chasing the player
	Pursuing = true
	// Chilling Enemy is following its default movement pattern
	Chilling = false

	// Visible Ghost visibility states
	// Visible Ghost is currently visible to the player
	Visible = true
	// UnVisible Ghost is currently invisible to the player
	UnVisible = false
)

// EnemyMover defines a strategy interface for enemy movement behavior.
// Implementations determine how enemies move when not directly pursuing the player.
type EnemyMover interface {
	Move(e *Enemy, d dungeon.Dungeon)
}

// Move executes the enemy's current movement strategy after updating its mode.
// It first calls SetMovingMode to determine the appropriate movement strategy,
// then invokes the mover's Move method if a strategy is set.
func (e *Enemy) Move(d dungeon.Dungeon) {
	if e.Mover != nil {
		e.SetMovingMode(d)
		e.Mover.Move(e, d)
	}
}

// InRoomWithPlayer checks if both the enemy and player are in the same room.
func (e *Enemy) InRoomWithPlayer(d dungeon.Dungeon) bool {
	p := d.PlayerCoords()
	r := FindRoomByCoords(p, d.Rooms[:])
	if r != nil && dungeon.IsCoordInRoom(e.GetCoords(), *r) && dungeon.IsCoordInRoom(p, *r) {
		return true
	}
	return false
}

// SetMovingMode determines and sets the enemy's movement strategy based on:
// - Proximity to player
// - Current location (room/corridor)
// - Path availability
// - Animosity (pursuit range)
func (e *Enemy) SetMovingMode(d dungeon.Dungeon) {
	c, _ := d.TileUnderEnemy(e.GetCoords())
	_, path := e.FindPathToPlayer(d)
	shouldPursue := path != -1 && path <= e.Animosity
	if e.IsPursuing && shouldPursue {
		e.Mover = PursuingMoving{}
		return
	}
	if !e.IsPursuing && e.InRoomWithPlayer(d) && shouldPursue {
		e.IsPursuing = Pursuing
		e.Mover = PursuingMoving{}
		return
	}
	e.IsPursuing = Chilling
	if c == common.CorridorTile || c == common.DoorTile {
		e.Mover = NonMoving{}
	} else {
		e.Mover = e.DefaultMover
	}
}

// FindPathToPlayer calculates the shortest path to the player using BFS.
// Returns the path (excluding current position) and its length.
// Returns empty slice and -1 if no path exists.
func (e *Enemy) FindPathToPlayer(d dungeon.Dungeon) ([]common.Coords, int) {
	start := e.GetCoords()
	target := d.PlayerCoords()
	if start == target {
		return []common.Coords{}, -1
	}
	queue := [][]common.Coords{{start}}
	visited := map[common.Coords]bool{start: true}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]
		for _, dir := range []Direction{Up, Down, Left, Right} {
			next := current
			switch dir {
			case Up:
				next.Y--
			case Down:
				next.Y++
			case Left:
				next.X--
			case Right:
				next.X++
			}
			if next.X < 0 || next.Y < 0 {
				continue
			}
			if visited[next] {
				continue
			}
			tile, err := d.Tile(next)
			if err != nil {
				continue
			}
			if tile == common.FloorTile || tile == common.DoorTile ||
				tile == common.CorridorTile || next == target {
				if next == target {
					newPath := append([]common.Coords(nil), path...)
					newPath = append(newPath, next)
					return newPath[1:], len(newPath)
				}
				visited[next] = true
				newPath := append([]common.Coords(nil), path...)
				newPath = append(newPath, next)
				queue = append(queue, newPath)
			}
		}
	}
	return []common.Coords{}, -1
}
