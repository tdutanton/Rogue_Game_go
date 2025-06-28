package unit

import "github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"

// NonMoving implements EnemyMover for stationary enemies that don't move.
type NonMoving struct{}

// Move implements the EnemyMover interface for stationary enemies (no-op).
func (n NonMoving) Move(_ *Enemy, _ dungeon.Dungeon) {}
