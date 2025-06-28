package render

import (
	"math"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"

	gc "github.com/gbin/goncurses"
)

// RenderGame - draw the dungeon in the console
func (v *View) RenderGame(d dungeon.Dungeon) {
	if v.ShowInventory {
		return
	}
	v.RenderRooms(d.Rooms, d.Player)
	v.RenderPassages(d.Passages)
	v.RenderItems(d)
	v.RenderEnemy(d)
	v.RenderPlayer(d)

	v.RenderExit(d)
	v.GameWindow.Refresh()
}

// RenderRooms - draw rooms in the dungeon
func (v *View) RenderRooms(rooms [common.MaxRoomCount]dungeon.Room, player dungeon.Coordinator) {
	for _, room := range rooms {
		if room.Visited {
			v.RenderRoom(room, player)
		}
	}
}

// FillRoomByFog - fill every room tile as closed from character's view
func (v *View) FillRoomByFog(room dungeon.Room) {
	startX := room.X
	startY := room.Y
	endX := room.X + room.Width - 1
	endY := room.Y + room.Height - 1

	for y := startY + 1; y < endY; y++ {
		for x := startX + 1; x < endX; x++ {
			v.draw(y, x, Fog, GreenBlack)
		}
	}
}

// ClearRoomFromFog - set every room tile as clean tile (empty)
func (v *View) ClearRoomFromFog(room dungeon.Room) {
	startX := room.X
	startY := room.Y
	endX := room.X + room.Width - 1
	endY := room.Y + room.Height - 1

	for y := startY + 1; y < endY; y++ {
		for x := startX + 1; x < endX; x++ {
			v.draw(y, x, EmptyFloor, GreenBlack)
		}
	}
}

// FillRoomByPartFog fill the room by fog around the view sector
func (v *View) FillRoomByPartFog(room dungeon.Room, player dungeon.Coordinator) {
	playerCoords := player.GetCoords()
	playerX := playerCoords.X
	playerY := playerCoords.Y

	isVertical := room.Visible == dungeon.VerticalFog
	ratio := 1.5
	if !isVertical {
		ratio = 3.5
	}

	startX := room.X + 1
	startY := room.Y + 1
	endX := room.X + room.Width - 2
	endY := room.Y + room.Height - 2

	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			newX := x - playerX
			newY := y - playerY
			absX := math.Abs(float64(newX))
			absY := math.Abs(float64(newY))

			visible := (isVertical && absY >= ratio*absX) ||
				(!isVertical && absX >= ratio*absY)

			if visible {
				v.draw(y, x, EmptyFloor, GreenBlack)
			} else {
				v.draw(y, x, Fog, GreenBlack)
			}
		}
	}
}

func (v *View) isCoordVisible(coord common.Coords, player dungeon.Coordinator, room dungeon.Room) bool {
	playerCoords := player.GetCoords()
	newX := coord.X - playerCoords.X
	newY := coord.Y - playerCoords.Y
	absX := math.Abs(float64(newX))
	absY := math.Abs(float64(newY))

	isVertical := room.Visible == dungeon.VerticalFog
	ratio := 1.5
	if !isVertical {
		ratio = 3.5
	}

	return (isVertical && absY >= ratio*absX) || (!isVertical && absX >= ratio*absY)
}

// RenderRoom - draw one room
func (v *View) RenderRoom(room dungeon.Room, player dungeon.Coordinator) {
	startX := room.X
	startY := room.Y
	endX := room.X + room.Width - 1
	endY := room.Y + room.Height - 1

	switch room.Visible {
	case dungeon.VerticalFog, dungeon.HorizontalFog:
		v.FillRoomByPartFog(room, player)
	case dungeon.FogClean:
		v.ClearRoomFromFog(room)
	default:
		v.FillRoomByFog(room)
	}

	for x := startX + 1; x < endX; x++ {
		v.draw(startY, x, WallHorizontal, YellowBlack)
		v.draw(endY, x, WallHorizontal, YellowBlack)
	}

	for y := startY + 1; y < endY; y++ {
		v.draw(y, startX, WallVertical, YellowBlack)
		v.draw(y, endX, WallVertical, YellowBlack)
	}

	v.draw(startY, startX, LowerLeftCorner, YellowBlack)
	v.draw(startY, endX, UpperRightCorner, YellowBlack)
	v.draw(endY, startX, UpperLeftCorner, YellowBlack)
	v.draw(endY, endX, LowerRightCorner, YellowBlack)

	for _, room := range room.Doors {
		v.draw(room.Y, room.X, Door, YellowBlack)
	}
}

