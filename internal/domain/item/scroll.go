package item

import "github.com/tdutanton/Rogue_Game_go/internal/domain/common"

// Scroll represents a magical scroll that grants temporary character enhancements.
type Scroll struct {
	Name      string        // Name of the scroll (e.g., "Scroll of Whispering Shadows")
	Agility   int           // Agility bonus while active
	Strength  int           // Strength bonus while active
	MaxHealth int           // Max health increase while active
	Coords    common.Coords // Position on the map
	Duration  int           // Duration of the effect in turns
	IsActive  bool          // Indicates whether the scroll's effect is currently active
}

// Type returns the item type, used for identification and rendering.
func (s Scroll) Type() Type {
	return ScrollType
}

// Info returns the display name of the item for UI and logs.
func (s Scroll) Info() string {
	return s.Name
}

// GetCoords returns the current position of the item on the map.
func (s Scroll) GetCoords() common.Coords {
	return s.Coords
}

// SetCoords updates the item's position on the map.
func (s *Scroll) SetCoords(c common.Coords) {
	s.Coords.X = c.X
	s.Coords.Y = c.Y
}
