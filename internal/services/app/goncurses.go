package app

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"

	gc "github.com/gbin/goncurses"
)

// initGoNcurses initializes the ncurses environment with basic settings:
// - Disables echo (input is not printed to the screen)
// - Hides cursor
// - Enables cbreak mode (no line buffering)
// - Enables raw mode (direct character input)
func initGoNcurses() {
	_, err := gc.Init()
	if err != nil {
		panic(err)
	}

	gc.Echo(false)
	gc.Cursor(0)
	gc.CBreak(true)
	gc.Raw(true)
}

// CreateInputWindow creates a window for user input at the bottom of the screen.
// Returns a pointer to the created window or nil if an error occurred.
func CreateInputWindow() (*gc.Window, error) {
	win, err := gc.NewWindow(1, 1, common.InfoHeight+common.MapHeight+common.StatisticHeight, 0)
	if err != nil {
		return nil, err
	}

	return win, nil
}

// CreateInfoWindow creates a window for displaying short messages and logs.
// Positioned at the top of the screen.
// Returns a pointer to the created window or nil if an error occurred.
func CreateInfoWindow() (*gc.Window, error) {
	win, err := gc.NewWindow(common.InfoHeight, common.InfoWidth, 0, 0)
	if err != nil {
		return nil, err
	}

	return win, nil
}

// CreateGameWindow creates a window for rendering the dungeon map.
// Positioned below the info window.
// Returns a pointer to the created window or nil if an error occurred.
func CreateGameWindow() (*gc.Window, error) {
	win, err := gc.NewWindow(common.MapHeight+1, common.MapWidth+1, common.InfoHeight, 0)
	if err != nil {
		return nil, err
	}

	return win, nil
}

// CreateStatisticWindow creates a window for displaying player statistics.
// Positioned below the game window.
// Returns a pointer to the created window or nil if an error occurred.
func CreateStatisticWindow() (*gc.Window, error) {
	win, err := gc.NewWindow(common.StatisticHeight, common.StatisticWidth, common.InfoHeight+common.MapHeight+1, 0)
	if err != nil {
		return nil, err
	}

	return win, nil
}

// CreateInventoryWindow creates a window for displaying the player's inventory.
// Positioned to the right of the game window.
// Returns a pointer to the created window or nil if an error occurred.
func CreateInventoryWindow() (*gc.Window, error) {
	win, err := gc.NewWindow(common.InventoryHeight, common.InventoryWidth, common.InfoHeight, 0)
	if err != nil {
		return nil, err
	}

	return win, nil
}

// CreateMainWindow creates the main window of the game
func CreateMainWindow() (*gc.Window, error) {
	win, err := gc.NewWindow(common.MainHeight, common.MainWidth, 0, 0)
	if err != nil {
		return nil, err
	}

	return win, nil
}
