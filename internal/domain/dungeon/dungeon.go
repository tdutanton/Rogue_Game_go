package dungeon

import (
	"fmt"
	"math"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
	"github.com/tdutanton/Rogue_Game_go/internal/domain/item"
)

// Dungeon represents a collection of rooms and corridors, connected together, and exit from level coords
type Dungeon struct {
	Rooms       [common.MaxRoomCount]Room
	Passages    []Passage
	Exit        common.Coords
	Player      Coordinator
	LevelNumber int
	Items       []item.Item
	Enemies     []Coordinator
	EventData   []string
}

// Passage represents a path connecting rooms in a dungeon
type Passage struct {
	Path []Corridor
}

// Contains checks if the given coordinates are within this passage boundaries.
func (p *Passage) Contains(c common.Coords) bool {
	for i := range p.Path {
		if p.Path[i].Contains(c) {
			return true
		}
	}
	return false
}

// NewDungeon create new dungeon
func NewDungeon(player Coordinator) *Dungeon {
	dungeon := GenerateDungeon()
	dungeon.Player = player

	return &dungeon
}

// GenerateDungeon creates a new dungeon by generating a fixed number of rooms and their connections.
func GenerateDungeon() Dungeon {
	var dungeon Dungeon
	for i := 0; i < common.MaxRoomCount; i++ {
		dungeon.Rooms[i] = generateRoom(i)
	}
	generateRoomConnections(dungeon.Rooms[:], &dungeon)

	return dungeon
}

// TileFromEntities - get tile type under the entities
func (d *Dungeon) TileFromEntities(c common.Coords) (common.TileType, bool) {
	for _, item := range d.Items {
		if c == item.GetCoords() {
			return common.ItemTile, true
		}
	}
	for _, enemy := range d.Enemies {
		if c == enemy.GetCoords() {
			return common.EnemyTile, true
		}
	}
	return common.UnknownTile, false
}

// TileFromItems - get tile type under the items
func (d *Dungeon) TileFromItems(c common.Coords) (common.TileType, bool) {
	for _, item := range d.Items {
		if c == item.GetCoords() {
			return common.ItemTile, true
		}
	}
	return common.UnknownTile, false
}

// TileFromRooms - get tile type in the room
func (d *Dungeon) TileFromRooms(c common.Coords) (common.TileType, bool) {
	for _, room := range d.Rooms {
		if IsCoordInRoom(c, room) {
			return common.FloorTile, true
		}
		if IsCoordInWall(c, room) {
			for _, door := range room.Doors {
				if c == door.Coords {
					return common.DoorTile, true
				}
			}
			return common.WallTile, true
		}
	}
	return common.UnknownTile, false
}

// TileFromPassages - get tile type in the passage
func (d *Dungeon) TileFromPassages(c common.Coords) (common.TileType, bool) {
	for _, passage := range d.Passages {
		for _, corridor := range passage.Path {
			if IsCoordInCorridor(c, corridor) {
				return common.CorridorTile, true
			}
		}
	}
	return common.UnknownTile, false
}

// Tile returns the type of tile at the specified coordinates in the dungeon.
// Returns UnknownTile and an error if coordinates are out of bounds.
func (d *Dungeon) Tile(c common.Coords) (common.TileType, error) {
	if c == d.Exit {
		return common.FinishTile, nil
	}
	if c == d.Player.GetCoords() {
		return common.PlayerTile, nil
	}
	if tile, ok := d.TileFromEntities(c); ok {
		return tile, nil
	}
	if tile, ok := d.TileFromRooms(c); ok {
		return tile, nil
	}
	if tile, ok := d.TileFromPassages(c); ok {
		return tile, nil
	}

	return common.UnknownTile, fmt.Errorf("no tile found at coordinates (%d, %d)", c.X, c.Y)
}

// TileUnderEnemy - gettile type under the Character
func (d *Dungeon) TileUnderEnemy(c common.Coords) (common.TileType, error) {
	if c == d.Exit {
		return common.FinishTile, nil
	}
	if c == d.Player.GetCoords() {
		return common.PlayerTile, nil
	}
	if tile, ok := d.TileFromItems(c); ok {
		return tile, nil
	}
	if tile, ok := d.TileFromRooms(c); ok {
		return tile, nil
	}
	if tile, ok := d.TileFromPassages(c); ok {
		return tile, nil
	}

	return common.UnknownTile, fmt.Errorf("no tile found at coordinates (%d, %d)", c.X, c.Y)
}

// PlayerCoords returns the current coordinates of the player in the dungeon.
func (d *Dungeon) PlayerCoords() common.Coords {
	return d.Player.GetCoords()
}

