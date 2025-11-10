/* Package world shows how world are created */
package world

import (
	"fmt"
	"log"
	"math"
)

const (
	North = "North"
	South = "South"

	West = "West"
	East = "East"

	ChunkLenght = 64
	ChunkWidth  = 24

	DefaultName = "NJQ"
)

type PlayerIDNotFoundError string

func (p PlayerIDNotFoundError) Error() string {
	return fmt.Sprintf("playerID: %s is not found", string(p))
}

type ChunkIDNotFoundError string

func (c ChunkIDNotFoundError) Error() string {
	return fmt.Sprintf("chunkID: %s is not found", string(c))
}

type IPosition interface {
	GetPosition() (int, int)
	SetPosition(x, y int)
}

type World struct {
	ChunkIDPosition map[ChunkPosition]string
	Chunks          map[string]*ChunkGraph
	Players         map[string]*Player
}

func CreateNewWorld(chunk *ChunkGraph, player *Player) *World {
	if chunk == nil {
		chunk = CreateNewChunk(0, 0)
	}
	if player == nil {
		player = CreateNewPlayer(DefaultName)
		player.ChunkID = chunk.ID
	}
	world := &World{
		ChunkIDPosition: map[ChunkPosition]string{
			*chunk.Position: chunk.ID,
		},
		Chunks: map[string]*ChunkGraph{
			chunk.ID: chunk,
		},
		Players: map[string]*Player{player.ID: player},
	}
	return world
}

func LoadWorld(playerName, playerID string) *World {
	// TODO: load world from file and check playerID is exist in the world

	player := CreateNewPlayer(playerName)
	player.ID = playerID
	newWorld := CreateNewWorld(nil, player)
	chunkID := newWorld.ChunkIDPosition[ChunkPosition{x: 0, y: 0}]
	player.ChunkID = chunkID
	chunk := newWorld.Chunks[chunkID]
	chunk.InsertPlayerID(playerID, *player.Position)
	return newWorld
}

func (w *World) MovePlayer(playerID string, x, y int) error {
	player, ok := w.Players[playerID]
	if !ok {
		return PlayerIDNotFoundError(playerID)
	}
	playerChunk, ok := w.Chunks[player.ChunkID]
	if !ok {
		return ChunkIDNotFoundError(player.ChunkID)
	}
	playerChunk.RemovePlayerID(*player.Position)

	if math.Abs(float64(x)) > float64(playerChunk.Length)/2 {
		if x > 0 {
			w.LoadAdjacentChunks(playerChunk)
			x = -playerChunk.Length / 2
			player.ChunkID = playerChunk.East.ID
			w.ChunkIDPosition[*playerChunk.East.Position] = player.ChunkID
		} else {
			w.LoadAdjacentChunks(playerChunk)
			x = playerChunk.Length / 2
			player.ChunkID = playerChunk.West.ID
			w.ChunkIDPosition[*playerChunk.West.Position] = player.ChunkID
		}
	}
	if math.Abs(float64(y)) > float64(playerChunk.Width)/2 {
		if y > 0 {
			w.LoadAdjacentChunks(playerChunk)
			y = -playerChunk.Width / 2
			player.ChunkID = playerChunk.North.ID
			w.ChunkIDPosition[*playerChunk.North.Position] = player.ChunkID
		} else {
			w.LoadAdjacentChunks(playerChunk)
			y = playerChunk.Width / 2
			player.ChunkID = playerChunk.South.ID
			w.ChunkIDPosition[*playerChunk.South.Position] = player.ChunkID
		}
	}

	log.Printf("player: %+v\n", player)
	log.Printf("chunk: %+v\n", playerChunk)

	playerChunk, ok = w.Chunks[player.ChunkID]
	if !ok {
		return ChunkIDNotFoundError(player.ChunkID)
	}
	player.SetPosition(x, y)
	playerChunk.InsertPlayerID(playerID, *player.Position)
	return nil
}

func (w *World) LoadAdjacentChunks(currentChunk *ChunkGraph) {
	x, y := currentChunk.GetPosition()
	chunkID := w.ChunkIDPosition[*currentChunk.Position.TranslateNew(0, 1)]
	chunkNorth, ok := w.Chunks[chunkID]
	if !ok {
		chunkNorth = CreateNewChunk(x, y+1)
		w.Chunks[chunkNorth.ID] = chunkNorth
		w.Chunks[fmt.Sprintf("%d:%d", x, y+1)] = chunkNorth
	}
	if currentChunk.North == nil {
		currentChunk.SetAdjacentChunkByDirection(North, chunkNorth)
	}

	chunkID = w.ChunkIDPosition[*currentChunk.Position.TranslateNew(0, -1)]
	chunkSouth, ok := w.Chunks[chunkID]
	if !ok {
		chunkSouth = CreateNewChunk(x, y-1)
		w.Chunks[chunkSouth.ID] = chunkSouth
		w.Chunks[fmt.Sprintf("%d:%d", x, y-1)] = chunkSouth
	}
	if currentChunk.South == nil {
		currentChunk.SetAdjacentChunkByDirection(South, chunkSouth)
	}

	chunkID = w.ChunkIDPosition[*currentChunk.Position.TranslateNew(1, 0)]
	chunkEast, ok := w.Chunks[chunkID]
	if !ok {
		chunkEast = CreateNewChunk(x+1, y)
		w.Chunks[chunkEast.ID] = chunkEast
		w.Chunks[fmt.Sprintf("%d:%d", x+1, y)] = chunkEast
	}
	if currentChunk.East == nil {
		currentChunk.SetAdjacentChunkByDirection(East, chunkEast)
	}

	chunkID = w.ChunkIDPosition[*currentChunk.Position.TranslateNew(-1, 0)]
	chunkWest, ok := w.Chunks[chunkID]
	if !ok {
		chunkWest = CreateNewChunk(x-1, y)
		w.Chunks[chunkWest.ID] = chunkWest
		w.Chunks[fmt.Sprintf("%d:%d", x-1, y)] = chunkWest
	}
	if currentChunk.West == nil {
		currentChunk.SetAdjacentChunkByDirection(West, chunkWest)
	}
}
