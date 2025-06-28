package render

import (
	"fmt"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// RenderStatistic displays the player's current statistics in the statistics window.
// Formats and renders the player's vital information including:
//   - Current level
//   - Treasure count
//   - Health status (current/max)
//   - Agility score
//   - Strength score
//
// The information is displayed in red text on black background for high visibility.
// Window is cleared before rendering and refreshed afterward to ensure proper display.
//
// Parameters:
//   - player: The player character whose statistics to display (unit.Character)
//
// Display Format:
//
//	"Level:[level]  Treasures:[count]  Health:[current]([max])  Agility:[value]  Strength:[value]"
//
// Example Output:
//
//	"Level:5  Treasures:12  Health:32(45)  Agility:7  Strength:10"
func (v *View) RenderStatistic(player unit.Character) {
	v.StatisticWindow.Erase()
	v.StatisticWindow.Box(0, 0)
	stats := player.Stats
	statistic := fmt.Sprintf(" Level:%v \t Treasures:%v \t Health:%v(%v) \t Agility:%v \t Strength:%v",
		stats.LevelAchieved,
		stats.TreasuresReceived,
		player.Health, player.MaxHealth,
		player.Agility,
		player.Strength,
	)

	v.StatisticWindow.ColorOn(RedBlack)
	v.StatisticWindow.MovePrint(1, 7, statistic)
	v.StatisticWindow.ColorOff(RedBlack)
	v.StatisticWindow.Refresh()
}
