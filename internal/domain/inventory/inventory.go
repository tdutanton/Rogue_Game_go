// Package inventory manages the player's collected items, including weapons, consumables,
// and treasure. It enforces limits and provides methods for adding and removing items.
package inventory

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
)

// MaxItems defines the maximum number of items allowed in each category (weapons, elixirs, etc.)
const MaxItems = 9

// Inventory holds all collectible items the player possesses.
type Inventory struct {
	Weapons  []item.Weapon // Equippable weapons
	Elixirs  []item.Elixir // Temporary stat-boosting potions
	Scrolls  []item.Scroll // Magic scrolls with one-time effects
	Foods    []item.Food   // Food items that restore health
	Treasure int           // Collected gold or currency
}

// Add attempts to add an item to the inventory.
// Returns false if the inventory is full for the given item type.
func (inv *Inventory) Add(itm item.Item) bool {
	switch itm.Type() {
	case item.WeaponType:
		if len(inv.Weapons) >= MaxItems {
			return false
		}

		weapon, _ := itm.(*item.Weapon)
		inv.Weapons = append(inv.Weapons, *weapon)

	case item.ElixirType:
		if len(inv.Weapons) >= MaxItems {
			return false
		}

		elixir, _ := itm.(*item.Elixir)
		inv.Elixirs = append(inv.Elixirs, *elixir)

	case item.ScrollType:
		if len(inv.Weapons) >= MaxItems {
			return false
		}

		scroll, _ := itm.(*item.Scroll)
		inv.Scrolls = append(inv.Scrolls, *scroll)

	case item.FoodType:
		if len(inv.Weapons) >= MaxItems {
			return false
		}

		food, _ := itm.(*item.Food)
		inv.Foods = append(inv.Foods, *food)

	default:
		break
	}

	return true
}

// Delete removes a specific item from the inventory.
// Returns true if the item was found and removed, false otherwise.
func (inv *Inventory) Delete(itm item.Item) bool {
	switch v := itm.(type) {
	case *item.Weapon:
		return deleteFromSlice(&inv.Weapons, *v)
	case *item.Elixir:
		return deleteFromSlice(&inv.Elixirs, *v)
	case *item.Scroll:
		return deleteFromSlice(&inv.Scrolls, *v)
	case *item.Food:
		return deleteFromSlice(&inv.Foods, *v)
	default:
		return false
	}
}

// deleteFromSlice is a generic helper function to remove an item from a slice.
// Returns true if the item was found and removed.
func deleteFromSlice[T comparable](slice *[]T, item T) bool {
	for i, v := range *slice {
		if v == item {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return true
		}
	}
	return false
}
