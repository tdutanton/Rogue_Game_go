package item

import "github.com/tdutanton/Rogue_Game_go/internal/domain/common"

// Type represents the category of an item in the game.
type Type int

// Item types available in the game.
const (
	EmptyType  Type = iota // Placeholder for no item
	FoodType               // Consumable that restores health
	ElixirType             // Potion that grants temporary stat boosts
	ScrollType             // Magic scroll with a one-time effect
	WeaponType             // Equippable weapon that boosts attack power
)

// Item is a common interface implemented by all collectible objects in the dungeon.
// It provides methods for position management, type identification, and display info.
type Item interface {
	GetCoords() common.Coords  // Get current map coordinates
	SetCoords(c common.Coords) // Set new map coordinates
	Type() Type                // Get item category
	Info() string              // Get human-readable description or name
}

// ItemsNames provides a string values for all of items types.
var ItemsNames = map[Type]string{
	FoodType:   "food",
	ElixirType: "elixir",
	ScrollType: "scroll",
	WeaponType: "weapon",
}
