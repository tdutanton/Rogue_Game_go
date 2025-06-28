package render

import (
	"fmt"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
)

// RenderMainWindow - draw main game window on screen
func (v *View) RenderMainWindow() {
	startX, startY := 20, 7
	v.MainWindow.Clear()
	v.MainWindow.Box(0, 0)
	logo := []string{
		" /$$$$$$$ ",
		"| $$__  $$",
		"| $$  \\ $$  /$$$$$$   /$$$$$$  /$$   /$$  /$$$$$$",
		"| $$$$$$$/ /$$__  $$ /$$__  $$| $$  | $$ /$$__  $$",
		"| $$__  $$| $$  \\ $$| $$  \\ $$| $$  | $$| $$$$$$$$",
		"| $$  \\ $$| $$  | $$| $$  | $$| $$  | $$| $$_____/",
		"| $$  | $$|  $$$$$$/|  $$$$$$$|  $$$$$$/|  $$$$$$$",
		"|__/  |__/ \\______/  \\____  $$ \\______/  \\_______/",
		"                     /$$  \\ $$",
		"                    |  $$$$$$/ ",
		"                     \\______/ ",
	}

	logoHeight := len(logo)
	v.MainWindow.ColorOn(GreenBlack)
	for i := 0; i < logoHeight; i++ {
		v.MainWindow.MovePrintf(startY+i, startX, logo[i])
	}
	v.MainWindow.ColorOff(GreenBlack)

	menuY, menuX := startY+logoHeight+5, startX+12
	v.MainWindow.MovePrintf(menuY, menuX, `New Game (Press Space)`)
	v.MainWindow.MovePrintf(menuY+1, menuX, `Load Game (Press L)`)
	v.MainWindow.MovePrintf(menuY+2, menuX, `LeaderBoard (Press S)`)
	v.MainWindow.MovePrintf(menuY+3, menuX, `Exit (Press Q)`)

	leftWeapon := []string{
		"  ,:\\      /:.",
		" //  \\_()_/  \\\\",
		"||   |    |   ||",
		"||   |    |   ||",
		"||   |____|   ||",
		" \\   / || \\  //",
		"  `:/  ||  \\:'",
		"       ||",
		"       ||",
		"       XX",
		"       XX",
		"       XX",
		"       XX",
		"       OO",
		"       `'",
	}

	v.MainWindow.ColorOn(YellowBlack)
	leftWeaponY, leftWeaponX := 18, 5
	for i := 0; i < len(leftWeapon); i++ {
		v.MainWindow.MovePrintf(leftWeaponY+i, leftWeaponX, leftWeapon[i])
	}

	rightWeapon := []string{
		"                _",
		"               /\\)",
		"              /\\/",
		"             /\\/",
		"           _/L/",
		"          (/\\_)",
		"          /#/",
		"         /#/",
		"        /#/",
		"       /#/",
		"      /#/",
		"     /#/",
		"    /#/",
		"   /#/",
		"  /#/",
		" /#/",
		"/,'",
		"'",
	}

	rightWeaponY, rightWeaponX := 15, 62
	for i := 0; i < len(rightWeapon); i++ {
		v.MainWindow.MovePrintf(rightWeaponY+i, rightWeaponX, rightWeapon[i])
	}
	v.MainWindow.ColorOff(YellowBlack)

	v.MainWindow.Refresh()
}

// RenderCharacterDeathWindow - draw player's death window on screen
func (v *View) RenderCharacterDeathWindow() {
	v.MainWindow.Clear()
	logo := []string{
		"__   __                                 _                _ ",
		"\\ \\ / /                                | |              | | ",
		" \\ V /___  _   _    __ _ _ __ ___    __| | ___  __ _  __| |  ",
		"  \\ // _ \\| | | |  / _` | '__/ _ \\  / _` |/ _ \\/ _` |/ _` |  ",
		"  | | (_) | |_| | | (_| | | |  __/ | (_| |  __/ (_| | (_| |_ _ _ ",
		"  \\_/\\___/ \\__,_|  \\__,_|_|  \\___|  \\__,_|\\___|\\__,_|\\__,_(_|_|_)",
	}

	v.MainWindow.ColorOn(RedBlack)
	logonY, logoX := 3, 12
	for i := 0; i < len(logo); i++ {
		v.MainWindow.MovePrintf(logonY+i, logoX, logo[i])
	}
	v.MainWindow.ColorOff(RedBlack)

	scull := []string{
		"                               .___.",
		"           /)               ,-^     ^-.",
		"          //               /           \\",
		" .-------| |--------------/  __     __  \\-------------------.__",
		" |WMWMWMW| |>>>>>>>>>>>>> | />>\\   />>\\ |>>>>>>>>>>>>>>>>>>>>>>:>",
		" `-------| |--------------| \\__/   \\__/ |-------------------'^^",
		"          \\\\               \\    /|\\    /",
		"           \\)               \\   \\_/   /",
		"                             |       |",
		"                             |+H+H+H+|",
		"                             \\       /",
		"                              ^-----^",
	}

	scullY, scullX := logonY+len(logo)+5, 10
	for i := 0; i < len(scull); i++ {
		v.MainWindow.MovePrintf(scullY+i, scullX, scull[i])
	}

	v.MainWindow.MovePrintf(scullY+len(scull)+4, 22, `Press Space to return to main menu`)
	v.MainWindow.MovePrintf(scullY+len(scull)+5, 22, `Press Q to quit`)
	v.MainWindow.Refresh()
}

