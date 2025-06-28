// Package common contains shared constants and configuration used across the game,
// such as map dimensions, room generation parameters, and UI layout settings.
package common

// Dungeon room layout configuration
const (
	MaxRoomCount  = 9 // Maximum number of rooms in a dungeon level
	RoomsInRow    = 3 // Number of rooms per row
	RoomsInColumn = 3 // Number of rows
)

// WallThickness rendering thickness (in characters)
const WallThickness = 1

// Game map dimensions (in characters)
const (
	MapHeight = 30 // Total height of the dungeon map
	MapWidth  = 90 // Total width of the dungeon map
)

// UI panel dimensions
const (
	InfoHeight = 2  // Height of the info panel
	InfoWidth  = 90 // Width of the info panel

	StatisticHeight = 3  // Height of the statistics panel
	StatisticWidth  = 90 // Width of the statistics panel

	InventoryHeight = 30                                           // Height of the inventory panel
	InventoryWidth  = 90                                           // Width of the inventory panel
	MainHeight      = MapHeight + InfoHeight + StatisticHeight + 1 // MainHeight of main window
	MainWidth       = 90                                           // MainWidth of main window
)

// Room generation constraints
const (
	MinRoomDistance = 3                                                               // Minimum space between rooms
	MinRoomHeight   = 5                                                               // Minimum room height
	MinRoomWidth    = 6                                                               // Minimum room width
	MaxRoomHeight   = (MapHeight - MinRoomDistance*(RoomsInColumn-1)) / RoomsInColumn // Max possible room height
	MaxRoomWidth    = (MapWidth - MinRoomDistance*(RoomsInRow-1)) / RoomsInRow        // Max possible room width
)
