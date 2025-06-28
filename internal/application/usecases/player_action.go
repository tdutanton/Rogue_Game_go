package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/tdutanton/Rogue_Game_go/internal/adapters/secondary/render"
	"github.com/tdutanton/Rogue_Game_go/internal/adapters/secondary/storage"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/logic"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// ActionResult represent current state to pass it into Input
type ActionResult int

const (
	// ContinueGame - Player can move
	ContinueGame ActionResult = iota
	// GameOver - Player lose
	GameOver
	// Win - Player Win
	Win
	// NextLevel - Player step on Exit tile
	NextLevel
)

// PlayerActionUseCase encapsulates the logic for processing player input and managing dungeon state.
type PlayerActionUseCase struct {
	character *unit.Character
	dungeon   dungeon.Dungeon
	view      render.View
	cfg       *storage.Config
	cancel    context.CancelFunc
}

// NewPlayerActionUseCase creates a new instance of PlayerActionUseCase with provided dependencies.
func NewPlayerActionUseCase(character *unit.Character, dung dungeon.Dungeon, view *render.View, cfg *storage.Config, cancel context.CancelFunc) *PlayerActionUseCase {
	return &PlayerActionUseCase{
		character: character,
		dungeon:   dung,
		view:      *view,
		cfg:       cfg,
		cancel:    cancel,
	}
}

// GetDungeon returns a pointer to the current dungeon instance.
func (uc *PlayerActionUseCase) GetDungeon() *dungeon.Dungeon {
	return &uc.dungeon
}

// Execute processes the player's directional input, updates dungeon state, handles rendering,
// checks for death or level completion, and progresses to the next level if needed.
func (uc *PlayerActionUseCase) Execute(direction unit.Direction) ActionResult {
	logic.HandleAction(direction, &uc.dungeon)
	uc.dungeon.Update()
	if uc.character.IsDead() {
		uc.view.RenderCharacterDeathWindow()
		uc.view.RenderInfo([]string{"You are dead"})
		uc.SaveStats()
		return GameOver
	}

	player, _ := uc.dungeon.Player.(*unit.Character)

	if uc.dungeon.IsExit() {

		if uc.dungeon.LevelNumber == 21 {
			uc.view.RenderCharacterWinWindow()
			uc.SaveStats()
			return Win
		}
		uc.dungeon = storage.GenerateDungeonFromConfig(uc.dungeon.LevelNumber+1, uc.cfg, player)
		uc.view.GameWindow.Erase()
		if err := uc.SaveGame(); err != nil {
			panic(fmt.Sprintf("save failed: %v", err))
		}
		return NextLevel
	}
	uc.view.Render(uc.dungeon)
	return ContinueGame
}

// RenderInitial forces an immediate render of the game state
func (uc *PlayerActionUseCase) RenderInitial() {
	uc.view.Render(uc.dungeon)
}

// RenderMainMenu get main window after death or win
func (uc *PlayerActionUseCase) RenderMainMenu() {
	uc.view.RenderMainWindow()
}

// SaveGame save current game to JSON file
func (uc *PlayerActionUseCase) SaveGame() error {
	stor := storage.NewJSONDungeonStorage()
	storDTO := storage.DungeonToDTO(uc.dungeon)
	err := stor.SaveGameState(storDTO)
	if err != nil {
		return fmt.Errorf("save error: %w", err)
	}
	player, ok := uc.dungeon.Player.(*unit.Character)
	if ok {
		err = stor.SaveLeaderboard(common.Stats(player.Stats))
		if err != nil {
			return fmt.Errorf("save error: %w", err)
		}
	}
	uc.view.RenderInfo([]string{"Game saved successfully!"})
	return nil
}

// SaveStats - push player's stats to leaderboard
func (uc *PlayerActionUseCase) SaveStats() error {
	stor := storage.NewJSONDungeonStorage()
	player, ok := uc.dungeon.Player.(*unit.Character)
	if err := stor.DeleteSave(); err != nil {
		return fmt.Errorf("save error: %w", err)
	}
	if ok {
		err := stor.SaveLeaderboard(common.Stats(player.Stats))
		if err != nil {
			return fmt.Errorf("save error: %w", err)
		}
	}
	uc.view.RenderInfo([]string{"Stats saved successfully!"})
	return nil
}

// LoadGame load last game from JSON
func (uc *PlayerActionUseCase) LoadGame() (*dungeon.Dungeon, error) {
	stor := storage.NewJSONDungeonStorage()
	loadedDTO, err := stor.LoadGameState()
	if err != nil {
		return nil, fmt.Errorf("load error: %w", err)
	}
	loadedDungeon := storage.DTOToDungeon(*loadedDTO)
	player, ok := loadedDungeon.Player.(*unit.Character)
	if !ok {
		return nil, fmt.Errorf("invalid player type in saved game")
	}
	cfg, err := storage.LoadDungeonConfig("configs/dungeon_config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	uc.dungeon = loadedDungeon
	uc.character = player
	uc.cfg = cfg
	if uc.character.CurrentWeapon != nil {
		uc.character.Strength -= uc.character.CurrentWeapon.Strength
		uc.character.CurrentWeapon = nil
	}
	return &loadedDungeon, nil
}

// NewGame create new game
func (uc *PlayerActionUseCase) NewGame() {
	cfgGame, err := storage.LoadDungeonConfig("configs/dungeon_config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	uc.dungeon = storage.GenerateDungeonFromConfig(1, cfgGame, nil)
	character, ok := uc.dungeon.Player.(*unit.Character)
	if !ok {
		panic("player is not of type *unit.Character")
	}
	uc.character = character
	uc.cfg = cfgGame
	uc.view.GameWindow.Clear()
}

// RenderLeaderBord displays the game's leaderboard statistics in MainWindow.
// It retrieves the leaderboard data, renders the statistics window, waits for user input,
// and then returns to the main window.
func (uc *PlayerActionUseCase) RenderLeaderBord() {
	stats := storage.GetLeaderboardSlice()
	uc.view.RenderStatisticWindow(stats)
	uc.view.MainWindow.GetChar()
	uc.view.RenderMainWindow()
}
