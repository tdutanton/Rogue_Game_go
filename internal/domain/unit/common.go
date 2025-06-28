package unit

// Direction represents cardinal movement directions.
type Direction int

const (
	// Down Move downward (increasing Y)
	Down Direction = iota

	// Left Move left (decreasing X)
	Left

	// Up Move upward (decreasing Y)
	Up

	// Right Move right (increasing X)
	Right
)

// AngleDirection represents diagonal movement directions.
type AngleDirection int

const (
	// DownLeft Diagonal: down and left
	DownLeft AngleDirection = iota

	// UpLeft Diagonal: up and left
	UpLeft

	// UpRight Diagonal: up and right
	UpRight

	// DownRight Diagonal: down and right
	DownRight
)
