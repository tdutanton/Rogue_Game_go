package ports

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
)

// Renderer interface with Render method to render (draw) something on the screen
type Renderer interface {
	Render(dungeon dungeon.Dungeon)
}
