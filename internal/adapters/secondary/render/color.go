package render

// Color pair constants for ncurses color attributes.
// These constants represent predefined color combinations used throughout the game.
// The values correspond to ncurses color pair indices (1 through 6).
const (
	// WhiteBlack represents white text on black background
	WhiteBlack = 1 + iota

	// BlackWhite represents black text on white background
	BlackWhite

	// RedBlack represents red text on black background
	RedBlack

	// GreenBlack represents green text on black background
	GreenBlack

	// GreenWhite represents green text on white background
	GreenWhite

	// YellowBlack represents yellow text on black background
	YellowBlack

	// BlueBlack represents blue text on black background
	BlueBlack
)
