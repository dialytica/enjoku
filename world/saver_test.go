package world_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/dialytica/enjoku/world"
)

func TestSaveJSONWorld(t *testing.T) {
	gameWorld := world.CreateNewWorld(nil, nil)
	gameWorld.Name = "Test World"
	gameWorld.ID = "0bbe2a06-8a6b-4131-837a-136c565fb38b"
	var gameChunkID string
	for id := range gameWorld.Chunks {
		gameChunkID = id
	}
	var gamePlayerID string
	for id := range gameWorld.Players {
		gamePlayerID = id
	}

	err := world.SaveJSONWorld(gameWorld)
	if err != nil {
		t.Fatal("Saving World into json file failed, error: ", err)
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("$HOME environment variable not found", err)
	}

	gamePath := filepath.Join(homePath,
		".local", "share", "enjoku", gameWorld.ID)
	worldPath := filepath.Join(gamePath, "world.json")

	worldFile, err := os.Open(worldPath)
	if err != nil {
		t.Fatal("world file not found", err)
	}

	var worldFlatJSON world.FlattenWorld
	t.Run("validating world json file", func(t *testing.T) {
		worldDecoder := json.NewDecoder(worldFile)

		err = worldDecoder.Decode(&worldFlatJSON)
		if err != nil {
			t.Fatalf("failed to load world.json, error: %s", err)
		}

		if worldFlatJSON.ID != gameWorld.ID {
			t.Errorf("expecting ID as %s but got %s",
				gameWorld.ID, worldFlatJSON.ID)
		}

		if worldFlatJSON.Name != gameWorld.Name {
			t.Errorf("expecting Name as %s but got %s",
				gameWorld.Name, worldFlatJSON.Name)
		}

		for _, chunkID := range worldFlatJSON.Chunks {
			_, ok := gameWorld.Chunks[chunkID]
			if !ok {
				t.Errorf("expecting chunk %s exists", chunkID)
				break
			}
		}

		for _, playerID := range worldFlatJSON.Players {
			_, ok := gameWorld.Players[playerID]
			if !ok {
				t.Errorf("expecting player %s exists", playerID)
				break
			}
		}
	})

	t.Run("validating chunks json file", func(t *testing.T) {
		chunkPath := filepath.Join(gamePath, "chunks",
			gameChunkID+".json")

		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			t.Fatalf("failed to open chunks/%s.json, error: %s",
				gameChunkID, err)
		}

		chunkDecoder := json.NewDecoder(chunkFile)

		var chunk world.ChunkGraph
		err = chunkDecoder.Decode(&chunk)
		if err != nil {
			t.Fatalf("failed to load chunks/%s.json, error: %s",
				gameChunkID, err)
		}
		_, ok := gameWorld.Chunks[chunk.ID]
		if !ok {
			t.Fatalf("chunk %s is not found in game world", chunk.ID)
		}

	})

	t.Run("validating players json file", func(t *testing.T) {
		playerPath := filepath.Join(gamePath, "players",
			gamePlayerID+".json")

		playerFile, err := os.Open(playerPath)
		if err != nil {
			t.Fatalf("failed to open players/%s.json, error: %s",
				gamePlayerID, err)
		}

		playerDecoder := json.NewDecoder(playerFile)

		var player world.Player
		err = playerDecoder.Decode(&player)
		if err != nil {
			t.Fatalf("failed to load players/%s.json, error: %s",
				gamePlayerID, err)
		}
		_, ok := gameWorld.Players[player.ID]
		if !ok {
			t.Fatalf("player %s is not found in game world", player.ID)
		}

	})
}
