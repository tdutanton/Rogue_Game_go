package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
)

const (
	// GameStateFile defines the path where game state is saved
	GameStateFile = "internal/adapters/secondary/storage/dungeon_save.json"

	// LeaderboardFile defines the path where leaderboard data is stored
	LeaderboardFile = "internal/adapters/secondary/storage/leaderboard.json"

	// MaxLeaderboardLen specifies the maximum number of entries kept in the leaderboard
	MaxLeaderboardLen = 30
)

// DungeonStorage defines the interface for persistent game data operations.
// It handles saving/loading game state and managing leaderboard entries.
type DungeonStorage interface {
	// SaveGameState persists the current dungeon state to storage
	SaveGameState(data DungeonData) error

	// LoadGameState retrieves the saved game state from storage
	LoadGameState() (*DungeonData, error)

	// SaveLeaderboard adds a new entry to the leaderboard and maintains sorting
	SaveLeaderboard(entry common.Stats) error

	// GetLeaderboard retrieves the current leaderboard sorted by highest score
	GetLeaderboard() ([]common.Stats, error)
}

// JSONDungeonStorage implements DungeonStorage using JSON files for persistence
type JSONDungeonStorage struct{}

// NewJSONDungeonStorage creates a new instance of JSON-based dungeon storage
func NewJSONDungeonStorage() *JSONDungeonStorage {
	return &JSONDungeonStorage{}
}

// SaveGameState serializes the dungeon data to JSON and writes it to GameStateFile.
// Returns an error if marshaling or file operations fail.
func (s *JSONDungeonStorage) SaveGameState(data DungeonData) error {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GameStateFile, dataBytes, 0o644)
}

// LoadGameState reads the game state from GameStateFile and deserializes it.
// Returns a pointer to DungeonData or error if file doesn't exist or is invalid.
func (s *JSONDungeonStorage) LoadGameState() (*DungeonData, error) {
	data, err := os.ReadFile(GameStateFile)
	if err != nil {
		return nil, err
	}

	var dungeonData DungeonData
	if err := json.Unmarshal(data, &dungeonData); err != nil {
		return nil, err
	}

	return &dungeonData, nil
}

// SaveLeaderboard adds a new entry to the leaderboard while maintaining:
// - Sorting by treasures received (descending)
// - Maximum length of MaxLeaderboardLen
// Returns error if file operations fail.
func (s *JSONDungeonStorage) SaveLeaderboard(entry common.Stats) error {
	entries, err := s.GetLeaderboard()
	if err != nil {
		entries = []common.Stats{}
	}

	entries = append(entries, entry)

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].TreasuresReceived > entries[j].TreasuresReceived
	})

	if len(entries) > MaxLeaderboardLen {
		entries = entries[:MaxLeaderboardLen]
	}

	dataBytes, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(LeaderboardFile, dataBytes, 0o644)
}

// GetLeaderboard retrieves the current leaderboard entries from storage.
// Returns empty slice if file doesn't exist yet.
// Returns error if file exists but contains invalid data.
func (s *JSONDungeonStorage) GetLeaderboard() ([]common.Stats, error) {
	data, err := os.ReadFile(LeaderboardFile)
	if err != nil {
		return nil, err
	}

	var entries []common.Stats
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}

// DeleteSave deletes current saving in *.json file
func (s *JSONDungeonStorage) DeleteSave() error {
	if err := os.Remove(GameStateFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete save file: %w", err)
	}
	return nil
}

// GetLeaderboardSlice returns slice of leaderboard stats from storage, sorted descending.
func GetLeaderboardSlice() []common.Stats {
	stats, err := NewJSONDungeonStorage().GetLeaderboard()

	if err == nil {
		sort.Slice(stats, func(i, j int) bool {
			return stats[i].TreasuresReceived > stats[j].TreasuresReceived
		})
		return stats
	}
	return nil
}

// RemoveLastRecord removes the last entry from the leaderboard file.
// It reads the current leaderboard, removes the last record, and writes the updated list back to the file.
// Returns an error if reading, unmarshaling, or writing the file fails.
func RemoveLastRecord() error {
	data, err := os.ReadFile(LeaderboardFile)
	if err != nil {
		return err
	}

	var stats []common.Stats
	if err := json.Unmarshal(data, &stats); err != nil {
		return err
	}

	if len(stats) > 0 {
		stats = stats[:len(stats)-1]
	}

	output, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(LeaderboardFile, output, 0o644)
	if err != nil {
		return err
	}
	return nil
}
