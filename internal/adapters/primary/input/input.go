package input

import (
	"context"

	"github.com/tdutanton/Rogue_Game_go/internal/adapters/secondary/storage"
	"github.com/tdutanton/Rogue_Game_go/internal/application/usecases"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"

	gc "github.com/gbin/goncurses"
)

// AppState represent the current game state to switch between screens
type AppState int

const (
	// AppStateMainMenu shows main menu
	AppStateMainMenu AppState = iota
	// AppStateInGame shows game screen
	AppStateInGame
	// AppStateGameOver shows screen if player's dead
	AppStateGameOver
	// AppStateWin shows screen if player's win
	AppStateWin
)

// InputHandler handles user input and translates it into game actions.
// It manages keyboard input and delegates actions to the appropriate use cases.
type InputHandler struct {
	stdscr            *gc.Window                       // ncurses standard screen window
	cancel            context.CancelFunc               // function to cancel the game context
	playerActionUC    *usecases.PlayerActionUseCase    // use case for player movements
	inventoryActionUC *usecases.InventoryActionUseCase // use case for inventory operations
	appState          AppState                         // current AppState
	MainWindow        *gc.Window                       // Start Screen
}

// NewInputHandler creates a new InputHandler instance.
//
// Parameters:
//   - stdscr: ncurses standard screen window
//   - cancel: context cancellation function to terminate the game
//   - playerActionUC: use case for handling player movement actions
//   - inventoryActionUC: use case for handling inventory actions
//
// Returns:
//   - *InputHandler: initialized input handler instance
func NewInputHandler(
	stdscr *gc.Window,
	cancel context.CancelFunc,
	playerActionUC *usecases.PlayerActionUseCase,
	inventoryActionUC *usecases.InventoryActionUseCase,
	mainWindow *gc.Window,
) *InputHandler {
	return &InputHandler{
		stdscr:            stdscr,
		cancel:            cancel,
		playerActionUC:    playerActionUC,
		inventoryActionUC: inventoryActionUC,
		appState:          AppStateMainMenu,
		MainWindow:        mainWindow,
	}
}

// Update processes the current input frame.
// It reads keyboard input and triggers corresponding game actions:
//   - Movement (WASD or arrow keys)
//   - Inventory operations (h,j,k,e keys)
//   - Item selection (number keys 0-9)
//   - Game termination (Ctrl+C)
//
// The method uses a 100ms timeout for non-blocking input.
func (h *InputHandler) Update() {
	h.stdscr.Timeout(100)
	key := h.stdscr.GetChar()

	switch h.appState {
	case AppStateMainMenu:
		switch key {
		case ' ', gc.KEY_ENTER:
			h.startNewGame()
		case 'l', 'L':
			storage.RemoveLastRecord()
			_, err := h.playerActionUC.LoadGame()
			if err != nil {
				h.MainWindow.MovePrintf(2, 12, "Failed to load game - there's no save file. Please start new game")
				h.MainWindow.Refresh()
			} else {
				h.appState = AppStateInGame
				h.playerActionUC.RenderInitial()
			}
		case 'q', 'Q':
			h.cancel()
		case 's', 'S':
			h.showLeaderBoard()
		case 0x03: // Ctrl+C
			h.cancel()
		}
	case AppStateInGame:
		switch key {
		case 'a', gc.KEY_LEFT:
			result := h.playerActionUC.Execute(unit.Left)
			h.handleActionResult(result)
		case 's', gc.KEY_DOWN:
			result := h.playerActionUC.Execute(unit.Down)
			h.handleActionResult(result)
		case 'w', gc.KEY_UP:
			result := h.playerActionUC.Execute(unit.Up)
			h.handleActionResult(result)
		case 'd', gc.KEY_RIGHT:
			result := h.playerActionUC.Execute(unit.Right)
			h.handleActionResult(result)
		case 'h':
			h.inventoryActionUC.Execute(item.WeaponType)
		case 'j':
			h.inventoryActionUC.Execute(item.FoodType)
		case 'k':
			h.inventoryActionUC.Execute(item.ElixirType)
		case 'e':
			h.inventoryActionUC.Execute(item.ScrollType)
		case 'q', 'Q':
			h.playerActionUC.SaveGame()
			h.cancel()
		case 0x03: // Ctrl+C
			h.cancel()
		}
	case AppStateGameOver, AppStateWin:
		switch key {
		case ' ':
			h.returnToMainMenu()
		case 'q', 'Q':
			h.cancel()
		case 0x03: // Ctrl+C
			h.cancel()
		}
	}
}

// SetAppState - switch current gamestate to new state
func (h *InputHandler) SetAppState(state AppState) {
	h.appState = state
}

// startNewGame - action in main menu if player chose start new game
func (h *InputHandler) startNewGame() {
	h.appState = AppStateInGame
	h.MainWindow.Erase()
	h.MainWindow.Refresh()
	h.playerActionUC.NewGame()
	h.playerActionUC.RenderInitial()
}

// showLeaderBoard - action in main menu if player chose look to statistics
func (h *InputHandler) showLeaderBoard() {
	// h.MainWindow.Erase()
	// h.MainWindow.Refresh()

	h.playerActionUC.RenderLeaderBord()
}

// handleActionResult check and handle different game states
func (h *InputHandler) handleActionResult(result usecases.ActionResult) {
	switch result {
	case usecases.GameOver:
		h.appState = AppStateGameOver
	case usecases.Win:
		h.appState = AppStateWin
	case usecases.NextLevel:
		h.playerActionUC.RenderInitial()
	}
}

// returnToMainMenu - get the main menu if player die or win
func (h *InputHandler) returnToMainMenu() {
	h.appState = AppStateMainMenu
	h.playerActionUC.RenderMainMenu()
}
