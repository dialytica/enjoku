package world_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dialytica/enjoku/world"
)

func TestSaveJSONWorld(t *testing.T) {
	gameWorld := world.CreateNewWorld(nil, nil)
	gameWorld.Name = "Test World"
	gameWorld.ID = "0bbe2a06-8a6b-4131-837a-136c565fb38b"

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

	_, err = os.Stat(worldPath)
	if err != nil {
		t.Fatal("world file not found", err)
	}
}
