package item

import "github.com/tdutanton/Rogue_Game_go/internal/domain/common"

// Elixir represents a consumable potion that temporarily boosts character attributes.
type Elixir struct {
	Name      string        // Name of the elixir (e.g., "Elixir of the Hollow Vein")
	Agility   int           // Agility boost value
	Strength  int           // Strength boost value
	MaxHealth int           // Max health boost value
	Coords    common.Coords // Position on the map
	Duration  int           // Number of turns the effect lasts
	IsActive  bool          // Whether the effect is currently active
}

// Type returns the item type, used for identification and rendering.
func (e Elixir) Type() Type {
	return ElixirType
}

// Info returns the display name of the item for UI and logs.
func (e Elixir) Info() string {
	return e.Name
}

// GetCoords returns the current position of the item on the map.
func (e Elixir) GetCoords() common.Coords {
	return e.Coords
}

// SetCoords updates the item's position on the map.
func (e *Elixir) SetCoords(c common.Coords) {
	e.Coords.X = c.X
	e.Coords.Y = c.Y
}
