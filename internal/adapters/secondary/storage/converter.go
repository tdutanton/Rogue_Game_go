package storage

import (
	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/dungeon"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/inventory"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/unit"
)

// CoordsToDTO converts domain Coordinates to DTO format.
func CoordsToDTO(c common.Coords) CoordsData {
	return CoordsData{c.X, c.Y}
}

// DTOtoCoords converts DTO Coordinates back to domain format.
func DTOtoCoords(cd CoordsData) common.Coords {
	return common.Coords{X: cd.X, Y: cd.Y}
}

// SizeToDTO converts domain Size to DTO format.
func SizeToDTO(s common.Size) SizeData {
	return SizeData{s.Width, s.Height}
}

// DTOToSize converts DTO Size back to domain format.
func DTOToSize(sd SizeData) common.Size {
	return common.Size{Width: sd.Width, Height: sd.Height}
}

// ItemToDTO converts a domain Item to its DTO representation.
// Handles all item types (Elixir, Scroll, Food, Weapon) with type-specific fields.
func ItemToDTO(i item.Item) ItemData {
	if i == nil {
		return ItemData{}
	}
	result := ItemData{
		Type:       int(i.Type()),
		CoordsData: CoordsData(i.GetCoords()),
		Name:       i.Info(),
	}
	switch v := i.(type) {
	case *item.Elixir:
		result.Agility = v.Agility
		result.Strength = v.Strength
		result.MaxHealth = v.MaxHealth
		result.Duration = v.Duration
		result.IsActive = v.IsActive
	case *item.Scroll:
		result.Agility = v.Agility
		result.Strength = v.Strength
		result.MaxHealth = v.MaxHealth
		result.Duration = v.Duration
		result.IsActive = v.IsActive
	case *item.Food:
		result.Value = v.Value
	case *item.Weapon:
		result.Strength = v.Strength
	}
	return result
}

// DTOToItem converts an Item DTO back to domain Item.
// Panics if the item type is unknown.
func DTOToItem(id ItemData) item.Item {
	switch id.Type {
	case int(item.FoodType):
		return &item.Food{
			Name:   id.Name,
			Value:  id.Value,
			Coords: common.Coords(id.CoordsData),
		}
	case int(item.ElixirType):
		return &item.Elixir{
			Name:      id.Name,
			Agility:   id.Agility,
			Strength:  id.Strength,
			MaxHealth: id.MaxHealth,
			Coords:    common.Coords(id.CoordsData),
			Duration:  id.Duration,
			IsActive:  id.IsActive,
		}
	case int(item.ScrollType):
		return &item.Scroll{
			Name:      id.Name,
			Agility:   id.Agility,
			Strength:  id.Strength,
			MaxHealth: id.MaxHealth,
			Coords:    common.Coords(id.CoordsData),
			Duration:  id.Duration,
			IsActive:  id.IsActive,
		}
	case int(item.WeaponType):
		return &item.Weapon{
			Name:     id.Name,
			Strength: id.Strength,
			Coords:   common.Coords(id.CoordsData),
		}
	default:
		panic("unknown item type")
	}
}

// StatsToDTO converts game statistics to DTO format.
func StatsToDTO(s common.Stats) StatsData {
	return StatsData{
		TreasuresReceived: s.TreasuresReceived,
		LevelAchieved:     s.LevelAchieved,
		EnemiesDefeated:   s.EnemiesDefeated,
		FoodEaten:         s.FoodEaten,
		ElixirsDrunk:      s.ElixirsDrunk,
		ScrollsRead:       s.ScrollsRead,
		HitsMade:          s.HitsMade,
		HitsMissed:        s.HitsMissed,
		CellsPassed:       s.CellsPassed,
	}
}

// DTOToStats converts statistics DTO back to domain format.
func DTOToStats(sd StatsData) common.Stats {
	return common.Stats{
		LevelAchieved:     sd.LevelAchieved,
		TreasuresReceived: sd.TreasuresReceived,
		EnemiesDefeated:   sd.EnemiesDefeated,
		FoodEaten:         sd.FoodEaten,
		ElixirsDrunk:      sd.ElixirsDrunk,
		ScrollsRead:       sd.ScrollsRead,
		HitsMade:          sd.HitsMade,
		HitsMissed:        sd.HitsMissed,
		CellsPassed:       sd.CellsPassed,
	}
}

