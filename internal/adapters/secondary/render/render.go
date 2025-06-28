package render

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"

	gc "github.com/gbin/goncurses"
)

// View represents the game's visual interface components.
// It manages all ncurses windows used to display different aspects of the game:
// - GameWindow: Main game area showing dungeon and entities
// - InfoWindow: Displays game events and messages
// - StatisticWindow: Shows player stats and character information
// - InventoryWindow: Displays player's inventory items
//
// The View handles all rendering operations and maintains color configurations.
type View struct {
	InfoWindow      *gc.Window // Window for game event messages
	StatisticWindow *gc.Window // Window for player statistics
	GameWindow      *gc.Window // Main game area window
	InventoryWindow *gc.Window // Inventory display window
	MainWindow      *gc.Window // MainWindow of Game
	ShowInventory   bool       // ShowInventory - check if inventory window is open now
}

// NewView creates and initializes a new View instance with the specified windows.
// It sets up the color pairs used throughout the game interface.
//
// Parameters:
//   - infoWindow: Window for displaying game messages and events
//   - StatisticWindow: Window for character statistics display
//   - gameWindow: Main game rendering window
//   - inventoryWindow: Window for inventory management
//
// Returns:
//   - *View: Initialized View instance with configured colors
func NewView(infoWindow *gc.Window, StatisticWindow *gc.Window, gameWindow *gc.Window, inventoryWindow *gc.Window, mainWindow *gc.Window) *View {
	view := &View{
		InfoWindow:      infoWindow,
		StatisticWindow: StatisticWindow,
		GameWindow:      gameWindow,
		InventoryWindow: inventoryWindow,
		MainWindow:      mainWindow,
	}
	view.adjustColors()

	return view
}

// Render updates all game view components with current game state.
// This is the main rendering function that coordinates updates to:
// - Inventory display
// - Game event messages
// - Main game area
// - Player statistics
//
// Parameters:
//   - d: Current dungeon state containing all game entities and events
func (v *View) Render(d dungeon.Dungeon) {
	if v.ShowInventory {
		v.InventoryWindow.Erase()
		v.InventoryWindow.Refresh()
	}
	v.RenderInfo(d.EventData)
	v.RenderGame(d)
	player, _ := d.Player.(*unit.Character)
	v.RenderStatistic(*player)
	if !v.ShowInventory {
		v.InventoryWindow.Erase()
	}
	gc.UpdatePanels()
}

// adjustColors initializes all color pairs used in the game interface.
// This private method is called during View initialization and sets up:
// - White on black (default text)
// - Black on white (inverted)
// - Red on black (warnings/danger)
// - Green on black (positive status)
// - Green on white (highlighted positive)
// - Yellow on black (special items/messages)
func (v *View) adjustColors() {
	gc.StartColor()
	gc.InitPair(WhiteBlack, gc.C_WHITE, gc.C_BLACK)
	gc.InitPair(BlackWhite, gc.C_BLACK, gc.C_WHITE)
	gc.InitPair(RedBlack, gc.C_RED, gc.C_BLACK)
	gc.InitPair(GreenBlack, gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(GreenWhite, gc.C_GREEN, gc.C_WHITE)
	gc.InitPair(YellowBlack, gc.C_YELLOW, gc.C_BLACK)
	gc.InitPair(BlueBlack, gc.C_BLUE, gc.C_BLACK)
}

// draw renders a single character at specified coordinates with given color.
// This is a low-level drawing primitive used by other rendering functions.
//
// Parameters:
//   - y: Vertical position (row) in the game window
//   - x: Horizontal position (column) in the game window
//   - symbol: Character to display
//   - color: Color pair to use (must be initialized via adjustColors)
func (v *View) draw(y int, x int, symbol gc.Char, color int16) {
	v.GameWindow.ColorOn(color)
	v.GameWindow.MoveAddChar(y, x, symbol)
	v.GameWindow.ColorOff(color)
}
