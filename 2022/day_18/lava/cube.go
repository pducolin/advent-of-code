package lava

type Direction int

const (
	TOP Direction = iota
	BOTTOM
	LEFT
	RIGHT
	FRONT
	BACK
)

type Cube struct {
	X int
	Y int
	Z int
}

var DIRECTIONS = []Direction{TOP, BOTTOM, LEFT, RIGHT, FRONT, BACK}

func (c *Cube) GetNeighbours() map[Direction]Cube {
	neighbours := map[Direction]Cube{}

	neighbours[LEFT] = Cube{
		X: c.X - 1,
		Y: c.Y,
		Z: c.Z,
	}
	neighbours[RIGHT] = Cube{
		X: c.X + 1,
		Y: c.Y,
		Z: c.Z,
	}
	neighbours[BOTTOM] = Cube{
		X: c.X,
		Y: c.Y - 1,
		Z: c.Z,
	}
	neighbours[TOP] = Cube{
		X: c.X,
		Y: c.Y + 1,
		Z: c.Z,
	}
	neighbours[BACK] = Cube{
		X: c.X,
		Y: c.Y,
		Z: c.Z - 1,
	}
	neighbours[FRONT] = Cube{
		X: c.X,
		Y: c.Y,
		Z: c.Z + 1,
	}

	return neighbours
}
