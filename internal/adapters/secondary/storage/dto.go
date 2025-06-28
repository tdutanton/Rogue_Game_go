package storage

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
)

// CoordsData represents 2D coordinates in the game world.
type CoordsData struct {
	X int `json:"x"` // X coordinate (horizontal position)
	Y int `json:"y"` // Y coordinate (vertical position)
}

// SizeData represents dimensions of game objects or spaces.
type SizeData struct {
	Width  int `json:"width"`  // Width dimension
	Height int `json:"height"` // Height dimension
}

// ItemData contains information about game items.
// It includes positional data, attributes, and state information.
type ItemData struct {
	Type       int             `json:"type"` // Item type identifier
	CoordsData `json:"coords"` // Position in the game world
	Name       string          `json:"name"`                 // Display name of the item
	Agility    int             `json:"agility,omitempty"`    // Agility modifier (if applicable)
	Strength   int             `json:"strength,omitempty"`   // Strength modifier (if applicable)
	MaxHealth  int             `json:"max_health,omitempty"` // Max health modifier (if applicable)
	Value      int             `json:"value_food,omitempty"` // Nutritional value (for food items)
	Duration   int             `json:"duration,omitempty"`   // Effect duration in turns
	IsActive   bool            `json:"is_active,omitempty"`  // Active state flag
}

// StatsData tracks various player statistics and achievements.
type StatsData struct {
	TreasuresReceived int `json:"treasures"`        // Total treasures collected
	LevelAchieved     int `json:"level_achieved"`   // Highest level reached
	EnemiesDefeated   int `json:"enemies_defeated"` // Total enemies defeated
	FoodEaten         int `json:"food_eaten"`       // Total food consumed
	ElixirsDrunk      int `json:"elixirs_drunk"`    // Total elixirs consumed
	ScrollsRead       int `json:"scrolls_read"`     // Total scrolls used
	HitsMade          int `json:"hits_ok"`          // Successful attacks
	HitsMissed        int `json:"hits_missed"`      // Missed attacks
	CellsPassed       int `json:"cells_passed"`     // Total movement steps taken
}

// InventoryData represents the player's inventory, containing various item categories.
type InventoryData struct {
	Weapons  []ItemData `json:"weapons,omitempty"` // Collected weapons
	Elixirs  []ItemData `json:"elixirs,omitempty"` // Collected elixirs
	Scrolls  []ItemData `json:"scrolls,omitempty"` // Collected scrolls
	Foods    []ItemData `json:"foods,omitempty"`   // Collected food items
	Treasure int        `json:"treasure"`          // Current treasure amount
}

// UnitData contains common attributes for both characters and enemies.
type UnitData struct {
	Health     int             `json:"health"`   // Current health points
	Agility    int             `json:"agility"`  // Agility attribute
	Strength   int             `json:"strength"` // Strength attribute
	CoordsData `json:"coords"` // Current position
	InBattle   bool            `json:"in_battle"` // Battle engagement status
}

// CharacterData represents the player character with all associated data.
type CharacterData struct {
	Unit          UnitData      `json:"base_data"`        // Basic unit attributes
	MaxHealth     int           `json:"max_health"`       // Maximum health capacity
	CurrentWeapon ItemData      `json:"weapon,omitempty"` // Currently equipped weapon
	Inventory     InventoryData `json:"backpack"`         // Player inventory
	Stats         StatsData     `json:"stats"`            // Game statistics
}

// RoomData describes a dungeon room with its properties and features.
type RoomData struct {
	SizeData   `json:"size"`                       // Room dimensions
	CoordsData `json:"room_up_left_corner_coords"` // Top-left corner coordinates
	Doors      []CoordsData                        `json:"doors"`   // Door locations
	Type       int                                 `json:"type"`    // Room type identifier
	Visited    bool                                `json:"visited"` // Room visited identified
	Visible    int                                 `json:"visible"` // Room visible identifier
}

// CorridorData represents a connecting path between two points.
type CorridorData struct {
	Begin   CoordsData `json:"begin"`   // Starting coordinates
	End     CoordsData `json:"end"`     // Ending coordinates
	Visited bool       `json:"visited"` // Corridor visited identifier
}

// PassageData contains a series of connected corridors forming a path.
type PassageData struct {
	Path []CorridorData `json:"corridors"` // Sequence of corridor segments
}

// EnemyData represents an enemy entity with combat attributes and behavior flags.
type EnemyData struct {
	Unit       UnitData `json:"base_data"`   // Basic unit attributes
	EnemyType  int      `json:"enemy_type"`  // Enemy classification
	Animosity  int      `json:"animosity"`   // Aggressiveness level
	Visibility bool     `json:"visibility"`  // Visibility status
	IsPursuing bool     `json:"is_pursuing"` // Pursuit behavior flag
	Treasure   int      `json:"treasure"`    // Treasure carried by enemy
}

// DungeonData contains all information about a dungeon level.
// It includes the layout, entities, and player state.
type DungeonData struct {
	LevelNumber int                           `json:"level"`     // Current dungeon level
	Rooms       [common.MaxRoomCount]RoomData `json:"rooms"`     // Array of rooms in the dungeon
	Passages    []PassageData                 `json:"passages"`  // Connecting passages between rooms
	Exit        CoordsData                    `json:"exit"`      // Dungeon exit location
	Player      CharacterData                 `json:"character"` // Player character data
	Items       []ItemData                    `json:"items"`     // Items present in the dungeon
	Enemies     []EnemyData                   `json:"enemies"`   // Enemies present in the dungeon
}
