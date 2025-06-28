package unit

import (
	"math/rand"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
)

// ApplyEffect make some battle effects from Enemies
func (e *Enemy) ApplyEffect(target Fighter) {
	switch e.EnemyType {
	case Vampire:
		if player, ok := target.(*Character); ok {
			stolenHealth := common.RandomInRange(MinTakenMaxHealth, MaxTakenMaxHealth)
			player.MaxHealth -= stolenHealth
			if player.MaxHealth < 0 {
				player.MaxHealth = 0
			}
			if player.Health > player.MaxHealth {
				player.Health = player.MaxHealth
			}
		}
	case SnakeWizard:
		if rand.Float64() < 0.3 {
			if player, ok := target.(*Character); ok {
				player.SkipNextTurn = true
			}
		}
	default:
		// TODO тут допустим логирование для неизвестного врага
	}
}

// EnemyType defines the type of enemy.
type EnemyType int

// EnemyType represents different categories of enemies in the game.
// The behavior and movement patterns vary by enemy type.
const (
	Zombie      EnemyType = iota // Slow-moving, basic enemy
	Vampire                      // Fast-moving, may have special abilities
	Ghost                        // Can teleport and become invisible
	Ogr                          // Strong but slow enemy
	SnakeWizard                  // Moves diagonally with magical abilities
)

// EnemyNames provides a string values for all of enemy types.
var EnemyNames = map[EnemyType]string{
	Zombie:      "Zombie",
	Vampire:     "Vampire",
	Ghost:       "Ghost",
	Ogr:         "Ogr",
	SnakeWizard: "Snake-Wizard",
}

// SetCoords for enemy
func (e *Enemy) SetCoords(c common.Coords) {
	e.Coords.X = c.X
	e.Coords.Y = c.Y
}

// Enemy represents an enemy entity with stats, type, hostility, visibility, and movement behavior.
type Enemy struct {
	Unit
	EnemyType
	Animosity    int        // How aggressive the enemy is toward the player.
	Visibility   bool       // Used for ghost invisibility logic.
	Mover        EnemyMover // current move pattern
	DefaultMover EnemyMover // for switch from Pursuing
	IsPursuing   bool       // Move toward to Character if Enemy noticed him
	Treasure     int        // Number of treasurre which Character will receive after kill the monster
}

// TreasureFactors define proportions used in treasure generation calculations.
const (
	TreasureFactorAnimosty = 0.2
	TreasureFactorStrength = 0.2
	TreasureFactorAgility  = 0.2
	TreasureFactorHealth   = 0.2
)

// NewEnemy creates a new enemy of the specified type at given coordinates.
func NewEnemy(t EnemyType, x, y int) *Enemy {
	e := &Enemy{
		Unit: Unit{
			Coords: common.Coords{X: x, Y: y},
		},
		EnemyType:  t,
		Visibility: true,
		IsPursuing: Chilling,
	}

	switch t {
	case Ghost:
		e.DefaultMover = GhostMoving{}
	case SnakeWizard:
		e.DefaultMover = SnakeWizardMoving{}
	case Ogr:
		e.DefaultMover = OgrMoving{}
	default:
		e.DefaultMover = NonMoving{}
	}

	e.Mover = e.DefaultMover

	return e
}

// NewEnemyWithoutCoords create new enemy without placing on the field
func NewEnemyWithoutCoords(t EnemyType) *Enemy {
	e := &Enemy{
		Unit: Unit{
			Coords: common.Coords{},
		},
		EnemyType:  t,
		Visibility: true,
		IsPursuing: Chilling,
	}

	switch t {
	case Ghost:
		e.DefaultMover = GhostMoving{}
	case SnakeWizard:
		e.DefaultMover = SnakeWizardMoving{}
	case Ogr:
		e.DefaultMover = OgrMoving{}
	default:
		e.DefaultMover = NonMoving{}
	}

	e.Mover = e.DefaultMover

	return e
}

// CurrentStrength returns the enemy's current strength value.
func (e *Enemy) CurrentStrength() int {
	return e.Strength
}

// WeaponInHands returns nil for enemies, assuming they do not use weapons.
func (e *Enemy) WeaponInHands() *item.Weapon {
	return nil
}

// GoodEnemyTileSet defines tile common enemies can move on.
var GoodEnemyTileSet = map[common.TileType]struct{}{
	common.FloorTile: {},
}

// IsWalkableForEnemy returns true if the given tile type is traversable by enemies.
func IsWalkableForEnemy(t common.TileType) bool {
	_, ok := GoodEnemyTileSet[t]
	return ok
}

// isPossibleEnemyMove checks if an enemy can move to the specified coordinates.
func isPossibleEnemyMove(c common.Coords, r dungeon.Room, d dungeon.Dungeon) bool {
	if !dungeon.IsCoordInRoom(c, r) {
		return false
	}
	t, _ := d.Tile(c)
	return IsWalkableForEnemy(t)
}

// GetCoords - get current enemy's position
func (e *Enemy) GetCoords() common.Coords {
	return e.Unit.Coords
}

// HitCharacter attempts to hit a character, with a guaranteed hit for Ogr enemies.
// It calculates and applies damage if the hit is successful, and applies any enemy-specific effects.
// Returns a boolean indicating hit success and a placeholder value (currently always 0).
func (e *Enemy) HitCharacter(character *Character) bool {
	if IsHitSuccessful(&e.Unit, &character.Unit) || e.EnemyType == Ogr {
		damage := CalculateDamage(e)
		if e.EnemyType != Vampire {
			ApplyDamage(&character.Unit, damage)
		}
		e.ApplyEffect(character)
		character.Stats.HitsMissed++

		return true
	}
	return false
}
