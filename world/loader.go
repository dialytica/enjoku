package world

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// flattenWorld is flatten representation fo the world data in JSON format, it contains
// reference id that useful for getting user / chunk data that stored elsewhere
type flattenWorld struct {
	ID      string   `json:"id"`
	Name    string   `json:"Name"`
	Chunks  []string `json:"chunks"`
	Players []string `json:"players"`
}

func LoadJSONWorld(worldID string) (*World, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf(
			"$HOME environment variable not found, error: %s", err)
	}

	gamepath := filepath.Join(homePath, ".local", "share", "enjoku", worldID)
	playersPath := filepath.Join(gamepath, "players")
	chunksPath := filepath.Join(gamepath, "chunks")
	err = os.MkdirAll(playersPath, 0750)
	if err != nil {
		return nil, fmt.Errorf("players path error: %s", err)
	}
	err = os.MkdirAll(chunksPath, 0750)
	if err != nil {
		return nil, fmt.Errorf("chunks path error: %s", err)
	}

	worldPath := filepath.Join(gamepath, "world.json")
	worldFile, err := os.Open(worldPath)
	if err != nil {
		return nil, fmt.Errorf("world file not found, error: %s", err)
	}
	defer worldFile.Close()

	flattenWorldEncoder := json.NewDecoder(worldFile)
	var flatWorld flattenWorld
	err = flattenWorldEncoder.Decode(&flatWorld)
	if err != nil {
		return nil, fmt.Errorf("World Data Decode error: %s", err)
	}

	gameWorld := &World{
		Name:            flatWorld.Name,
		ID:              flatWorld.ID,
		ChunkIDPosition: make(map[ChunkPosition]string),
		Chunks:          make(map[string]*ChunkGraph),
		Players:         make(map[string]*Player),
	}

	for _, chunkRefID := range flatWorld.Chunks {
		chunkPath := fmt.Sprintf("%s.json", filepath.Join(chunksPath, chunkRefID))
		chunk, err := loadJSONChunk(chunkPath)
		if err != nil {
			// TODO: make a better chunk error handling later
			continue
		}
		gameWorld.Chunks[chunk.ID] = chunk
		gameWorld.ChunkIDPosition[*chunk.Position] = chunk.ID
	}

	for _, playerRefID := range flatWorld.Players {
		playerPath := fmt.Sprintf("%s.json", filepath.Join(playersPath, playerRefID))
		player, err := loadJSONPlayer(playerPath)
		if err != nil {
			// TODO: make a better player error handling later
			continue
		}
		gameWorld.Players[player.ID] = player
		playerChunk, ok := gameWorld.Chunks[player.ChunkID]
		if !ok {
			// TODO: make a better player chunk error handling
			continue
		}
		playerChunk.InsertPlayerID(player.ID, *player.Position)
	}

	return gameWorld, nil
}

func loadJSONChunk(chunkPath string) (*ChunkGraph, error) {
	chunkFile, err := os.Open(chunkPath)
	if err != nil {
		return nil, err
	}
	defer chunkFile.Close()

	var chunk ChunkGraph
	chunkDecoder := json.NewDecoder(chunkFile)

	err = chunkDecoder.Decode(&chunk)
	if err != nil {
		return nil, err
	}

	return &chunk, nil
}

func loadJSONPlayer(playerPath string) (*Player, error) {
	playerFile, err := os.Open(playerPath)
	if err != nil {
		return nil, err
	}
	defer playerFile.Close()

	var player Player
	playerDecoder := json.NewDecoder(playerFile)

	err = playerDecoder.Decode(&player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
