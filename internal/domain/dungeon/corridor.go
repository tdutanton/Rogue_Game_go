package dungeon

import (
	"math/rand"

	"github.com/tdutanton/Rogue_Game_go/internal/domain/common"
)

// Door represents an exit from Room to passage located at a specific coordinate position
type Door struct {
	common.Coords
}

// Corridor represents a part of Passage
type Corridor struct {
	Begin   common.Coords // Coordinates of corridor begin
	End     common.Coords // Coordinates of corridor end
	Visited bool          // Indicates whether the corridor has been visited by a player
}

// getNeighbors returns a slice of room indices that are adjacent to the room at the given index.
func getNeighbors(index int) []int {
	var neighbors []int
	rows, cols := common.RoomsInRow, common.RoomsInColumn
	row, col := index/cols, index%cols

	if row > 0 {
		neighbors = append(neighbors, (row-1)*cols+col)
	}
	if row < rows-1 {
		neighbors = append(neighbors, (row+1)*cols+col)
	}
	if col > 0 {
		neighbors = append(neighbors, row*cols+(col-1))
	}
	if col < cols-1 {
		neighbors = append(neighbors, row*cols+(col+1))
	}
	return neighbors
}

// shuffle randomly reorders the elements of the input slice using a permutation.
// It returns a new slice with the same elements in a random order.
func shuffle(nums []int) []int {
	r := rand.Perm(len(nums))
	shuffled := make([]int, len(nums))
	for i, v := range r {
		shuffled[i] = nums[v]
	}
	return shuffled
}

// connectRoomsDFS performs a depth-first search to connect rooms in the dungeon.
// It takes the current room index, a slice of rooms, and a visited tracker as input.
// The function recursively explores unvisited neighboring rooms, generating passages
// between connected rooms and marking rooms as plain type during traversal.
// It returns a slice of passages created during the room connection process.
func connectRoomsDFS(index int, rooms []Room, visited []bool) []Passage {
	visited[index] = true
	var passages []Passage

	neighbors := shuffle(getNeighbors(index))
	for _, n := range neighbors {
		if !visited[n] {

			passage := generatePassage(rooms, index, n)
			passages = append(passages, passage)

			rooms[n].Type = RoomPlain

			childPassages := connectRoomsDFS(n, rooms, visited)
			passages = append(passages, childPassages...)
		}
	}
	return passages
}

// generateRoomConnections creates connections between rooms in a dungeon by performing a depth-first search traversal.
// It randomly selects a start and end room, generates passages between rooms, and marks the start and end rooms.
func generateRoomConnections(rooms []Room, dg *Dungeon) {
	visited := make([]bool, len(rooms))

	start := rand.Intn(len(rooms))
	end := rand.Intn(len(rooms) - 1)
	if end >= start {
		end++
	}

	dg.Passages = connectRoomsDFS(start, rooms, visited)
	rooms[start].Type = RoomStart
	rooms[start].Visited = true
	rooms[end].Type = RoomEnd
	dg.Exit = generateExitPoint(&rooms[end])
}

// generatePassage creates a passage between two rooms in a dungeon grid.
func generatePassage(rooms []Room, src int, dest int) Passage {
	if dest < src {
		src, dest = dest, src
	}
	var passage Passage

	cols := common.RoomsInColumn
	rowSrc, colSrc := src/cols, src%cols
	rowDest, colDest := dest/cols, dest%cols

	if rowSrc == rowDest {
		generateHorizontalCorridor(&rooms[src], &rooms[dest], &passage)
	}

	if colSrc == colDest {
		generateVerticalCorridor(&rooms[src], &rooms[dest], &passage)
	}

	return passage
}

