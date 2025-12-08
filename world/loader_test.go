package world_test

import (
	"testing"

	"github.com/dialytica/enjoku/world"
)

func TestLoadJSONWorld(t *testing.T) {
	// Setup
	gameWorld := world.CreateNewWorld(nil, nil)
	gameWorld.Name = "Test World Loading"
	gameWorld.ID = "05920c7e-b5f9-440b-b619-c10238b9eabf"
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

	// Testing Scenarios

	loadedWorld, err := world.LoadJSONWorld(gameWorld.ID)
	if err != nil {
		t.Fatalf("failed to load game world. error: %s", err)
	}

	if loadedWorld.ID != gameWorld.ID {
		t.Errorf("expecting world ID %s but got %s",
			gameWorld.ID, loadedWorld.ID)
	}

	if _, ok := loadedWorld.Chunks[gameChunkID]; !ok {
		t.Errorf("expecting chunkID %s exist", gameChunkID)
	}

	if _, ok := loadedWorld.Players[gamePlayerID]; !ok {
		t.Errorf("expecting playerID %s exist", gamePlayerID)
	}
}
