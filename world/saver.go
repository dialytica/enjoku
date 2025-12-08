package world

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSONWorld(gameWorld *World) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf(
			"$HOME environment variable not found, error: %s", err)
	}

	gamepath := filepath.Join(homePath, ".local", "share", "enjoku", gameWorld.ID)
	playersPath := filepath.Join(gamepath, "players")
	chunksPath := filepath.Join(gamepath, "chunks")
	err = os.MkdirAll(playersPath, 0750)
	if err != nil {
		return fmt.Errorf("players path error: %s", err)
	}
	err = os.MkdirAll(chunksPath, 0750)
	if err != nil {
		return fmt.Errorf("chunks path error: %s", err)
	}

	flatWorld := FlattenWorld{
		ID:      gameWorld.ID,
		Name:    gameWorld.Name,
		Chunks:  make([]string, 0),
		Players: make([]string, 0),
	}

	for id, chunk := range gameWorld.Chunks {
		flatWorld.Chunks = append(flatWorld.Chunks, id)
		err = saveJSONChunk(chunk, chunksPath)
		if err != nil {
			// TODO: make a better chunk saving error handling
			continue
		}
	}

	for id, player := range gameWorld.Players {
		flatWorld.Players = append(flatWorld.Players, id)
		err = saveJSONPlayer(player, playersPath)
		if err != nil {
			// TODO: make a better player saving error handling
			continue
		}
	}

	worldPath := filepath.Join(gamepath, "world.json")
	worldFile, err := os.Create(worldPath)
	if err != nil {
		return fmt.Errorf("world file error: %s", err)
	}
	defer worldFile.Close()

	worldEncoder := json.NewEncoder(worldFile)
	worldEncoder.SetIndent("", "\t")
	if err = worldEncoder.Encode(&flatWorld); err != nil {
		return err
	}

	return nil
}

func saveJSONChunk(chunk *ChunkGraph, chunksPath string) error {
	chunkPath := fmt.Sprintf(
		"%s.json", filepath.Join(chunksPath, chunk.ID))
	chunkFile, err := os.Create(chunkPath)
	if err != nil {
		return err
	}
	defer chunkFile.Close()

	chunkEncoder := json.NewEncoder(chunkFile)
	chunkEncoder.SetIndent("", "\t")
	if err = chunkEncoder.Encode(chunk); err != nil {
		return err
	}

	return nil
}

func saveJSONPlayer(player *Player, playersPath string) error {
	playerPath := fmt.Sprintf(
		"%s.json", filepath.Join(playersPath, player.ID))
	playerFile, err := os.Create(playerPath)
	if err != nil {
		return err
	}
	defer playerFile.Close()

	playerEncoder := json.NewEncoder(playerFile)
	playerEncoder.SetIndent("", "\t")
	if err = playerEncoder.Encode(player); err != nil {
		return err
	}

	return nil
}