// generateHorizontalCorridor creates a horizontal corridor between two rooms in a dungeon.
// It randomly selects door positions on the source and destination rooms' walls,
// and generates a path with optional turns to connect the rooms horizontally.
func generateHorizontalCorridor(srcRoom *Room, destRoom *Room, pass *Passage) {
	srcWall, dstWall := srcRoom.X+srcRoom.Width-1, destRoom.X

	srcMaxBorder, srcMinBorder := srcRoom.Y+srcRoom.Height-2, srcRoom.Y+1
	srcY := rand.Intn(srcMaxBorder-srcMinBorder) + srcMinBorder

	dstMaxBorder, dstMinBorder := destRoom.Y+destRoom.Height-2, destRoom.Y+1
	dstY := rand.Intn(dstMaxBorder-dstMinBorder) + dstMinBorder

	srcRoom.Doors = append(srcRoom.Doors, Door{common.Coords{X: srcWall, Y: srcY}})
	destRoom.Doors = append(destRoom.Doors, Door{common.Coords{X: dstWall, Y: dstY}})

	if srcY == dstY {
		addCorridor(pass, srcWall+1, srcY, dstWall-1, dstY)
	} else {
		turnX := getRandomTurnCoord(srcWall, dstWall)

		addCorridor(pass, srcWall+1, srcY, turnX, srcY)
		addCorridor(pass, turnX, srcY, turnX, dstY)
		addCorridor(pass, turnX, dstY, dstWall-1, dstY)
	}
}

// generateVerticalCorridor creates a vertical corridor between two rooms in a dungeon.
// It randomly selects door positions on the source and destination rooms' walls,
// and generates a path with optional turns to connect the rooms vertically.
func generateVerticalCorridor(srcRoom *Room, destRoom *Room, pass *Passage) {
	srcWall, dstWall := srcRoom.Y+srcRoom.Height-1, destRoom.Y

	srcMaxBorder, srcMinBorder := srcRoom.X+srcRoom.Width-2, srcRoom.X+1
	srcX := rand.Intn(srcMaxBorder-srcMinBorder) + srcMinBorder

	dstMaxBorder, dstMinBorder := destRoom.X+destRoom.Width-2, destRoom.X+1
	dstX := rand.Intn(dstMaxBorder-dstMinBorder) + dstMinBorder

	srcRoom.Doors = append(srcRoom.Doors, Door{common.Coords{X: srcX, Y: srcWall}})
	destRoom.Doors = append(destRoom.Doors, Door{common.Coords{X: dstX, Y: dstWall}})

	if srcX == dstX {
		addCorridor(pass, srcX, srcWall+1, dstX, dstWall-1)
	} else {
		turnY := getRandomTurnCoord(srcWall, dstWall)

		addCorridor(pass, srcX, srcWall+1, srcX, turnY)
		addCorridor(pass, srcX, turnY, dstX, turnY)
		addCorridor(pass, dstX, turnY, dstX, dstWall-1)
	}
}

// getRandomTurnCoord calculates a random intermediate coordinate between two wall positions,
// ensuring the turn point is not too close to the source or destination walls.
func getRandomTurnCoord(srcWall, dstWall int) int {
	randCoord := rand.Intn(dstWall-srcWall-common.MinRoomDistance) + srcWall + 1
	if randCoord == srcWall+1 {
		randCoord++
	}
	if randCoord == dstWall-1 {
		randCoord--
	}
	return randCoord
}

// addCorridor adds a new corridor segment to the passage's path, representing a connection
// between two coordinates.
func addCorridor(pass *Passage, x1, y1, x2, y2 int) {
	pass.Path = append(pass.Path, Corridor{
		Begin:   common.Coords{X: x1, Y: y1},
		End:     common.Coords{X: x2, Y: y2},
		Visited: false,
	})
}

// Contains - check if coords is in the Corridor
func (c Corridor) Contains(coords common.Coords) bool {
	beginX := c.Begin.X
	beginY := c.Begin.Y
	endX := c.End.X
	endY := c.End.Y

	if beginX > endX {
		beginX, endX = endX, beginX
	}
	if beginY > endY {
		beginY, endY = endY, beginY
	}

	return beginX <= coords.X && beginY <= coords.Y &&
		endX >= coords.X && endY >= coords.Y
}
