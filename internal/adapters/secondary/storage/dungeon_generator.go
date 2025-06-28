package storage

import (
	"math/rand"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// GenerateDungeonFromConfig creates a complete dungeon level based on configuration.
// It handles player placement, item generation, enemy spawning, and room assignment.
// Parameters:
//   - level: The dungeon level to generate
//   - cfg: Configuration containing game parameters and probabilities
//   - player: Optional existing player character (new one created if nil)
//
// Returns a fully populated Dungeon structure ready for gameplay.
func GenerateDungeonFromConfig(level int, cfg *Config, player *unit.Character) dungeon.Dungeon {
	lvlCfg := pickLevelConfig(cfg.Levels, level)

	d := dungeon.GenerateDungeon()
	d.LevelNumber = level

	if player == nil {
		player = createPlayer(cfg.CharacterStartParams)
	}
	player.Stats.LevelAchieved = level
	d.Player = player

	d.Items = generateItems(cfg, lvlCfg.ItemsCount)
	d.Enemies = generateEnemies(cfg, lvlCfg.EnemyChances, lvlCfg.EnemyCount, lvlCfg.Treasure)
	var startRoom, endRoom *dungeon.Room
	for i := range d.Rooms {
		if d.Rooms[i].Type == dungeon.RoomStart {
			startRoom = &d.Rooms[i]
		} else if d.Rooms[i].Type == dungeon.RoomEnd {
			endRoom = &d.Rooms[i]
		}
	}

	if startRoom != nil {
		startCoords := getRandomFloorCoord(*startRoom)
		d.Player.SetCoords(startCoords)
		startRoom.Visible = dungeon.FogClean
	} else {
		panic("no start room found")
	}

	occupiedCoords := make(map[common.Coords]bool)
	for i := range d.Items {
		var coord common.Coords
		var attempts int
		maxAttempts := 100

		for attempts = 0; attempts < maxAttempts; attempts++ {
			room := getRandomNonSpecialRoom(d.Rooms[:], startRoom, endRoom)
			coord = getRandomFloorCoord(room)
			if (!occupiedCoords[coord]) && (endRoom == nil || coord != d.Exit) {
				break
			}
		}

		d.Items[i].SetCoords(coord)
		occupiedCoords[coord] = true
	}

	for _, e := range d.Enemies {
		var coord common.Coords
		var attempts int
		maxAttempts := 100
		for attempts = 0; attempts < maxAttempts; attempts++ {
			room := getRandomNonSpecialRoom(d.Rooms[:], startRoom, endRoom)
			coord = getRandomFloorCoord(room)
			if endRoom == nil || coord != d.Exit {
				break
			}
		}
		if attempts >= maxAttempts {
			room := getRandomNonSpecialRoom(d.Rooms[:], startRoom, endRoom)
			coord = getRandomFloorCoord(room)
		}
		e.SetCoords(coord)
	}
	return d
}

// createPlayer initializes a new player character with starting parameters.
// It sets base attributes and initializes default statistics.
func createPlayer(params Character) *unit.Character {
	p := &unit.Character{}
	p.MaxHealth = params.MaxHealth
	p.Health = params.Health
	p.Strength = params.Strength
	p.Agility = params.Agility
	p.Stats.LevelAchieved = 1
	return p
}

// generateEnemies creates a set of enemies for the dungeon based on configuration.
// Parameters:
//   - cfg: Game configuration containing enemy definitions
//   - chances: Probability weights for different enemy types
//   - countRange: Minimum and maximum number of enemies to generate
//   - treasureRange: Range of treasure values for enemies
//
// Returns a slice of enemies implementing the Coordinator interface.
func pickLevelConfig(levels []Level, current int) Level {
	for _, lvl := range levels {
		if current >= lvl.Range[0] && current <= lvl.Range[1] {
			return lvl
		}
	}
	return levels[len(levels)-1]
}

// createEnemy constructs a specific enemy instance with randomized attributes.
// Parameters:
//   - cfg: Game configuration containing enemy segments
//   - etype: Enemy type identifier
//   - treasureRange: Range for random treasure value
//
// Returns an enemy implementing the Coordinator interface.
func generateEnemies(cfg *Config, chances map[string]int, countRange [2]int, treasureRange [2]int) []dungeon.Coordinator {
	count := common.RandomInRange(countRange[0], countRange[1])
	enemies := make([]dungeon.Coordinator, 0, count)

	weights := make([]string, 0)
	for enemy, chance := range chances {
		for i := 0; i < chance; i++ {
			weights = append(weights, enemy)
		}
	}
	for i := 0; i < count; i++ {
		etype := weights[rand.Intn(len(weights))]
		enemies = append(enemies, createEnemy(cfg, etype, treasureRange))
	}
	return enemies
}

// createEnemy constructs a specific enemy instance with randomized attributes.
// Parameters:
//   - cfg: Game configuration containing enemy segments
//   - etype: Enemy type identifier
//   - treasureRange: Range for random treasure value
//
// Returns an enemy implementing the Coordinator interface.
func createEnemy(cfg *Config, etype string, treasureRange [2]int) dungeon.Coordinator {
	einfo := cfg.Enemies[etype]

	agilityRange := cfg.EnemyAgility[einfo.EnemyAgility]
	strengthRange := cfg.EnemyStrength[einfo.EnemyStrength]
	animosityRange := cfg.EnemyAnimosity[einfo.EnemyAnimosity]
	healthRange := cfg.EnemyHealth[einfo.EnemyHealth]

	e := EnemyNum{
		"zombie":       0,
		"vampire":      1,
		"ghost":        2,
		"ogr":          3,
		"snake_wizard": 4,
	}
	enemyIndex, ok := e[etype]
	if !ok {
		panic("unknown enemy type: " + etype)
	}
	enemyType := unit.EnemyType(enemyIndex)

	enemy := unit.NewEnemyWithoutCoords(enemyType)

	enemy.Agility = common.RandomInRange(agilityRange[0], agilityRange[1])
	enemy.Strength = common.RandomInRange(strengthRange[0], strengthRange[1])
	enemy.Animosity = common.RandomInRange(animosityRange[0], animosityRange[1])
	enemy.Health = common.RandomInRange(healthRange[0], healthRange[1])
	enemy.Treasure = common.RandomInRange(treasureRange[0], treasureRange[1])

	return enemy
}

// generateItems creates random items to populate the dungeon.
// Parameters:
//   - cfg: Game configuration containing item definitions
//   - countRange: Minimum and maximum number of items to generate
//
// Returns a slice of generated items.
func generateItems(cfg *Config, countRange [2]int) []item.Item {
	count := common.RandomInRange(countRange[0], countRange[1])
	items := make([]item.Item, 0, count)

	for i := 0; i < count; i++ {
		switch rand.Intn(4) {
		case 0:
			items = append(items, createElixir(cfg.Elixir))
		case 1:
			items = append(items, createScroll(cfg.Scroll))
		case 2:
			items = append(items, createFood(cfg.Food))
		case 3:
			items = append(items, createWeapon(cfg.Weapon))
		}
	}

	return items
}

// createElixir generates an elixir item with random properties.
// The elixir will have one randomly selected beneficial effect.
func createElixir(e ItemEffects) item.Item {
	el := &item.Elixir{}
	el.Duration = common.RandomInRange(e.Duration[0], e.Duration[1])
	el.IsActive = false
	el.Name = e.Name[rand.Intn(len(e.Name))]

	switch rand.Intn(3) {
	case 0:
		el.Agility = common.RandomInRange(e.Agility[0], e.Agility[1])
		el.Strength = 0
		el.MaxHealth = 0
	case 1:
		el.Agility = 0
		el.Strength = common.RandomInRange(e.Strength[0], e.Strength[1])
		el.MaxHealth = 0
	case 2:
		el.Agility = 0
		el.Strength = 0
		el.MaxHealth = common.RandomInRange(e.MaxHealth[0], e.MaxHealth[1])
	}

	return el
}

// createScroll generates a scroll item with random properties.
// The scroll will have one randomly selected magical effect.
func createScroll(s ItemEffects) item.Item {
	sc := &item.Scroll{}
	sc.Duration = common.RandomInRange(s.Duration[0], s.Duration[1])
	sc.IsActive = false
	sc.Name = s.Name[rand.Intn(len(s.Name))]

	switch rand.Intn(3) {
	case 0:
		sc.Agility = common.RandomInRange(s.Agility[0], s.Agility[1])
		sc.Strength = 0
		sc.MaxHealth = 0
	case 1:
		sc.Agility = 0
		sc.Strength = common.RandomInRange(s.Strength[0], s.Strength[1])
		sc.MaxHealth = 0
	case 2:
		sc.Agility = 0
		sc.Strength = 0
		sc.MaxHealth = common.RandomInRange(s.MaxHealth[0], s.MaxHealth[1])
	}

	return sc
}

// createFood generates a food item with random nutritional value.
func createFood(f FoodEffects) item.Item {
	food := &item.Food{}
	food.Value = common.RandomInRange(f.Health[0], f.Health[1])
	food.Name = f.Name[rand.Intn(len(f.Name))]
	return food
}

// createWeapon generates a weapon with random damage properties.
func createWeapon(w WeaponEffects) item.Item {
	weapon := &item.Weapon{
		Name:     w.Name[rand.Intn(len(w.Name))],
		Strength: common.RandomInRange(w.Strength[0], w.Strength[1]),
	}
	return weapon
}

// getRandomFloorCoord returns a random walkable position within a room.
// The position will be within the room's boundaries excluding walls.
func getRandomFloorCoord(room dungeon.Room) common.Coords {
	minX := room.X + 1
	maxX := room.X + room.Width - 2
	minY := room.Y + 1
	maxY := room.Y + room.Height - 2

	x := common.RandomInRange(minX, maxX)
	y := common.RandomInRange(minY, maxY)

	return common.Coords{X: x, Y: y}
}

// getRandomNonSpecialRoom selects a random room that isn't the start or end room.
// Falls back to any non-start room if no other options exist.
// Panics if no suitable rooms are available.
func getRandomNonSpecialRoom(rooms []dungeon.Room, startRoom, endRoom *dungeon.Room) dungeon.Room {
	var candidates []dungeon.Room
	for _, room := range rooms {
		if (startRoom == nil || room.Coords != startRoom.Coords) &&
			(endRoom == nil || room.Coords != endRoom.Coords) {
			candidates = append(candidates, room)
		}
	}
	if len(candidates) == 0 {
		for _, room := range rooms {
			if startRoom == nil || room.Coords != startRoom.Coords {
				candidates = append(candidates, room)
			}
		}
		if len(candidates) == 0 {
			panic("no suitable rooms available")
		}
	}
	return candidates[rand.Intn(len(candidates))]
}
