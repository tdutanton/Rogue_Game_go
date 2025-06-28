package dungeon

import (
	"math/rand"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
)

// RoomType represents the different common of rooms in a dungeon
type RoomType int

const (
	// RoomPlain is common room type
	RoomPlain RoomType = iota
	// RoomStart is room where player starts level
	RoomStart
	// RoomEnd is room with exit to the next level
	RoomEnd
)

// Visibility represent status of area to check if it needed to be filled by a fog of war
type Visibility int

const (
	// FogCovered All room covered by fog of war
	FogCovered Visibility = iota
	// FogClean All room is clean (character in room)
	FogClean
	// VerticalFog Character stay at the northern or southern door - draw vertical sector of fog
	VerticalFog
	// HorizontalFog Character stay at the western or eastern door - draw horizontal sector of fog
	HorizontalFog
)

// Room represents a rectangular area with a specific size and coordinates.
type Room struct {
	common.Size              // Room size including walls
	common.Coords            // Coordinates of the upper left corner of the room
	Doors         []Door     // Slice of room doors coordinates
	Type          RoomType   // Type of room in dungeon
	Visited       bool       // Indicates whether the room has been visited by a player
	Visible       Visibility // Indicates is room visible by a player
}

// generateRoom creates a randomly sized and positioned room within a grid layout.
// Returns a ready Room.
func generateRoom(number int) Room {
	var room Room

	col := number % common.RoomsInRow
	row := number / common.RoomsInColumn

	room.Height = rand.Intn(common.MaxRoomHeight-common.MinRoomHeight+1) + common.MinRoomHeight
	room.Width = rand.Intn(common.MaxRoomWidth-common.MinRoomWidth+1) + common.MinRoomWidth

	cellX := col * (common.MaxRoomWidth + common.MinRoomDistance)
	cellY := row * (common.MaxRoomHeight + common.MinRoomDistance)

	room.X = cellX + rand.Intn(common.MaxRoomWidth-room.Width+1)
	room.Y = cellY + rand.Intn(common.MaxRoomHeight-room.Height+1)
	room.Visited = false

	return room
}

// FloorWidth returns the interior width of the room (excluding walls).
func (r *Room) FloorWidth() int {
	return r.Width - 2
}

// FloorHeight returns the interior height of the room (excluding walls).
func (r *Room) FloorHeight() int {
	return r.Height - 2
}

// Contains checks if the given coordinates are within this room's boundaries.
func (r *Room) Contains(c common.Coords) bool {
	return c.X > r.X && c.X < r.X+r.Width-1 &&
		c.Y > r.Y && c.Y < r.Y+r.Height-1
}

// ContainsIncludeWalls checks if the given coordinates are within this room's boundaries include walls.
func (r *Room) ContainsIncludeWalls(c common.Coords) bool {
	return c.X >= r.X && c.X <= r.X+r.Width-1 &&
		c.Y >= r.Y && c.Y <= r.Y+r.Height-1
}

// generateExitPoint creates point for exit from level in specified room
func generateExitPoint(room *Room) common.Coords {
	minX, minY := room.X+1, room.Y+1
	maxX, maxY := room.X+room.Width-2, room.Y+room.Height-2

	randXCoord := rand.Intn(maxX-minX-1) + minX + 1
	randYCoord := rand.Intn(maxY-minY-1) + minY + 1

	return common.Coords{X: randXCoord, Y: randYCoord}
}
