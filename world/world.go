/* Package world shows how world are created */
package world

import "fmt"

const (
	North = "North"
	South = "South"

	West = "West"
	East = "East"

	ChunkLenght = 128
	ChunkWidth  = 96

	DefaultName = "NJQ"
)

type IPosition interface {
	GetPosition() (int, int)
	SetPosition(x, y int)
}

type World struct {
	Chunks  map[string]*ChunkGraph
	Players []*Player
}

func CreateNewWorld() *World {
	chunk := CreateNewChunk(0, 0)
	player := CreateNewPlayer(DefaultName)
	player.ChunkID = &chunk.ID
	world := &World{
		Chunks: map[string]*ChunkGraph{
			"0:0":    chunk,
			chunk.ID: chunk,
		},
		Players: []*Player{player},
	}
	return world
}

func LoadWorld() *World {
	newWorld := CreateNewWorld()
	return newWorld
}

func (w *World) LoadAdjacentChunks(currentChunk *ChunkGraph) {
	x, y := currentChunk.GetPosition()
	chunkNorth, ok := w.Chunks[fmt.Sprintf("%d:%d", x, y+1)]
	if !ok {
		chunkNorth = CreateNewChunk(x, y+1)
		w.Chunks[chunkNorth.ID] = chunkNorth
		w.Chunks[fmt.Sprintf("%d:%d", x, y+1)] = chunkNorth
	}
	currentChunk.SetAdjacentChunkByDirection(North, chunkNorth)
	chunkSouth, ok := w.Chunks[fmt.Sprintf("%d:%d", x, y-1)]
	if !ok {
		chunkSouth = CreateNewChunk(x, y-1)
		w.Chunks[chunkSouth.ID] = chunkSouth
		w.Chunks[fmt.Sprintf("%d:%d", x, y-1)] = chunkSouth
	}
	currentChunk.SetAdjacentChunkByDirection(South, chunkSouth)
	chunkEast, ok := w.Chunks[fmt.Sprintf("%d:%d", x+1, y)]
	if !ok {
		chunkEast = CreateNewChunk(x+1, y)
		w.Chunks[chunkEast.ID] = chunkEast
		w.Chunks[fmt.Sprintf("%d:%d", x+1, y)] = chunkEast
	}
	currentChunk.SetAdjacentChunkByDirection(East, chunkEast)
	chunkWest, ok := w.Chunks[fmt.Sprintf("%d:%d", x-1, y)]
	if !ok {
		chunkWest = CreateNewChunk(x-1, y)
		w.Chunks[chunkWest.ID] = chunkWest
		w.Chunks[fmt.Sprintf("%d:%d", x-1, y)] = chunkWest
	}
	currentChunk.SetAdjacentChunkByDirection(West, chunkWest)
}