// InventoryToDTO converts player inventory to DTO format.
// Includes all inventory categories: foods, elixirs, scrolls, and weapons.
func InventoryToDTO(i inventory.Inventory) InventoryData {
	result := InventoryData{
		Treasure: i.Treasure,
	}
	for _, food := range i.Foods {
		result.Foods = append(result.Foods, ItemData{
			Type:       int(food.Type()),
			CoordsData: CoordsData(food.GetCoords()),
			Name:       food.Info(),
			Value:      food.Value,
		})
	}
	for _, elixir := range i.Elixirs {
		result.Elixirs = append(result.Elixirs, ItemData{
			Type:       int(elixir.Type()),
			CoordsData: CoordsData(elixir.GetCoords()),
			Name:       elixir.Info(),
			Agility:    elixir.Agility,
			Strength:   elixir.Strength,
			MaxHealth:  elixir.MaxHealth,
			Duration:   elixir.Duration,
			IsActive:   elixir.IsActive,
		})
	}
	for _, scroll := range i.Scrolls {
		result.Scrolls = append(result.Scrolls, ItemData{
			Type:       int(scroll.Type()),
			CoordsData: CoordsData(scroll.GetCoords()),
			Name:       scroll.Info(),
			Agility:    scroll.Agility,
			Strength:   scroll.Strength,
			MaxHealth:  scroll.MaxHealth,
			Duration:   scroll.Duration,
			IsActive:   scroll.IsActive,
		})
	}
	for _, weapon := range i.Weapons {
		result.Weapons = append(result.Weapons, ItemData{
			Type:       int(weapon.Type()),
			CoordsData: CoordsData(weapon.GetCoords()),
			Name:       weapon.Info(),
			Strength:   weapon.Strength,
		})
	}
	return result
}

// DTOToInventory converts inventory DTO back to domain format.
func DTOToInventory(dto InventoryData) *inventory.Inventory {
	inv := &inventory.Inventory{
		Treasure: dto.Treasure,
	}
	for _, foodDTO := range dto.Foods {
		inv.Foods = append(inv.Foods, item.Food{
			Name:   foodDTO.Name,
			Value:  foodDTO.Value,
			Coords: common.Coords(foodDTO.CoordsData),
		})
	}
	for _, elixirDTO := range dto.Elixirs {
		inv.Elixirs = append(inv.Elixirs, item.Elixir{
			Name:      elixirDTO.Name,
			Agility:   elixirDTO.Agility,
			Strength:  elixirDTO.Strength,
			MaxHealth: elixirDTO.MaxHealth,
			Coords:    common.Coords(elixirDTO.CoordsData),
			Duration:  elixirDTO.Duration,
			IsActive:  elixirDTO.IsActive,
		})
	}
	for _, scrollDTO := range dto.Scrolls {
		inv.Scrolls = append(inv.Scrolls, item.Scroll{
			Name:      scrollDTO.Name,
			Agility:   scrollDTO.Agility,
			Strength:  scrollDTO.Strength,
			MaxHealth: scrollDTO.MaxHealth,
			Coords:    common.Coords(scrollDTO.CoordsData),
			Duration:  scrollDTO.Duration,
			IsActive:  scrollDTO.IsActive,
		})
	}
	for _, weaponDTO := range dto.Weapons {
		inv.Weapons = append(inv.Weapons, item.Weapon{
			Name:     weaponDTO.Name,
			Strength: weaponDTO.Strength,
			Coords:   common.Coords(weaponDTO.CoordsData),
		})
	}
	return inv
}

// UnitToDTO converts basic unit attributes to DTO format.
func UnitToDTO(u unit.Unit) UnitData {
	return UnitData{
		Health:     u.Health,
		Agility:    u.Agility,
		Strength:   u.Strength,
		CoordsData: CoordsData(u.Coords),
		InBattle:   u.InBattle,
	}
}

