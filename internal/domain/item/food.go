package item

import "github.com/tdutanton/Rogue_Game_go/internal/domain/common"

// Food represents a consumable food that increase Character's health.
type Food struct {
	Name   string        // Name of food
	Value  int           // Number of health
	Coords common.Coords // Food's coords
}

// Type returns the item type, used for identification and rendering.
func (f Food) Type() Type {
	return FoodType
}

// Info returns the display name of the item for UI and logs.
func (f Food) Info() string {
	return f.Name
}

// GetCoords returns the current position of the item on the map.
func (f Food) GetCoords() common.Coords {
	return f.Coords
}

// SetCoords updates the item's position on the map.
func (f *Food) SetCoords(c common.Coords) {
	f.Coords.X = c.X
	f.Coords.Y = c.Y
}
