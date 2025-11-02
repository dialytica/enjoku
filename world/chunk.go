package world

import "github.com/google/uuid"

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

func (c *ChunkPosition) TranslateNew(x, y int) *ChunkPosition {
	return &ChunkPosition{
		x: c.x + x,
		y: c.y + y,
	}
}

type ChunkGraph struct {
	ID                string
	Position          *ChunkPosition
	PlayerIDsPosition map[PlayerPosition]string

	Length int
	Width  int

	North *ChunkGraph
	South *ChunkGraph
	West  *ChunkGraph
	East  *ChunkGraph
}

func CreateNewChunk(x, y int) *ChunkGraph {
	chunkPosition := &ChunkPosition{
		x,
		y,
	}

	chunk := &ChunkGraph{
		PlayerIDsPosition: make(map[PlayerPosition]string),

		ID:       uuid.NewString(),
		Position: chunkPosition,
		Length:   ChunkLenght,
		Width:    ChunkWidth,
	}
	return chunk
}

func (c *ChunkGraph) GetPosition() (int, int) {
	return c.Position.GetPosition()
}

func (c *ChunkGraph) SetPosition(x, y int) {
	c.Position.SetPosition(x, y)
}

func (c *ChunkGraph) InsertPlayerID(playerID string, position PlayerPosition) {
	c.PlayerIDsPosition[position] = playerID
}

func (c *ChunkGraph) RemovePlayerID(position PlayerPosition) string {
	if playerID, ok := c.PlayerIDsPosition[position]; ok {
		delete(c.PlayerIDsPosition, position)
		return playerID
	}
	return ""
}

func (c *ChunkGraph) SetAdjacentChunkByDirection(direction string, chunk *ChunkGraph) {
	switch direction {
	case North:
		c.North = chunk
		chunk.South = c
	case South:
		c.South = chunk
		chunk.North = c
	case West:
		c.West = chunk
		chunk.East = c
	case East:
		c.East = chunk
		chunk.West = c
	}
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