// FreeCoords - get slice of empty coords in dungeon
func (d *Dungeon) FreeCoords() []common.Coords {
	var coords []common.Coords
	for _, room := range d.Rooms {
		for y := 1; y < room.Size.Height-1; y++ {
			for x := 1; x < room.Size.Width-1; x++ {
				c := common.Coords{X: room.Coords.X + x, Y: room.Coords.Y + y}
				coords = append(coords, c)
			}
		}
	}
	return coords
}

// CurrentRoom - get pointer to Room in which Character is in now
func (d *Dungeon) CurrentRoom() *Room {
	for _, room := range d.Rooms {
		if room.Contains(d.PlayerCoords()) {
			return &room
		}
	}

	return nil
}

// CurrentRoomWithWalls - get pointer to Room in which Character is in now include walls
func (d *Dungeon) CurrentRoomWithWalls() *Room {
	for _, room := range d.Rooms {
		if room.ContainsIncludeWalls(d.PlayerCoords()) {
			return &room
		}
	}

	return nil
}

// FindDropPosition - get the nearest empty point to frop somethng on it
func (d *Dungeon) FindDropPosition() *common.Coords {
	center := d.PlayerCoords()
	var bestPos *common.Coords
	var minDist float64 = math.MaxFloat64
	for dx := -2; dx <= 2; dx++ {
		for dy := -2; dy <= 2; dy++ {
			checkPos := common.Coords{X: center.X + dx, Y: center.Y + dy}
			tile, _ := d.Tile(checkPos)
			if tile == common.FloorTile {
				dist := math.Abs(float64(dx)) + math.Abs(float64(dy))
				if dist < minDist {
					minDist = dist
					bestPos = &checkPos
				}
			}
		}
	}
	return bestPos
}

// AddItemToNearestPosition - place Item on the nearest empty tile
func (d *Dungeon) AddItemToNearestPosition(weapon item.Weapon) bool {
	dropPos := d.FindDropPosition()
	if dropPos == nil {
		return false
	}
	weapon.Coords = *dropPos
	d.Items = append(d.Items, &weapon)
	return true
}

// UpdateVisibleRoomsStatus check is Player in the Room and switch visible status
func (d *Dungeon) UpdateVisibleRoomsStatus() {
	pos := d.PlayerCoords()
	for i := range d.Rooms {
		d.Rooms[i].Visible = FogCovered
		if d.Rooms[i].Contains(pos) {
			d.Rooms[i].Visible = FogClean
		}
		for _, door := range d.Rooms[i].Doors {
			if door.Coords == pos {
				x1, _ := d.Tile(common.Coords{X: pos.X - 1, Y: pos.Y})
				x2, _ := d.Tile(common.Coords{X: pos.X + 1, Y: pos.Y})
				if x1 == common.WallTile && x2 == common.WallTile {
					d.Rooms[i].Visible = VerticalFog
				} else {
					d.Rooms[i].Visible = HorizontalFog
				}
			}
		}
	}
}

// IsExit - check is Character on the Exit tile
func (d *Dungeon) IsExit() bool {
	return d.PlayerCoords() == d.Exit
}

// CurrentPassage - return current passage where's the Character is on
func (d *Dungeon) CurrentPassage() *Passage {
	coords := d.PlayerCoords()

	for i, passage := range d.Passages {
		for _, corridor := range passage.Path {
			if corridor.Contains(coords) {
				return &d.Passages[i]
			}
		}
	}

	return nil
}

// UpdateVisibleArea marks rooms and passages as visited when the player is located in it.
func (d *Dungeon) UpdateVisibleArea() {
	player := d.PlayerCoords()

	for i := range d.Rooms {
		room := &d.Rooms[i]
		if IsCoordInRoom(player, *room) || IsCoordInWall(player, *room) {
			room.Visited = true
		}
	}

	for i := range d.Passages {
		passage := &d.Passages[i]
		for j := range passage.Path {
			corridor := &passage.Path[j]
			if IsCoordInCorridor(player, *corridor) {
				corridor.Visited = true
			}
		}
	}
}

// Update - combine two methods to correct updating the dungeon parameters
func (d *Dungeon) Update() {
	d.UpdateVisibleArea()
	d.UpdateVisibleRoomsStatus()
}

// AddEventData appends a new event data string to the dungeon's event slice for rendering on game screen.
func (d *Dungeon) AddEventData(data string) {
	d.EventData = append(d.EventData, data)
}

// ClearEventData removes all event data from the dungeon, resetting the event slice to empty.
func (d *Dungeon) ClearEventData() {
	d.EventData = nil
}

// ReplaceEventData replaces an old event data by new string.
func (d *Dungeon) ReplaceEventData(data string) {
	d.ClearEventData()
	d.AddEventData(data)
}
