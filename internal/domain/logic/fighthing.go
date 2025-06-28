package logic

import (
	"strconv"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// UpdateFights manages the battle state of enemies near the player.
func UpdateFights(dg *dungeon.Dungeon) {
	for i := range dg.Enemies {
		enemy := dg.Enemies[i].(*unit.Enemy)
		if IsEnemyNearby(dg.Player, enemy) {
			if !enemy.InBattle {
				SwitchToBattleMode(enemy)
			}
		} else {
			if enemy.InBattle {
				enemy.InBattle = false
			}
		}
	}
}

// SwitchToBattleMode sets monsters properties, which are used in battles.
func SwitchToBattleMode(enemy *unit.Enemy) {
	enemy.InBattle = true
	enemy.Visibility = true
	if enemy.EnemyType == unit.Ogr {
		enemy.SkipNextTurn = false
	}
	if enemy.EnemyType == unit.Vampire {
		enemy.FirstHitMiss = true
	}
}

// IsPlayerAttacked checks if the player is attacked when moving in a specific direction.
// Returns true if the player is attacked (or try to attack), false otherwise.
func IsPlayerAttacked(dir unit.Direction, dg *dungeon.Dungeon) bool {
	newPlayerCoords := GetCoordsAfterMoving(dg.Player.GetCoords(), dir)
	player := dg.Player.(*unit.Character)
	isAttacked := false

	for i := range dg.Enemies {
		enemy := dg.Enemies[i].(*unit.Enemy)
		if newPlayerCoords == enemy.Coords {
			isAttacked = true
			if !player.SkipNextTurn {
				hit, gold := player.HitEnemy(enemy)
				UpdateAttackData(dg, player, enemy, hit, gold)
				break
			} else {
				player.SkipNextTurn = false
			}
		}
	}

	return isAttacked
}

// MonsterAttack handles enemy attacks on the player when enemies are nearby and in battle mode.
func MonsterAttack(dg *dungeon.Dungeon) {
	player := dg.Player.(*unit.Character)

	for i := range dg.Enemies {
		enemy := dg.Enemies[i].(*unit.Enemy)
		if IsEnemyNearby(player, enemy) && enemy.InBattle {
			skip := enemy.SkipNextTurn

			if enemy.EnemyType == unit.Ogr {
				enemy.SkipNextTurn = !enemy.SkipNextTurn
			}

			if !skip {
				hit := enemy.HitCharacter(player)
				UpdateAttackData(dg, enemy, player, hit, 0)
				break
			}
		}
	}
}

// UpdateAttackData add the result of an attack between a fighter and a defender in EventData array
// for output in game screen.
func UpdateAttackData(dg *dungeon.Dungeon, attacker, defender unit.Fighter, goodHit bool, gold int) {
	if enemy, ok := defender.(*unit.Enemy); ok {
		if !goodHit {
			dg.AddEventData("You missed " + unit.EnemyNames[enemy.EnemyType] + "...")
		} else if goodHit && gold == 0 {
			dg.AddEventData("You attacked " + unit.EnemyNames[enemy.EnemyType] + "!")
		} else if goodHit && gold > 0 {
			dg.AddEventData("You defeated " + unit.EnemyNames[enemy.EnemyType] + "! You got " + strconv.Itoa(gold) + " gold.")
		}
	}

	if enemy, ok := attacker.(*unit.Enemy); ok {
		if !goodHit {
			dg.AddEventData(unit.EnemyNames[enemy.EnemyType] + " missed.")
		} else if goodHit {
			dg.AddEventData(unit.EnemyNames[enemy.EnemyType] + " attacked!")
		}
	}
}

// RemoveDeadMonsters removes dead enemies from the dungeon's enemy list.
func RemoveDeadMonsters(dg *dungeon.Dungeon) {
	for i := len(dg.Enemies) - 1; i >= 0; i-- {
		enemy := dg.Enemies[i].(*unit.Enemy)
		if enemy.IsDead() {
			dg.Enemies = append(dg.Enemies[:i], dg.Enemies[i+1:]...)
		}
	}
}
