package unit

import (
	"math/rand"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

const (
	// BaseAgility is the baseline chance to hit before agility difference adjustment.
	BaseAgility = 0.5

	// MinAgility is the minimum chance to hit.
	MinAgility = 0.1

	// MaxAgility is the maximum chance to hit.
	MaxAgility = 0.9

	// AgilityStep is the multiplier applied to the agility difference to adjust chance to hit.
	AgilityStep = 0.05

	// MinTakenMaxHealth Vampire can take some max health - min value
	MinTakenMaxHealth = 1

	// MaxTakenMaxHealth Vampire can take some max health - max value
	MaxTakenMaxHealth = 3
)

// Unit represents a generic entity with health, agility, strength, and coordinates.
type Unit struct {
	Health       int
	Agility      int
	Strength     int
	Coords       common.Coords
	InBattle     bool
	FirstHitMiss bool // if Unit is a Vampire - Character miss first hit
	SkipNextTurn bool // for Character - wizard can throw him in a sleep, for Enemy - ogr skip every second attack
}

// GetCoords - get current unit's position
func (u *Unit) GetCoords() common.Coords {
	return u.Coords
}

// IsDead - check if unit is dead
func (u *Unit) IsDead() bool {
	return u.Health <= 0
}

// SetCoords - set new position to unit
func (u *Unit) SetCoords(c common.Coords) {
	u.Coords.X = c.X
	u.Coords.Y = c.Y
}

// Fighter defines the interface for entities that can perform attacks.
// It requires methods to get current strength and equipped weapon.
type Fighter interface {
	// CurrentStrength returns the fighter's effective strength value.
	CurrentStrength() int
	// // WeaponInHands returns a pointer to the weapon currently held, or nil if unarmed.
	// WeaponInHands() *item.Weapon
}

// ChanceToHit calculates the probability that an attacker hits the defender
// based on their agility values. The result is clamped between MinAgility and MaxAgility.
func ChanceToHit(attacker, defender Unit) float64 {
	base := BaseAgility + float64(attacker.Agility-defender.Agility)*AgilityStep
	if base < MinAgility {
		base = MinAgility
	} else if base > MaxAgility {
		base = MaxAgility
	}
	return base
}

// IsHitSuccessful determines whether an attacker's hit attempt succeeds
// taking into account the special first-hit miss condition and a random roll.
func IsHitSuccessful(attacker, defender *Unit) bool {
	if defender.FirstHitMiss {
		defender.FirstHitMiss = false
		return false
	}
	return rand.Float64() < ChanceToHit(*attacker, *defender)
}

// CalculateDamage computes the damage dealt by the attacker.
// It adds the attacker's current strength and weapon power if applicable.
func CalculateDamage(attacker Fighter) int {
	damage := attacker.CurrentStrength()
	if p, ok := attacker.(*Character); ok {
		if w := p.WeaponInHands(); w != nil {
			damage += w.Strength
		}
	}
	return damage
}

// ApplyDamage applies the given damage amount to the target unit,
// reducing its health but not allowing it to go below zero.
func ApplyDamage(target *Unit, dmg int) {
	target.Health -= dmg
	if target.Health <= 0 {
		target.Health = 0
	}
}

// FindRoomByCoords locates the room containing the specified coordinates.
// Returns nil if no room contains the coordinates.
func FindRoomByCoords(c common.Coords, rooms []dungeon.Room) *dungeon.Room {
	for i := range rooms {
		if rooms[i].Contains(c) {
			return &rooms[i]
		}
	}
	return nil
}
