package dungeon

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
)

// Coordinator interface to get abstract objects with coords
type Coordinator interface {
	GetCoords() common.Coords
	SetCoords(c common.Coords)
}

// IsCoordInRoom checks if coordinates are within a room's interior space.
func IsCoordInRoom(c common.Coords, r Room) bool {
	return c.X > r.X &&
		c.X <= r.X+r.FloorWidth() &&
		c.Y > r.Y &&
		c.Y <= r.Y+r.FloorHeight()
}

// IsCoordInWall checks if coordinates are on a room's perimeter walls.
func IsCoordInWall(c common.Coords, r Room) bool {
	left := c.X == r.X && c.Y >= r.Y && c.Y < r.Y+r.Height
	right := c.X == r.X+r.Width-1 && c.Y >= r.Y && c.Y < r.Y+r.Height
	top := c.Y == r.Y && c.X >= r.X && c.X < r.X+r.Width
	bottom := c.Y == r.Y+r.Height-1 && c.X >= r.X && c.X < r.X+r.Width
	return left || right || top || bottom
}

// IsCoordInCorridor checks if coordinates are within a corridor's path.
func IsCoordInCorridor(c common.Coords, corridor Corridor) bool {
	// Horizontal corridor
	if corridor.Begin.Y == corridor.End.Y {
		y := corridor.Begin.Y
		x1, x2 := corridor.Begin.X, corridor.End.X
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		return c.Y == y && c.X >= x1 && c.X <= x2
	}
	// Vertical corridor
	if corridor.Begin.X == corridor.End.X {
		x := corridor.Begin.X
		y1, y2 := corridor.Begin.Y, corridor.End.Y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		return c.X == x && c.Y >= y1 && c.Y <= y2
	}
	return false
}
