package world

type ChunkPosition struct {
	x int
	y int
}

func (c *ChunkPosition) GetPosition() (int, int) {
	return c.x, c.y
}

func (c *ChunkPosition) SetPosition(x, y int) {
	c.x = x
	c.y = y
}

type ChunkGraph struct {
	ID       string
	Position ChunkPosition
	Length   int
	Width    int
	North    *ChunkGraph
	South    *ChunkGraph
	West     *ChunkGraph
	East     *ChunkGraph
}

func (c *ChunkGraph) GetPosition() (int, int) {
	return c.Position.GetPosition()
}

func (c *ChunkGraph) SetPosition(x, y int) {
	c.Position.SetPosition(x, y)
}

func (c *ChunkGraph) Navigate(direction string) *ChunkGraph {
	switch direction {
	case North:
		return c.North
	case South:
		return c.South
	case West:
		return c.West
	case East:
		return c.East
	default:
		return nil
	}
}
