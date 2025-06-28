package usecases

import (
	"fmt"

	"github.com/tdutanton/Rogue_Game_go/internal/adapters/secondary/render"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// InventoryActionUseCase handles inventory display and item selection.
type InventoryActionUseCase struct {
	dungeon          *dungeon.Dungeon
	view             render.View
	selectedItemType item.Type
}

// NewInventoryAction creates new inventory use case instance.
func NewInventoryAction(dung *dungeon.Dungeon, view *render.View) *InventoryActionUseCase {
	return &InventoryActionUseCase{
		dungeon: dung,
		view:    *view,
	}
}

// Execute displays inventory items of specified type.
func (uc *InventoryActionUseCase) Execute(itemType item.Type) {
	uc.view.ShowInventory = true
	uc.view.InventoryWindow.Erase()
	defer func() {
		uc.view.ShowInventory = false
		uc.view.InventoryWindow.Erase()
		uc.view.InventoryWindow.Refresh()
		uc.view.Render(*uc.dungeon)
	}()
	player, _ := uc.dungeon.Player.(*unit.Character)
	uc.selectedItemType = item.EmptyType
	uc.view.InventoryWindow.Box(0, 0)
	switch itemType {
	case item.WeaponType:
		uc.view.RenderWeapon(player)
		uc.selectedItemType = item.WeaponType
	case item.ElixirType:
		uc.view.RenderElixirs(player.Inventory.Elixirs)
		uc.selectedItemType = item.ElixirType
	case item.FoodType:
		uc.view.RenderFoods(player.Inventory.Foods)
		uc.selectedItemType = item.FoodType
	case item.ScrollType:
		uc.view.RenderScrolls(player.Inventory.Scrolls)
		uc.selectedItemType = item.ScrollType
	}
	key := uc.view.InventoryWindow.GetChar()
	uc.dungeon.Update()
	switch key {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		uc.Select(int(key - '0'))
	default:
		break
	}
}

// Select chooses and uses inventory item by index.
// Handles weapon equipping, consumable usage, and updates player stats.
func (uc *InventoryActionUseCase) Select(num int) {
	if uc.selectedItemType == item.EmptyType {
		return
	}
	player, _ := uc.dungeon.Player.(*unit.Character)

	switch uc.selectedItemType {
	case item.WeaponType:
		if len(player.Inventory.Weapons) >= num {
			if num == 0 {
				if player.CurrentWeapon != nil {
					player.Strength -= player.CurrentWeapon.Strength
					player.CurrentWeapon = nil
					uc.dungeon.ReplaceEventData("You put the weapon in the backpack")
				}
			} else {
				weapon := &player.Inventory.Weapons[num-1]
				if player.CurrentWeapon == weapon {
					return
				}
				newWeapon := *weapon
				uc.dungeon.ReplaceEventData(fmt.Sprintf("You chose the %s", weapon.Name))
				player.DropWeapon(uc.dungeon)
				player.ChooseWeapon(&newWeapon)
			}
		}
	case item.ElixirType:
		if len(player.Inventory.Elixirs) > num {
			elixir := &player.Inventory.Elixirs[num]
			if elixir.IsActive {
				return
			}
			uc.dungeon.ReplaceEventData(fmt.Sprintf("You drank the %s", elixir.Name))
			player.DrinkElixir(*elixir)
			elixir.IsActive = true
		}
	case item.FoodType:
		if len(player.Inventory.Foods) > num {
			food := &player.Inventory.Foods[num]
			uc.dungeon.ReplaceEventData(fmt.Sprintf("You ate the %s", food.Name))
			player.EatFood(food)
		}
	case item.ScrollType:
		if len(player.Inventory.Scrolls) > num {
			scroll := &player.Inventory.Scrolls[num]
			if scroll.IsActive {
				return
			}
			uc.dungeon.ReplaceEventData(fmt.Sprintf("You read the %s", scroll.Name))
			player.UseScroll(*scroll)
			scroll.IsActive = true
		}
	}
	uc.view.RenderStatistic(*player)
}
