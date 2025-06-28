// Package app handles terminal-based UI rendering using the goncurses library.
// It initializes the screen, creates windows for different game interface components,
// and sets up input handling for the roguelike game.
package app

import (
	"context"
	"log"
	"os"

	"github.com/tdutanton/Rogue_Game_go/internal/adapters/secondary/render"
	"github.com/tdutanton/Rogue_Game_go/internal/adapters/secondary/storage"
	"github.com/tdutanton/Rogue_Game_go/internal/application/usecases"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"

	"github.com/tdutanton/Rogue_Game_go/internal/adapters/primary/input"
	"github.com/tdutanton/Rogue_Game_go/internal/services/config"
	"github.com/tdutanton/Rogue_Game_go/pkg/logger"

	gc "github.com/gbin/goncurses"
)

// Run initializes the game environment, loads configuration, creates windows and UI components,
// sets up the dungeon, and starts the main game loop, using some help functions.
func Run() {
	initGoNcurses()
	defer gc.End()

	windows := createAllWindows()
	view := render.NewView(
		windows.Info,
		windows.Statistic,
		windows.Game,
		windows.Inventory,
		windows.Main,
	)

	cfgGame := loadDungeonConfig()
	d := storage.GenerateDungeonFromConfig(1, cfgGame, nil)
	character := getCharacterFromDungeon(d)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	playerActionUC := usecases.NewPlayerActionUseCase(character, d, view, cfgGame, cancel)
	inventoryActionUC := usecases.NewInventoryAction(playerActionUC.GetDungeon(), view)

	inputHandler := input.NewInputHandler(
		windows.Input,
		cancel,
		playerActionUC,
		inventoryActionUC,
		windows.Main,
	)

	view.RenderMainWindow()
	gameLoop(ctx, inputHandler)
}

// windows struct with all game windows
type windows struct {
	Input     *gc.Window
	Info      *gc.Window
	Statistic *gc.Window
	Game      *gc.Window
	Inventory *gc.Window
	Main      *gc.Window
}

// createAllWindows just create all windows (surprise!)
func createAllWindows() windows {
	inputWindow, _ := CreateInputWindow()
	infoWindow, _ := CreateInfoWindow()
	statisticWindow, _ := CreateStatisticWindow()
	gameWindow, _ := CreateGameWindow()
	inventoryWindow, _ := CreateInventoryWindow()
	mainWindow, _ := CreateMainWindow()

	return windows{
		Input:     inputWindow,
		Info:      infoWindow,
		Statistic: statisticWindow,
		Game:      gameWindow,
		Inventory: inventoryWindow,
		Main:      mainWindow,
	}
}

// loadDungeonConfig - get .yaml
func loadDungeonConfig() *storage.Config {
	cfgGame, err := storage.LoadDungeonConfig("configs/dungeon_config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfgGame
}

// getCharacterFromDungeon - just get Player's pointer at the Game
func getCharacterFromDungeon(d dungeon.Dungeon) *unit.Character {
	character, ok := d.Player.(*unit.Character)
	if !ok {
		panic("player is not of type *unit.Character")
	}
	return character
}

// gameLoop runs the main loop that listens for input or exit signals.
// Exits when the context is canceled (e.g., player dies or quits).
func gameLoop(ctx context.Context, inputHandler *input.InputHandler) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			inputHandler.Update()
		}
	}
}

// RunDebug initializes the game environment with debug logger, loads configuration, creates windows and UI components,
// sets up the dungeon, and starts the main game loop.
func RunDebug() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is empty")
	}

	cfg := config.MustLoadConfig(configPath)

	logger.SetSettings(cfg.LoggerInfo)

	initGoNcurses()
	defer gc.End()

	inputWindow, _ := CreateInputWindow()
	infoWindow, _ := CreateInfoWindow()
	statisticWindow, _ := CreateStatisticWindow()
	gameWindow, _ := CreateGameWindow()
	inventoryWindow, _ := CreateInventoryWindow()
	mainWindow, _ := CreateMainWindow()
	view := render.NewView(infoWindow, statisticWindow, gameWindow, inventoryWindow, mainWindow)

	cfgGame, err := storage.LoadDungeonConfig("configs/dungeon_config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	level := 1

	d := storage.GenerateDungeonFromConfig(level, cfgGame, nil)
	character, ok := d.Player.(*unit.Character)
	if !ok {
		panic("player is not of type *unit.Character")
	}

	ctx, cancel := context.WithCancel(context.Background())
	playerActionUC := usecases.NewPlayerActionUseCase(character, d, view, cfgGame, cancel)
	inventoryActionUC := usecases.NewInventoryAction(playerActionUC.GetDungeon(), view)

	inputHandler := input.NewInputHandler(inputWindow, cancel, playerActionUC, inventoryActionUC, mainWindow)

	view.RenderMainWindow()
	gameLoop(ctx, inputHandler)
}
