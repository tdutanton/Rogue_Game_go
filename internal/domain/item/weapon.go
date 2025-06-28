package item

import "github.com/tdutanton/Rogue_Game_go/internal/domain/common"

// Weapon represents an equippable item that boosts the player's attack power.
type Weapon struct {
	Name     string        // Name of the weapon (e.g., "Dagger of the Forgotten Shadow")
	Strength int           // Bonus damage or attack power provided by the weapon
	Coords   common.Coords // Position of the weapon on the map
}

// Type returns the item type, used for identification and rendering.
func (w Weapon) Type() Type {
	return WeaponType
}

// Info returns the display name of the item for UI and logs.
func (w Weapon) Info() string {
	return w.Name
}

// GetCoords returns the current position of the item on the map.
func (w Weapon) GetCoords() common.Coords {
	return w.Coords
}

// SetCoords updates the item's position on the map.
func (w *Weapon) SetCoords(c common.Coords) {
	w.Coords.X = c.X
	w.Coords.Y = c.Y
}

// SameWeapons checks if two weapons is equal
func SameWeapons(first, second Weapon) bool {
	return first.Name == second.Name &&
		first.Strength == second.Strength
}
