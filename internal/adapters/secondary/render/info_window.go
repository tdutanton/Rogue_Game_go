package render

// RenderInfo displays informational messages in the game's info window.
// It clears the window, renders each line of text with yellow-on-black coloring,
// and refreshes the display to make changes visible.
//
// Parameters:
//   - eventInfo: A slice of strings where each string represents a line of information
//     to be displayed in the info window. Line breaks should be handled
//     by separate strings in the slice rather than newline characters.
//
// Behavior:
//   - Clears the info window before rendering new content
//   - Applies yellow text on black background coloring
//   - Prints each string on a separate line (starting from top)
//   - Ensures proper ncurses refresh to display changes
//
// Example:
//
//	view.RenderInfo([]string{
//	    "Player hit Goblin for 5 damage",
//	    "Found Potion of Healing",
//	})
func (v *View) RenderInfo(eventInfo []string) {
	v.InfoWindow.Erase()
	v.InfoWindow.ColorOn(YellowBlack)

	for i := range eventInfo {
		v.InfoWindow.MovePrint(i, 0, eventInfo[i])
	}

	v.InfoWindow.ColorOff(YellowBlack)
	v.InfoWindow.Refresh()
}
