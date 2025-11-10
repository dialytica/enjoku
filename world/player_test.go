package world_test

import (
	"testing"

	"github.com/dialytica/enjoku/world"
)

func TestPlayerPosition(t *testing.T) {
	playerPosition := world.PlayerPosition{}

	t.Run("initial player position is 0,0", func(t *testing.T) {
		if x, y := playerPosition.GetPosition(); x != 0 || y != 0 {
			t.Errorf("want player position at x:0, y:0 got x:%d, y:%d", x, y)
		}
	})

	t.Run("set player position replace with new value", func(t *testing.T) {
		playerPosition.SetPosition(2, 3)
		if x, y := playerPosition.GetPosition(); x != 2 || y != 3 {
			t.Errorf("want player position at x:0, y:0 got x:%d, y:%d", x, y)
		}
	})
}

func TestPlayer(t *testing.T) {
	playerName := "testPlayer"
	player := world.CreateNewPlayer(playerName)

	t.Run("newly created player name is as provided", func(t *testing.T) {
		if player.Name != playerName {
			t.Errorf("unexpected player name, want:%s, got:%s", playerName, player.Name)
		}
	})

	t.Run("player ID is auto generated", func(t *testing.T) {
		if player.ID == "" {
			t.Error("player ID is empty string")
		}
	})

	t.Run("player position is auto generated", func(t *testing.T) {
		if player.Position == nil {
			t.Fatal("player position is nil")
		}
		if x, y := player.GetPosition(); x != 0 || y != 0 {
			t.Errorf("player position should be initiated at (%d,%d)", x, y)
		}
	})

	t.Run("player chunk ID is still empty", func(t *testing.T) {
		if player.ChunkID != "" {
			t.Errorf("player chunkID is not empty")
		}
	})
}