// DTOToUnit converts unit DTO back to domain format.
func DTOToUnit(ud UnitData) unit.Unit {
	return unit.Unit{
		Health:   ud.Health,
		Agility:  ud.Agility,
		Strength: ud.Strength,
		Coords:   common.Coords(ud.CoordsData),
		InBattle: ud.InBattle,
	}
}

// CharacterToDTO converts player character to DTO format.
// Includes unit attributes, inventory, and statistics.
func CharacterToDTO(c unit.Character) CharacterData {
	var weaponDTO ItemData
	if c.CurrentWeapon != nil {
		weaponDTO = ItemToDTO(c.CurrentWeapon)
	}
	return CharacterData{
		Unit:          UnitToDTO(c.Unit),
		MaxHealth:     c.MaxHealth,
		CurrentWeapon: weaponDTO,
		Inventory:     InventoryToDTO(c.Inventory),
		Stats:         StatsToDTO(c.Stats),
	}
}

// DTOToCharacter converts character DTO back to domain format.
func DTOToCharacter(cd CharacterData) unit.Character {
	weapon := item.Weapon{}
	if cd.CurrentWeapon.Type == int(item.WeaponType) {
		weapon = item.Weapon{
			Name:     cd.CurrentWeapon.Name,
			Strength: cd.CurrentWeapon.Strength,
			Coords:   common.Coords(cd.CurrentWeapon.CoordsData),
		}
	} else {
		weapon = item.Weapon{}
	}
	return unit.Character{
		Unit:          DTOToUnit(cd.Unit),
		MaxHealth:     cd.MaxHealth,
		CurrentWeapon: &weapon,
		Inventory:     *DTOToInventory(cd.Inventory),
		Stats:         DTOToStats(cd.Stats),
	}
}

// RoomToDTO converts dungeon room to DTO format.
// Includes room coordinates, size, doors, and type.
func RoomToDTO(r dungeon.Room) RoomData {
	d := make([]CoordsData, len(r.Doors))
	for i, v := range r.Doors {
		d[i] = CoordsToDTO(v.Coords)
	}
	return RoomData{
		CoordsData: CoordsToDTO(r.Coords),
		SizeData:   SizeToDTO(r.Size),
		Doors:      d,
		Type:       int(r.Type),
		Visited:    r.Visited,
		Visible:    int(r.Visible),
	}
}

// DTOToRoom converts room DTO back to domain format.
func DTOToRoom(rd RoomData) dungeon.Room {
	d := make([]dungeon.Door, len(rd.Doors))
	for i, v := range rd.Doors {
		d[i].Coords = DTOtoCoords(v)
	}
	room := dungeon.Room{
		Doors:   d,
		Type:    dungeon.RoomType(rd.Type),
		Visited: rd.Visited,
		Visible: dungeon.Visibility(rd.Visible),
	}
	room.Size = DTOToSize(rd.SizeData)
	room.Coords = DTOtoCoords(rd.CoordsData)
	return room
}

// CorridorToDTO converts dungeon corridor to DTO format.
func CorridorToDTO(c dungeon.Corridor) CorridorData {
	return CorridorData{
		Begin:   CoordsData(c.Begin),
		End:     CoordsData(c.End),
		Visited: c.Visited,
	}
}

// DTOToCorridor converts corridor DTO back to domain format.
func DTOToCorridor(cd CorridorData) dungeon.Corridor {
	return dungeon.Corridor{
		Begin:   common.Coords(cd.Begin),
		End:     common.Coords(cd.End),
		Visited: cd.Visited,
	}
}

// PassageToDTO converts dungeon passage (collection of corridors) to DTO format.
func PassageToDTO(p dungeon.Passage) PassageData {
	result := make([]CorridorData, len(p.Path))
	for i, v := range p.Path {
		result[i] = CorridorToDTO(v)
	}
	return PassageData{Path: result}
}

// DTOToPassage converts passage DTO back to domain format.
func DTOToPassage(pd PassageData) dungeon.Passage {
	result := make([]dungeon.Corridor, len(pd.Path))
	for i, v := range pd.Path {
		result[i] = DTOToCorridor(v)
	}
	return dungeon.Passage{Path: result}
}

