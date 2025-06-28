package common

// Coords represents a 2D coordinate position with X and Y values as unsigned integers.
type Coords struct {
	X, Y int
}

// Size represents the dimensions of a rectangular area with a width and height.
type Size struct {
	Width, Height int
}

// TileType represents the different common of tiles in the dungeon.
type TileType int

// TileType represents different types of dungeon tiles.
// Used for game logic and rendering differentiation.
const (
	UnknownTile  TileType = iota // Undiscovered or invalid tile
	FloorTile                    // Walkable floor space
	WallTile                     // Impassable wall
	ItemTile                     // Contains an item (weapon, food, etc)
	EnemyTile                    // Contains an enemy unit
	FinishTile                   // Level exit/win condition
	DoorTile                     // Door/passage between areas
	PlayerTile                   // Current player position
	CorridorTile                 // Connecting passage between rooms
)

// Stats contains player statistics throughout the game
// including treasures collected, enemies defeated, and other
type Stats struct {
	TreasuresReceived int
	LevelAchieved     int
	EnemiesDefeated   int
	FoodEaten         int
	ElixirsDrunk      int
	ScrollsRead       int
	HitsMade          int
	HitsMissed        int
	CellsPassed       int
}
