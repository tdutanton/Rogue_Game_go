package render

import (
	"fmt"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// RenderElixirs displays the player's elixir inventory in the inventory window.
// Clears the window and shows either "Empty elixirs" or a numbered list of elixirs.
// Each elixir is displayed in the format: "index.name" (e.g., "0.Potion of Healing").
//
// Parameters:
//   - elixirs: Slice of Elixir items to display. If empty or nil, shows empty message.
//
// Display Format:
//
//	Empty elixirs
//	OR
//	0.Elixir of Power
//	1.Potion of Invisibility
func (v *View) RenderElixirs(elixirs []item.Elixir) {
	startX, startY := 10, 10
	v.InventoryWindow.MovePrintf(startY, startX, "Choose elixir:")

	if elixirs == nil || len(elixirs) == 0 {
		v.InventoryWindow.MovePrintf(startY+2, startX, "You haven't elixirs!")
	}
	for i, elixir := range elixirs {
		if elixir.IsActive {
			v.InventoryWindow.MovePrintf(startY+i+2, startX, "* ")
		} else {
			v.InventoryWindow.MovePrintf(startY+i+2, startX, "  ")
		}

		if elixir.Agility > 0 {
			v.InventoryWindow.MovePrintf(startY+i+2, startX+2, fmt.Sprintf("%d.%s (+%d agility for %d steps)", i, elixir.Name, elixir.Agility, elixir.Duration))
		} else if elixir.MaxHealth > 0 {
			v.InventoryWindow.MovePrintf(startY+i+2, startX+2, fmt.Sprintf("%d.%s (+%d max health for %d steps)", i, elixir.Name, elixir.MaxHealth, elixir.Duration))
		} else if elixir.Strength > 0 {
			v.InventoryWindow.MovePrintf(startY+i+2, startX+2, fmt.Sprintf("%d.%s (+%d strength for %d steps)", i, elixir.Name, elixir.Strength, elixir.Duration))
		}
	}
	v.InventoryWindow.MovePrintf(startY+len(elixirs)+5, startX, "Press any key to continue ...")
}

// RenderScrolls displays the player's scroll inventory in the inventory window.
// Clears the window and shows either "Empty scrolls" or a numbered list of scrolls.
// Each scroll is displayed in the format: "index.name" (e.g., "1.Scroll of Fireball").
//
// Parameters:
//   - scrolls: Slice of Scroll items to display. If empty or nil, shows empty message.
//
// Display Format:
//
//	Empty scrolls
//	OR
//	0.Scroll of Identify
//	1.Scroll of Teleportation
func (v *View) RenderScrolls(scrolls []item.Scroll) {
	startX, startY := 10, 10
	v.InventoryWindow.MovePrintf(startY, startX, "Choose scroll:")
	if scrolls == nil || len(scrolls) == 0 {
		v.InventoryWindow.MovePrintf(startY+2, startX, "You haven't scrolls!")
	}
	for i, scroll := range scrolls {
		if scroll.IsActive {
			v.InventoryWindow.MovePrintf(startY+i+2, startX, "* ")
		} else {
			v.InventoryWindow.MovePrintf(startY+i+2, startX, "  ")
		}

		if scroll.Agility > 0 {
			v.InventoryWindow.MovePrintf(startY+i+2, startX+2, fmt.Sprintf("%d.%s (+%d agility for %d steps)", i, scroll.Name, scroll.Agility, scroll.Duration))
		} else if scroll.MaxHealth > 0 {
			v.InventoryWindow.MovePrintf(startY+i+2, startX+2, fmt.Sprintf("%d.%s (+%d max health for %d steps)", i, scroll.Name, scroll.MaxHealth, scroll.Duration))
		} else if scroll.Strength > 0 {
			v.InventoryWindow.MovePrintf(startY+i+2, startX+2, fmt.Sprintf("%d.%s (+%d strength for %d steps)", i, scroll.Name, scroll.Strength, scroll.Duration))
		}
	}
	v.InventoryWindow.MovePrintf(startY+len(scrolls)+5, startX, "Press any key to continue ...")
}

// RenderFoods displays the player's food inventory in the inventory window.
// Clears the window and shows either "Empty foods" or a numbered list of food items.
// Each food item is displayed in the format: "index.name" (e.g., "2.Apple").
//
// Parameters:
//   - foods: Slice of Food items to display. If empty or nil, shows empty message.
//
// Display Format:
//
//	Empty foods
//	OR
//	0.Bread
//	1.Cheese
func (v *View) RenderFoods(foods []item.Food) {
	startX, startY := 10, 10
	v.InventoryWindow.MovePrintf(startY, startX, "Choose food:")
	if foods == nil || len(foods) == 0 {
		v.InventoryWindow.MovePrintf(startY+2, startX, "You haven't food!")
	}
	for i, food := range foods {
		v.InventoryWindow.MovePrintf(startY+i+2, startX, fmt.Sprintf("%d.%s (+%d health)", i, food.Name, food.Value))
	}

	v.InventoryWindow.MovePrintf(startY+len(foods)+5, startX, "Press any key to continue ...")
}

// RenderWeapon displays the player's weapon inventory with special handling.
// Shows "Empty weapons" if no weapons available, otherwise displays:
// - Option 0: "Without weapon" (default no-weapon state)
// - Subsequent options: numbered list of available weapons (starting from 1)
//
// Parameters:
//   - weapons: Slice of Weapon items to display. Empty slice shows empty message.
//
// Display Format:
//
//	Empty weapons
//	OR
//	0.Without weapon
//	1.Short Sword
//	2.Long Bow
func (v *View) RenderWeapon(ch *unit.Character) {
	startX, startY := 10, 10
	weapons := ch.Inventory.Weapons
	v.InventoryWindow.MovePrintf(startY, startX, "Choose weapon:")
	if weapons == nil || len(weapons) == 0 {
		v.InventoryWindow.MovePrintf(startY+2, startX, "You haven't weapons!")
	} else {
		if ch.CurrentWeapon == nil {
			v.InventoryWindow.MovePrintf(startY+2, startX, "> 0.Without weapon (+0)")
		} else {
			v.InventoryWindow.MovePrintf(startY+2, startX, "  0.Without weapon (+0)")
		}
	}
	for i, weapon := range weapons {
		if ch.CurrentWeapon != nil && weapon == *ch.CurrentWeapon {
			v.InventoryWindow.MovePrintf(startY+i+3, startX, "> ")
		} else {
			v.InventoryWindow.MovePrintf(startY+i+3, startX, "  ")
		}
		v.InventoryWindow.MovePrintf(startY+i+3, startX+2, fmt.Sprintf("%d.%s (+%d)", i+1, weapon.Name, weapon.Strength))
	}

	v.InventoryWindow.MovePrintf(startY+len(weapons)+5, startX, "Press any key to continue ...")
}
