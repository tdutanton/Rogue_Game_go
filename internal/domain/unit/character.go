package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/inventory"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
)

// Character represents the player character, including combat stats, weapon, and inventory.
type Character struct {
	Unit
	MaxHealth     int
	CurrentWeapon *item.Weapon
	Inventory     inventory.Inventory
	Stats         common.Stats
}

// NewCharacter - create new Player
func NewCharacter() *Character {
	return &Character{}
}

// SetCoords - set character's new position
func (ch *Character) SetCoords(c common.Coords) {
	ch.Unit.Coords.X = c.X
	ch.Unit.Coords.Y = c.Y
}

// CurrentStrength returns the player's current strength value.
func (ch *Character) CurrentStrength() int {
	return ch.Strength
}

// GetCoords - get current character's position
func (ch *Character) GetCoords() common.Coords {
	return ch.Unit.Coords
}

// WeaponInHands returns a pointer to the player's currently equipped weapon.
func (ch *Character) WeaponInHands() *item.Weapon {
	return ch.CurrentWeapon
}

// GoodPlayerTileSet defines tile common the player can walk on.
var GoodPlayerTileSet = map[common.TileType]struct{}{
	common.FloorTile:    {},
	common.ItemTile:     {},
	common.FinishTile:   {},
	common.DoorTile:     {},
	common.CorridorTile: {},
}

// IsWalkableForPlayer returns true if the given tile type is traversable by the player.
func IsWalkableForPlayer(t common.TileType) bool {
	_, ok := GoodPlayerTileSet[t]
	return ok
}

// IsPossiblePlayerMove checks if the player can move to the specified coordinates.
func IsPossiblePlayerMove(c common.Coords, d dungeon.Dungeon) bool {
	t, _ := d.Tile(c)
	return IsWalkableForPlayer(t)
}

// Step moves the unit one cell in the specified direction if possible.
func (ch *Character) Step(dir Direction, d dungeon.Dungeon) {
	switch dir {
	case Down:
		ch.stepDown(d)
	case Left:
		ch.stepLeft(d)
	case Up:
		ch.stepUp(d)
	case Right:
		ch.stepRight(d)
	}
}

// stepDown moves the unit downward (Y+1) if the target position is walkable.
func (ch *Character) stepDown(d dungeon.Dungeon) {
	newY := ch.Coords.Y + 1
	if IsPossiblePlayerMove(common.Coords{X: ch.Coords.X, Y: newY}, d) {
		ch.Coords.Y = newY
		ch.increaseStepsCount()
	}
}

// stepLeft moves the unit leftward (X-1) if the target position is walkable.
func (ch *Character) stepLeft(d dungeon.Dungeon) {
	newX := ch.Coords.X - 1
	if IsPossiblePlayerMove(common.Coords{X: newX, Y: ch.Coords.Y}, d) {
		ch.Coords.X = newX
		ch.increaseStepsCount()
	}
}

// stepUp moves the unit upward (Y-1) if the target position is walkable.
func (ch *Character) stepUp(d dungeon.Dungeon) {
	newY := ch.Coords.Y - 1
	if IsPossiblePlayerMove(common.Coords{X: ch.Coords.X, Y: newY}, d) {
		ch.Coords.Y = newY
		ch.increaseStepsCount()
	}
}

// stepRight moves the unit rightward (X+1) if the target position is walkable.
func (ch *Character) stepRight(d dungeon.Dungeon) {
	newX := ch.Coords.X + 1
	if IsPossiblePlayerMove(common.Coords{X: newX, Y: ch.Coords.Y}, d) {
		ch.Coords.X = newX
		ch.increaseStepsCount()
	}
}

// increaseStepsCount - just increment the number of character steps
func (ch *Character) increaseStepsCount() {
	ch.Stats.CellsPassed++
}
