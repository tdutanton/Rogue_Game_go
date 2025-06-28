// Package storage provides the persistence layer for the Rogue game engine.
//
// It handles everything from dungeon generation to game state persistence,
// acting as the bridge between volatile game state and permanent storage.
//
// The package offers:
//   - Dungeon configuration loading (YAML)
//   - Game state management (JSON persistence)
//   - Leaderboard functionality
//   - DTO conversions between domains
//   - Procedural dungeon generation logic
//
// Key components:
//   - Config - Central game configuration structure
//   - DungeonData - Complete game state representation
//   - DungeonStorage - Interface for persistence operations
//   - JSONDungeonStorage - Concrete JSON-based implementation
//
// Usage example:
//
//	  Load game configuration:
//	  	cfg, err := storage.LoadDungeonConfig("config.yaml")
//
//	  Generate new dungeon:
//	  	dungeon := storage.GenerateDungeonFromConfig(1, cfg, nil)
//	  	(1 - is level number - you can put enything else)
//
//	  Initialize storage:
//	  	store := storage.NewJSONDungeonStorage()
//
//	  Save Game:
//			store.SaveGameState(dungeon)
//
//	  Load Game (last try):
//	  	loadedDungeon, err := stor.LoadGameState()
//
//	  Update Leader Board:
//	  		err = stor.SaveLeaderboard(common.Stats(storDTO.Player.Stats))
//				if err != nil {
//					log.Fatalf("Leaderboard error: %v", err)
//				}
//
//	  Get Leader Board:
//	  		leaderboard, err := stor.GetLeaderboard()
//				if err != nil {
//					log.Fatalf("Leaderboard error: %v", err)
//				}
//
//	  Print actual Leader Board (if you want to test the Package)
//				fmt.Println("\nLeaderboard:")
//				for i, entry := range leaderboard {
//					fmt.Printf("%d. Level: %d | Treasures: %d | Killed: %d\n",
//					i+1, entry.LevelAchieved, entry.TreasuresReceived, entry.EnemiesDefeated)
//				}
//
// The package follows clean architecture principles, keeping domain logic
// separate from storage concerns while providing convenient data conversion.
package storage
