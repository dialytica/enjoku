package world_test

import (
	"testing"

	"github.com/dialytica/enjoku/world"
)

func TestChunkPosition(t *testing.T) {
	chunkPosition := world.ChunkPosition{}

	t.Run("initial position is 0,0", func(t *testing.T) {
		if x, y := chunkPosition.GetPosition(); x != 0 || y != 0 {
			t.Errorf("want chunk position x:0, y:0 but got x:%d, y:%d", x, y)
		}
	})

	t.Run("set position replace its value", func(t *testing.T) {
		chunkPosition.SetPosition(2, 3)
		if x, y := chunkPosition.GetPosition(); x != 2 || y != 3 {
			t.Errorf("want chunk position x:2, y:3 but got x:%d, y:%d", x, y)
		}
	})

	t.Run("translate position adds new chunkPosition with new value", func(t *testing.T) {
		chunkPosition.TranslateNew(2, 3)
		newChunkP := chunkPosition.TranslateNew(1, -2)
		if x, y := newChunkP.GetPosition(); x != 3 || y != 1 {
			t.Errorf("want new chunk position x:3, y:1 but got x:%d, y:%d", x, y)
		}
	})
}

func TestChunkGraph(t *testing.T) {
	chunk := world.CreateNewChunk(1, 1)
	playerID := "61054519-421e-431e-a9db-e50f852eceea"

	t.Run("chunk position is as initiated", func(t *testing.T) {
		if x, y := chunk.GetPosition(); x != 1 || y != 1 {
			t.Errorf("want position at (%d,%d) got (%d,%d)", 1, 1, x, y)
		}
	})

	t.Run("chunk ID is auto generated", func(t *testing.T) {
		if chunk.ID == "" {
			t.Error("chunk ID is empty")
		}
	})

	t.Run("chunk PlayerIDs map is empty", func(t *testing.T) {
		if len(chunk.PlayerIDsPosition) != 0 {
			t.Fatal("player ID is not empty")
		}
	})

	t.Run("insert playerID position", func(t *testing.T) {
		chunk.InsertPlayerID(playerID, world.PlayerPosition{})

		if len(chunk.PlayerIDsPosition) == 0 {
			t.Fatal("player ID is empty")
		}

		if id, ok := chunk.PlayerIDsPosition[world.PlayerPosition{}]; !ok || id != playerID {
			t.Fatalf("want player ID: %s but got: %s", playerID, id)
		}
	})

	t.Run("remove playerID", func(t *testing.T) {
		id := chunk.RemovePlayerID(world.PlayerPosition{})

		if id != playerID {
			t.Fatalf("want to remove player ID: %s but got: %s", playerID, id)
		}

		if len(chunk.PlayerIDsPosition) != 0 {
			t.Fatal("player ID is not empty")
		}
	})

	t.Run("set adjacent chunk by direction", func(t *testing.T) {
		t.Run("North", func(t *testing.T) {
			newChunk := world.CreateNewChunk(1, 2)
			chunk.SetAdjacentChunkByDirection(world.North, newChunk)
			if chunk.North == nil {
				t.Fatal("chunk North is empty")
			}
			if chunk.North.South == nil || chunk.North.South != chunk {
				t.Fatal("chunk North is not connected")
			}
		})

		t.Run("South", func(t *testing.T) {
			newChunk := world.CreateNewChunk(1, 0)
			chunk.SetAdjacentChunkByDirection(world.South, newChunk)
			if chunk.South == nil {
				t.Fatal("chunk South is empty")
			}
			if chunk.South.North == nil || chunk.South.North != chunk {
				t.Fatal("chunk South is not connected")
			}
		})

		t.Run("East", func(t *testing.T) {
			newChunk := world.CreateNewChunk(2, 1)
			chunk.SetAdjacentChunkByDirection(world.East, newChunk)
			if chunk.East == nil {
				t.Fatal("chunk East is empty")
			}
			if chunk.East.West == nil || chunk.East.West != chunk {
				t.Fatal("chunk East is not connected")
			}
		})

		t.Run("West", func(t *testing.T) {
			newChunk := world.CreateNewChunk(0, 1)
			chunk.SetAdjacentChunkByDirection(world.West, newChunk)
			if chunk.West == nil {
				t.Fatal("chunk West is empty")
			}
			if chunk.West.East == nil || chunk.West.East != chunk {
				t.Fatal("chunk West is not connected")
			}
		})
	})
}
