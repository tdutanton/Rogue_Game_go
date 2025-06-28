package logic

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// HandleAction processes a player's actions in the game: fighting, moving and collecting items.
// Monsters actions updating after every player's movement.
func HandleAction(dir unit.Direction, dg *dungeon.Dungeon) {
	// player := dg.Player.(*unit.Character)
	dg.ClearEventData()
	UpdateFights(dg)
	pAttacked := IsPlayerAttacked(dir, dg)
	RemoveDeadMonsters(dg)
	MonsterAttack(dg)

	if !pAttacked /* && !player.IsDead() */ {
		PlayerMove(dir, dg)
		CheckConsumables(dg)
	}

	MonstersMove(dg)
	UpdateItemsEffects(dg)
}

// IsEnemyNearby checks if the enemy is neighbour for the player.
// It considers horizontal, vertical, and diagonal (for snake wizard) neighbourhood.
// Returns true if an enemy is nearby, false otherwise.
func IsEnemyNearby(character dungeon.Coordinator, monster dungeon.Coordinator) bool {
	enemy := monster.GetCoords()
	player := character.GetCoords()

	horisontalNeighbour := (enemy.X == player.X+1 || enemy.X == player.X-1) && enemy.Y == player.Y
	verticalNeighbour := (enemy.Y == player.Y+1 || enemy.Y == player.Y-1) && enemy.X == player.X

	if horisontalNeighbour || verticalNeighbour {
		return true
	}

	if e, ok := monster.(*unit.Enemy); ok && e.EnemyType == unit.SnakeWizard {

		leftDiagonal := (enemy.X == player.X-1 && (enemy.Y == player.Y-1 || enemy.Y == player.Y+1))
		rightDiagonal := (enemy.X == player.X+1 && (enemy.Y == player.Y-1 || enemy.Y == player.Y+1))

		if leftDiagonal || rightDiagonal {
			return true
		}
	}
	return false
}

// GetCoordsAfterMoving calculates the new coordinates for a character after moving in a specified direction.
// It returns the updated coordinates based on the input direction, or the original coordinates if an invalid direction is provided.
func GetCoordsAfterMoving(player common.Coords, dir unit.Direction) common.Coords {
	switch dir {
	case unit.Up:
		return common.Coords{X: player.X, Y: player.Y - 1}
	case unit.Down:
		return common.Coords{X: player.X, Y: player.Y + 1}
	case unit.Left:
		return common.Coords{X: player.X - 1, Y: player.Y}
	case unit.Right:
		return common.Coords{X: player.X + 1, Y: player.Y}
	default:
		return player
	}
}

// MonstersMove iterates through all enemies in the dungeon and calls their individual Move method,
// allowing each enemy to update its position within the dungeon's current state.
func MonstersMove(dg *dungeon.Dungeon) {
	for i := range dg.Enemies {
		enemy := dg.Enemies[i].(*unit.Enemy)
		if !enemy.InBattle {
			enemy.Move(*dg)
		}
	}
}

// UpdateItemsEffects manages the duration and effects of active elixirs and scrolls in the player's inventory.
func UpdateItemsEffects(dg *dungeon.Dungeon) {
	player := dg.Player.(*unit.Character)

	for i := len(player.Inventory.Elixirs) - 1; i >= 0; i-- {
		elixir := &player.Inventory.Elixirs[i]
		if elixir.IsActive {
			elixir.Duration--
		}
		if elixir.Duration <= 0 {
			elixir.IsActive = false
			DeleteEffectValues(dg, elixir.Agility, elixir.Strength, elixir.MaxHealth)
			player.Inventory.Delete(elixir)
		}
	}

	for i := len(player.Inventory.Scrolls) - 1; i >= 0; i-- {
		scroll := &player.Inventory.Scrolls[i]
		if scroll.IsActive {
			scroll.Duration--
		}
		if scroll.Duration <= 0 {
			scroll.IsActive = false
			DeleteEffectValues(dg, scroll.Agility, scroll.Strength, scroll.MaxHealth)
			player.Inventory.Delete(scroll)
		}
	}
}

// DeleteEffectValues removes the effects of an item from the player's attributes
// and remove item from the player's inventory.
func DeleteEffectValues(dg *dungeon.Dungeon, agility, strength, maxHealth int) {
	player := dg.Player.(*unit.Character)

	player.Agility -= agility
	if player.Agility < 0 {
		player.Agility = 0
	}

	player.Strength -= strength
	if player.Strength < 0 {
		player.Strength = 0
	}

	player.MaxHealth -= maxHealth
	if player.MaxHealth <= 0 {
		player.MaxHealth = 5
	}
	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}
}