// RenderCharacterWinWindow - draw player's win window on screen
func (v *View) RenderCharacterWinWindow() {
	v.MainWindow.Clear()

	logo := []string{
		" __     __                    _       _ ",
		" \\ \\   / /                   (_)     | |",
		"  \\ \\_/ /__  _   _  __      ___ _ __ | |",
		"   \\   / _ \\| | | | \\ \\ /\\ / / | '_ \\| |",
		"    | | (_) | |_| |  \\ V  V /| | | | |_|",
		"    |_|\\___/ \\__,_|   \\_/\\_/ |_|_| |_(_)",
	}

	v.MainWindow.ColorOn(GreenBlack)
	logoY, logoX := 2, 22
	for i := 0; i < len(logo); i++ {
		v.MainWindow.MovePrintf(logoY+i, logoX, logo[i])
	}
	v.MainWindow.ColorOff(GreenBlack)

	man := []string{
		"   .",
		"  / \\",
		"  |.|",
		"  |:|      __",
		",_|:|_,   /  )",
		"  (Oo    / _I_",
		"   +\\ \\  || __|",
		"      \\ \\||___|",
		"        \\ /.:.\\-\\",
		"         |.:. /-----\\",
		"         |___|::oOo::|",
		"         /   |:<_T_>:|",
		"        |_____\\ ::: /",
		"         | |  \\ \\:/",
		"         | |   | |",
		"         \\ /   | \\___",
		"         / |   \\_____\\",
		"         `-'",
	}

	manY, manX := logoY+len(logo)+2, 29
	for i := 0; i < len(man); i++ {
		v.MainWindow.MovePrintf(manY+i, manX, man[i])
	}

	textY := manY + len(man) + 2
	v.MainWindow.MovePrintf(textY, logoX+2, `Press Space to return to main menu`)
	v.MainWindow.MovePrintf(textY+1, logoX+2, `Press Q to quit`)
	v.MainWindow.Refresh()
}

// RenderStatisticWindow - draw leaderboard on screen
func (v *View) RenderStatisticWindow(stats []common.Stats) {
	startX, startY := 3, 3
	v.MainWindow.Clear()
	v.MainWindow.Box(0, 0)
	sword := []string{
		"              />",
		" ()          //---------------------------------------------------------(",
		"(*)OXOXOXOXO(*>                      LEADERBOARD                         \\",
		" ()          \\\\-----------------------------------------------------------)",
		"              \\>",
	}

	swordY, swordX := 1, 8

	v.MainWindow.ColorOn(GreenBlack)
	for i := 0; i < len(sword); i++ {
		v.MainWindow.MovePrintf(swordY+i, swordX, sword[i])
	}
	v.MainWindow.ColorOff(GreenBlack)

	header := " #   Treasures   Level   Enemies   Food   Elixirs   Scrolls   Hits   Misses   Steps"
	v.MainWindow.MovePrint(startY+4, startX, header)
	startY += 6

	for row, stat := range stats {
		if row >= 21 {
			break
		}

		if row%2 == 0 {
			v.MainWindow.ColorOn(YellowBlack)
		}

		v.MainWindow.MovePrint(startY+row, startX+1, fmt.Sprintf("%d", row+1))
		v.MainWindow.MovePrint(startY+row, startX+7, fmt.Sprintf("%d", stat.TreasuresReceived))
		v.MainWindow.MovePrint(startY+row, startX+18, fmt.Sprintf("%d", stat.LevelAchieved))
		v.MainWindow.MovePrint(startY+row, startX+27, fmt.Sprintf("%d", stat.EnemiesDefeated))
		v.MainWindow.MovePrint(startY+row, startX+36, fmt.Sprintf("%d", stat.FoodEaten))
		v.MainWindow.MovePrint(startY+row, startX+44, fmt.Sprintf("%d", stat.ElixirsDrunk))
		v.MainWindow.MovePrint(startY+row, startX+54, fmt.Sprintf("%d", stat.ScrollsRead))
		v.MainWindow.MovePrint(startY+row, startX+63, fmt.Sprintf("%d", stat.HitsMade))
		v.MainWindow.MovePrint(startY+row, startX+71, fmt.Sprintf("%d", stat.HitsMissed))
		v.MainWindow.MovePrint(startY+row, startX+78, fmt.Sprintf("%d", stat.CellsPassed))

		if row%2 == 0 {
			v.MainWindow.ColorOff(YellowBlack)
		}
	}
	v.MainWindow.MovePrintf(32, 30, "Press any key to continue ...")
	v.MainWindow.Refresh()
}
