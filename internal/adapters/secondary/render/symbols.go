package render

import gc "github.com/gbin/goncurses"

// Character represents the player's symbol in the game world.
const Character = '@'

// Item symbols used for rendering inventory items in the game world.
const (
	Food   = 'f' // Symbol for food items
	Elixir = 'e' // Symbol for elixir/potion items
	Weapon = 'w' // Symbol for weapon items
	Scroll = 's' // Symbol for scroll items
)

// Enemy symbols used for rendering different enemy types.
const (
	Zombie      = 'Z' // Symbol for zombie enemies
	Vampire     = 'V' // Symbol for vampire enemies
	Ghost       = 'G' // Symbol for ghost enemies
	Ogr         = 'O' // Symbol for ogre enemies
	SnakeWizard = 'S' // Symbol for snake wizard enemies
)

// Dungeon structure symbols using ncurses ACS characters.
// These constants define the visual representation of dungeon elements.
const (
	WallHorizontal   = gc.ACS_HLINE    // Horizontal wall segment ─
	WallVertical     = gc.ACS_VLINE    // Vertical wall segment │
	UpperLeftCorner  = gc.ACS_ULCORNER // Upper left corner ┌
	UpperRightCorner = gc.ACS_URCORNER // Upper right corner ┐
	LowerLeftCorner  = gc.ACS_LLCORNER // Lower left corner └
	LowerRightCorner = gc.ACS_LRCORNER // Lower right corner ┘
	Passage          = gc.ACS_CKBOARD  // Passage or tunnel block ▒
	Door             = gc.ACS_PLUS     // Door symbol +
	EmptyFloor       = ' '             // Empty floor space
	Exit             = 'E'             // Dungeon exit symbol
	Fog              = '.'             // Unexplored area/for symbol
)