// EnemyToDTO converts enemy unit to DTO format.
// Includes enemy-specific attributes like type, behavior flags, and treasure.
func EnemyToDTO(e unit.Enemy) EnemyData {
	return EnemyData{
		Unit:       UnitToDTO(e.Unit),
		EnemyType:  int(e.EnemyType),
		Animosity:  e.Animosity,
		Visibility: e.Visibility,
		IsPursuing: e.IsPursuing,
		Treasure:   e.Treasure,
	}
}

// DTOToEnemy converts enemy DTO back to domain format.
// Initializes appropriate movement behavior based on enemy type.
func DTOToEnemy(ed EnemyData) unit.Enemy {
	result := unit.Enemy{
		Unit:       DTOToUnit(ed.Unit),
		EnemyType:  unit.EnemyType(ed.EnemyType),
		Animosity:  ed.Animosity,
		Visibility: ed.Visibility,
		IsPursuing: ed.IsPursuing,
		Treasure:   ed.Treasure,
	}
	switch result.EnemyType {
	case unit.Ghost:
		result.DefaultMover = unit.GhostMoving{}
	case unit.SnakeWizard:
		result.DefaultMover = unit.SnakeWizardMoving{}
	case unit.Ogr:
		result.DefaultMover = unit.OgrMoving{}
	default:
		result.DefaultMover = unit.NonMoving{}
	}
	if result.IsPursuing {
		result.Mover = unit.PursuingMoving{}
	} else {
		result.Mover = result.DefaultMover
	}
	return result
}

// DungeonToDTO converts complete dungeon state to DTO format.
// Includes rooms, passages, items, enemies, and player character.
// Panics if player type assertion fails.
func DungeonToDTO(d dungeon.Dungeon) DungeonData {
	player, ok := d.Player.(*unit.Character)
	if !ok {
		panic("wrong player type")
	}
	result := DungeonData{
		LevelNumber: d.LevelNumber,
		Exit:        CoordsToDTO(d.Exit),
		Player:      CharacterToDTO(*player),
	}
	result.Rooms = [common.MaxRoomCount]RoomData{}
	for i, v := range d.Rooms {
		result.Rooms[i] = RoomToDTO(v)
	}
	p := make([]PassageData, len(d.Passages))
	for i, v := range d.Passages {
		p[i] = PassageToDTO(v)
	}
	it := make([]ItemData, len(d.Items))
	for i, v := range d.Items {
		it[i] = ItemToDTO(v)
	}
	e := make([]EnemyData, len(d.Enemies))
	for i, coord := range d.Enemies {
		enemy, ok := coord.(*unit.Enemy)
		if !ok {
			panic("wrong enemy type")
		}
		e[i] = EnemyToDTO(*enemy)
	}
	result.Passages = p
	result.Items = it
	result.Enemies = e
	return result
}

// DTOToDungeon converts dungeon DTO back to complete domain representation.
// Reconstructs all dungeon components including entities and their relationships.
func DTOToDungeon(dd DungeonData) dungeon.Dungeon {
	var rooms [common.MaxRoomCount]dungeon.Room
	for i, v := range dd.Rooms {
		rooms[i] = DTOToRoom(v)
	}
	passages := make([]dungeon.Passage, len(dd.Passages))
	for i, v := range dd.Passages {
		passages[i] = DTOToPassage(v)
	}
	items := make([]item.Item, len(dd.Items))
	for i, v := range dd.Items {
		items[i] = DTOToItem(v)
	}
	enemies := make([]dungeon.Coordinator, len(dd.Enemies))
	for i, v := range dd.Enemies {
		enemy := DTOToEnemy(v)
		enemies[i] = new(unit.Enemy)
		*enemies[i].(*unit.Enemy) = enemy
	}
	player := DTOToCharacter(dd.Player)
	return dungeon.Dungeon{
		LevelNumber: dd.LevelNumber,
		Rooms:       rooms,
		Passages:    passages,
		Exit:        DTOtoCoords(dd.Exit),
		Player:      &player,
		Items:       items,
		Enemies:     enemies,
	}
}
