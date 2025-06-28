package unit

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
)

// IsDead checks if the character's health has dropped to zero or below.
func (ch *Character) IsDead() bool {
	return ch.Health <= 0
}

// MoveRight moves the character one step to the right on the map.
func (ch *Character) MoveRight() {
	ch.Coords.X++
	ch.increaseStepsCount()
}

// MoveLeft moves the character one step to the left on the map.
func (ch *Character) MoveLeft() {
	ch.Coords.X--
	ch.increaseStepsCount()
}

// MoveDown moves the character one step downward on the map.
func (ch *Character) MoveDown() {
	ch.Coords.Y++
	ch.increaseStepsCount()
}

// MoveUp moves the character one step upward on the map.
func (ch *Character) MoveUp() {
	ch.Coords.Y--
	ch.increaseStepsCount()
}

// TakeItem adds an item to the character's inventory.
func (ch *Character) TakeItem(item item.Item) {
	ch.Inventory.Add(item)
}

// DropItem removes an item from the character's inventory and places it on the map.
func (ch *Character) DropItem(item item.Item) {
}

// EatFood restores the character's health by the food's value up to max health.
func (ch *Character) EatFood(food *item.Food) {
	ch.Health += food.Value
	if ch.Health > ch.MaxHealth {
		ch.Health = ch.MaxHealth
	}
	ch.Stats.FoodEaten++
	ch.Inventory.Delete(food)
}

// DrinkElixir applies the elixir's temporary stat boosts to the character.
func (ch *Character) DrinkElixir(elixir item.Elixir) {
	ch.MaxHealth += elixir.MaxHealth
	ch.Agility += elixir.Agility
	ch.Strength += elixir.Strength
	ch.Stats.ElixirsDrunk++
}

// UseScroll activates the scroll's effect.
func (ch *Character) UseScroll(scroll item.Scroll) {
	ch.MaxHealth += scroll.MaxHealth
	ch.Agility += scroll.Agility
	ch.Strength += scroll.Strength
	ch.Stats.ScrollsRead++
}

// IsEnemyLeft checks whether any enemies remain in the current dungeon level.
func (ch *Character) IsEnemyLeft() {
}

// HitEnemy attempts to hit an enemy and returns whether the hit was successful and the amount of gold earned.
func (ch *Character) HitEnemy(enemy *Enemy) (bool, int) {
	if IsHitSuccessful(&ch.Unit, &enemy.Unit) {
		damage := CalculateDamage(ch)
		ApplyDamage(&enemy.Unit, damage)

		gold := 0
		if enemy.IsDead() {
			gold = enemy.Treasure
			ch.Inventory.Treasure += gold
			ch.Stats.TreasuresReceived += gold
			ch.Stats.EnemiesDefeated++
		}
		ch.Stats.HitsMade++

		return true, gold
	}
	return false, 0
}

// ChooseWeapon equips a weapon if none is currently equipped.
func (ch *Character) ChooseWeapon(weapon *item.Weapon) {
	if ch.CurrentWeapon == nil {
		ch.CurrentWeapon = weapon
		ch.Strength += ch.CurrentWeapon.Strength
	}
}

// DropWeapon - drop weapon in hands to empty tile
func (ch *Character) DropWeapon(d *dungeon.Dungeon) bool {
	if ch.CurrentWeapon == nil {
		return false
	}
	weapon := ch.CurrentWeapon
	ch.Strength -= weapon.Strength
	ch.CurrentWeapon = nil
	if d.AddItemToNearestPosition(*weapon) {
		ch.Inventory.Delete(weapon)
		return true
	}
	return false
}