// RenderPassages - draw passages slice
func (v *View) RenderPassages(passages []dungeon.Passage) {
	for _, passage := range passages {
		v.RenderPassage(passage)
	}
}

// RenderPassage - draw one passage
func (v *View) RenderPassage(passage dungeon.Passage) {
	for _, path := range passage.Path {
		if !path.Visited {
			continue
		}
		if path.Begin.X == path.End.X {
			startY := min(path.Begin.Y, path.End.Y)
			endY := max(path.Begin.Y, path.End.Y)
			for y := startY; y <= endY; y++ {
				v.GameWindow.MoveAddChar(y, path.Begin.X, Passage)
			}
		} else {
			startX := min(path.Begin.X, path.End.X)
			endX := max(path.Begin.X, path.End.X)
			for x := startX; x <= endX; x++ {
				v.GameWindow.MoveAddChar(path.Begin.Y, x, Passage)
			}
		}
	}
}

// RenderItems - draw items slice on the map
func (v *View) RenderItems(d dungeon.Dungeon) {
	currentRoom := d.CurrentRoomWithWalls()
	if currentRoom == nil {
		return
	}

	for _, i := range d.Items {
		coords := i.GetCoords()

		if !currentRoom.Contains(coords) ||
			(currentRoom.Visible != dungeon.FogClean &&
				!v.isCoordVisible(coords, d.Player, *currentRoom)) {
			continue
		}

		switch i.Type() {
		case item.FoodType:
			v.GameWindow.MoveAddChar(coords.Y, coords.X, Food)
		case item.ElixirType:
			v.GameWindow.MoveAddChar(coords.Y, coords.X, Elixir)
		case item.WeaponType:
			v.GameWindow.MoveAddChar(coords.Y, coords.X, Weapon)
		case item.ScrollType:
			v.GameWindow.MoveAddChar(coords.Y, coords.X, Scroll)
		default:
			break
		}
	}
}

// RenderEnemy - draw one enemy
func (v *View) RenderEnemy(d dungeon.Dungeon) {
	currentRoom := d.CurrentRoomWithWalls()
	currentPassage := d.CurrentPassage()

	for _, u := range d.Enemies {
		enemy, ok := u.(*unit.Enemy)
		if !ok {
			continue
		}

		coords := enemy.GetCoords()

		if currentRoom != nil {
			if !currentRoom.Contains(coords) || (currentRoom.Visible != dungeon.FogClean && !v.isCoordVisible(coords, d.Player, *currentRoom)) {
				continue
			}
		} else if currentPassage != nil {
			if !currentPassage.Contains(coords) {
				continue
			}
		}

		switch enemy.EnemyType {
		case unit.Zombie:
			v.draw(coords.Y, coords.X, Zombie, GreenBlack)
		case unit.Vampire:
			v.draw(coords.Y, coords.X, Vampire, RedBlack)
		case unit.Ghost:
			if enemy.Visibility {
				v.GameWindow.MoveAddChar(coords.Y, coords.X, Ghost)
			}
		case unit.Ogr:
			v.draw(coords.Y, coords.X, Ogr, YellowBlack)
		case unit.SnakeWizard:
			v.GameWindow.MoveAddChar(coords.Y, coords.X, SnakeWizard)
		default:
			break
		}
	}
}

// RenderPlayer - draw the Character
func (v *View) RenderPlayer(d dungeon.Dungeon) {
	coords := d.Player.GetCoords()
	if _, isCoridor := d.TileFromPassages(coords); isCoridor {
		v.draw(coords.Y, coords.X, Character, BlackWhite)
	} else {
		v.draw(coords.Y, coords.X, Character, WhiteBlack)
	}
}

// RenderExit - draw Exit point
func (v *View) RenderExit(d dungeon.Dungeon) {
	coords := d.Exit
	currentRoom := d.CurrentRoomWithWalls()
	if currentRoom == nil || !currentRoom.Contains(coords) {
		return
	}

	v.GameWindow.AttrOn(gc.A_BLINK)
	v.draw(d.Exit.Y, d.Exit.X, Exit, GreenBlack)
	v.GameWindow.AttrOff(gc.A_BLINK)
}
