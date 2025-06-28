package logic

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// PlayerMove handle player's moving in the specified direction within the dungeon.
// It checks if the move is possible before updating the player's coordinates.
// If the move is valid, the player's position is updated according to the given direction.
func PlayerMove(dir unit.Direction, dg *dungeon.Dungeon) {
	player := dg.Player.(*unit.Character)

	if unit.IsPossiblePlayerMove(GetCoordsAfterMoving(player.Coords, dir), *dg) {
		switch dir {
		case unit.Up:
			player.MoveUp()
		case unit.Down:
			player.MoveDown()
		case unit.Left:
			player.MoveLeft()
		case unit.Right:
			player.MoveRight()
		}
	}
}

// CheckConsumables checks if the player is standing on any items in the dungeon
// and attempts to add them to the player's inventory.
func CheckConsumables(dg *dungeon.Dungeon) {
	player := dg.Player.(*unit.Character)
	for i := len(dg.Items) - 1; i >= 0; i-- {
		it := dg.Items[i]
		if it.GetCoords() == dg.PlayerCoords() {
			if player.Inventory.Add(it) {
				dg.AddEventData("You take the " + item.ItemsNames[it.Type()] + "!")
				// !удаляется предмет из dg.Items, но должен ли?...
				dg.Items = append(dg.Items[:i], dg.Items[i+1:]...)
			}
		}
	}
}
