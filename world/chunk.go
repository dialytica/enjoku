package world

import "github.com/google/uuid"

type ChunkPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c *ChunkPosition) GetPosition() (int, int) {
	return c.X, c.Y
}

func (c *ChunkPosition) SetPosition(x, y int) {
	
	c.X = x
	c.Y = y
}

func (c *ChunkPosition) TranslateNew(x, y int) *ChunkPosition {
	return &ChunkPosition{
		X: c.X + x,
		Y: c.Y + y,
	}
}

type ChunkGraph struct {
	ID       string         `json:"id"`
	Position *ChunkPosition `json:"position"`

	Length int `json:"length"`
	Width  int `json:"width"`

	North *ChunkGraph `json:"-"`
	South *ChunkGraph `json:"-"`
	West  *ChunkGraph `json:"-"`
	East  *ChunkGraph `json:"-"`

	PlayerIDsPosition map[PlayerPosition]string `json:"-"`
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
