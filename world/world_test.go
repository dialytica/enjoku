package world_test

import (
	"errors"
	"testing"

	"github.com/dialytica/enjoku/world"
)

func TestCreateNewWorld(t *testing.T) {
	t.Run("empty world", func(t *testing.T) {
		gameWorld := world.CreateNewWorld(nil, nil)

		if len(gameWorld.Chunks) != 1 {
			t.Fatal("initial chunk is not generated")
		}

		if len(gameWorld.Players) != 1 {
			t.Fatal("initial player is not generated")
		}

		if len(gameWorld.ChunkIDPosition) != 1 {
			t.Fatal("chunk ID position is not mapped")
		}
	})

	t.Run("world with existing player", func(t *testing.T) {
		playerName := "testPlayer"
		player := world.CreateNewPlayer(playerName)
		gameWorld := world.CreateNewWorld(nil, player)

		if len(gameWorld.Chunks) != 1 {
			t.Fatal("initial chunk is not generated")
		}

		if len(gameWorld.Players) != 1 {
			t.Fatal("existing player is not added")
		}

		if _, ok := gameWorld.Players[player.ID]; !ok {
			t.Fatal("existing player ID is not found")
		}

		if len(gameWorld.ChunkIDPosition) != 1 {
			t.Fatal("chunk ID position is not mapped")
		}
	})

	t.Run("world with existing chunk", func(t *testing.T) {
		chunk := world.CreateNewChunk(0, 0)
		gameWorld := world.CreateNewWorld(chunk, nil)

		if len(gameWorld.Chunks) != 1 {
			t.Fatal("existing chunk is not added")
		}

		if _, ok := gameWorld.Chunks[chunk.ID]; !ok {
			t.Fatal("existing chunk ID is not found")
		}

		if len(gameWorld.Players) != 1 {
			t.Fatal("initial player is not generated")
		}

		if len(gameWorld.ChunkIDPosition) != 1 {
			t.Fatal("chunk ID position is not mapped")
		}
	})
}

func TestLoadWorld(t *testing.T) {
	playerName := "TesPlayer"
	playerID := "ffe7e548-30ab-4e1b-9746-eb55ad2e4946"
	worldID := "2e919013-87c9-4821-a446-737c48983eda"

	t.Run("non existent worldID", func(t *testing.T) {
		gameWorld := world.LoadWorld(playerName, playerID,
			"5616479e-c421-4e55-ad38-224f177a5c44")
		if gameWorld != nil {
			t.Errorf("the game world ID %s should not be exist, you must be lucky",
				"5616479e-c421-4e55-ad38-224f177a5c44")
		}
	})

	// Setup
	savedPlayer := world.CreateNewPlayer(playerName)
	savedPlayer.ID = playerID
	savedWorld := world.CreateNewWorld(nil,
		savedPlayer)
	savedWorld.ID = worldID

	err := world.SaveJSONWorld(savedWorld)
	if err != nil {
		t.Fatalf("setup save world failed. error: %s", err)
	}

	t.Run("load saved World", func(t *testing.T) {
		loadedWorld := world.LoadWorld(playerName, playerID, worldID)
		if loadedWorld == nil {
			t.Fatal("Failed to load game world")
		}

		if loadedWorld.ID != savedWorld.ID {
			t.Errorf("expecting worldID %s but got %s",
				loadedWorld.ID, savedWorld.ID)
		}

		if len(loadedWorld.Players) != 1 {
			t.Errorf("unexpected number of players. players: %+v",
				loadedWorld.Players)
		}
	})
}

func TestMovePlayer(t *testing.T) {
	playerName := "testPlayerMovement"
	player := world.CreateNewPlayer(playerName)
	gameWorld := world.CreateNewWorld(nil, player)
	chunkID := gameWorld.ChunkIDPosition[world.ChunkPosition{}]
	player.ChunkID = chunkID
	chunk := gameWorld.Chunks[chunkID]
	chunk.InsertPlayerID(player.ID, *player.Position)

	t.Run("inside chunk", func(t *testing.T) {
		// move to the north 1 step
		gameWorld.MovePlayer(player.ID, 0, 1)
		if x, y := player.GetPosition(); x != 0 || y != 1 {
			t.Fatal("player is not moving north")
		}
		// move to the east 1 step
		gameWorld.MovePlayer(player.ID, 1, 0)
		if x, y := player.GetPosition(); x != 1 || y != 0 {
			t.Fatal("player is not moving east")
		}
		// move to the south 1 step
		gameWorld.MovePlayer(player.ID, 0, -1)
		if x, y := player.GetPosition(); x != 0 || y != -1 {
			t.Fatal("player is not moving south")
		}
		// move to the west 1 step
		gameWorld.MovePlayer(player.ID, -1, 0)
		if x, y := player.GetPosition(); x != -1 || y != 0 {
			t.Fatal("player is not moving west")
		}
	})

	t.Run("across chunk", func(t *testing.T) {
		// move to the north chunk
		err := gameWorld.MovePlayer(player.ID, 0, world.ChunkWidth)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if x, y := player.GetPosition(); x != 0 || y != -world.ChunkWidth/2 {
			t.Fatalf("player is not moving north, expected: (%d,%d) got: (%d, %d)",
				0, -world.ChunkWidth/2, x, y)
		}
		if player.ChunkID != gameWorld.ChunkIDPosition[*chunk.North.Position] {
			t.Logf("chunk(%+v): %s", chunk.Position, chunk.ID)
			t.Logf("chunk.North(%+v): %s", chunk.North.Position, chunk.North.ID)
			t.Fatalf("player is not moving to north chunk, expected: %s got:%s",
				player.ChunkID, gameWorld.ChunkIDPosition[*chunk.North.Position])
		}
		// move to the east chunk
		gameWorld.MovePlayer(player.ID, world.ChunkLenght, 0)
		if x, y := player.GetPosition(); x != -world.ChunkLenght/2 || y != 0 {
			t.Fatalf("player is not moving east, expected: (%d,%d) got: (%d, %d)",
				-world.ChunkLenght/2, 0, x, y)
		}
		if player.ChunkID != gameWorld.ChunkIDPosition[*chunk.North.East.Position] {
			t.Logf("chunk.North(%+v): %s", chunk.North.Position, chunk.ID)
			t.Logf("chunk.North.East(%+v): %s", chunk.North.East.Position, chunk.North.East.ID)
			t.Fatalf("player is not moving to east chunk, expected: %s got:%s",
				player.ChunkID, gameWorld.ChunkIDPosition[*chunk.North.East.Position])
		}
		// move to the south chunk
		gameWorld.MovePlayer(player.ID, 0, -world.ChunkWidth)
		if x, y := player.GetPosition(); x != 0 || y != world.ChunkWidth/2 {
			t.Fatalf("player is not moving south, expected: (%d,%d) got: (%d, %d)",
				0, world.ChunkLenght/2, x, y)
		}
		if player.ChunkID != gameWorld.ChunkIDPosition[*chunk.East.Position] {
			t.Logf("chunk(%+v): %s", chunk.Position, chunk.ID)
			t.Logf("chunk.East(%+v): %s", chunk.East.Position, chunk.East.ID)
			t.Fatalf("player is not moving to south chunk, expected: %s got:%s",
				player.ChunkID, gameWorld.ChunkIDPosition[*chunk.East.Position])
		}
		// move to the west chunk
		gameWorld.MovePlayer(player.ID, -world.ChunkLenght, 0)
		if x, y := player.GetPosition(); x != world.ChunkLenght/2 || y != 0 {
			t.Fatalf("player is not moving west, expected: (%d,%d) got: (%d, %d)",
				world.ChunkLenght/2, 0, x, y)
		}
		if player.ChunkID != gameWorld.ChunkIDPosition[*chunk.Position] {
			t.Logf("chunk(%+v): %s", chunk.Position, chunk.ID)
			t.Fatalf("player is not moving to west chunk, expected: %s got:%s",
				player.ChunkID, gameWorld.ChunkIDPosition[*chunk.Position])
		}
	})

	t.Run("player not registered", func(t *testing.T) {
		err := gameWorld.MovePlayer("bc2f94c2-ac9f-4454-8049-fcedb5f7865a", 1, 2)
		if !errors.Is(err, world.PlayerIDNotFoundError("bc2f94c2-ac9f-4454-8049-fcedb5f7865a")) {
			t.Fatal("unidentifed error: ", err)
		}
	})
}

func TestLoadAdjacentChunks(t *testing.T) {
	chunk := world.CreateNewChunk(0, 0)
	gameWorld := world.CreateNewWorld(chunk, nil)
	chunkIDs := make(map[string]string)

	t.Run("adjacent chunk generated if empty", func(t *testing.T) {
		gameWorld.LoadAdjacentChunks(chunk)
		if chunk.North == nil {
			t.Fatalf("chunk North is not loaded")
		}
		chunkIDs[world.North] = chunk.North.ID
		if chunk.South == nil {
			t.Fatalf("chunk South is not loaded")
		}
		chunkIDs[world.South] = chunk.South.ID
		if chunk.East == nil {
			t.Fatalf("chunk East is not loaded")
		}
		chunkIDs[world.East] = chunk.East.ID
		if chunk.West == nil {
			t.Fatalf("chunk West is not loaded")
		}
		chunkIDs[world.West] = chunk.West.ID
	})

	t.Run("loaded chunk still have the same id", func(t *testing.T) {
		gameWorld.LoadAdjacentChunks(chunk)
		if chunkIDs[world.North] != chunk.North.ID {
			t.Fatalf("chunk North mismatch, expected: %s got: %s",
				chunkIDs[world.North], chunk.North.ID)
		}
		if chunkIDs[world.South] != chunk.South.ID {
			t.Fatalf("chunk South mismatch, expected: %s got: %s",
				chunkIDs[world.South], chunk.South.ID)
		}
		if chunkIDs[world.East] != chunk.East.ID {
			t.Fatalf("chunk East mismatch, expected: %s got: %s",
				chunkIDs[world.East], chunk.East.ID)
		}
		if chunkIDs[world.West] != chunk.West.ID {
			t.Fatalf("chunk West mismatch, expected: %s got: %s",
				chunkIDs[world.West], chunk.West.ID)
		}
	})

	t.Run("partially loaded chunk properly loaded", func(t *testing.T) {
		gameWorld.LoadAdjacentChunks(chunk.North)
		if chunk.North.North == nil {
			t.Fatalf("chunk North.North is not loaded")
		}
		if chunk.North.South == nil {
			t.Fatalf("chunk North.South is not loaded")
		}
		if chunk.ID != chunk.North.South.ID {
			t.Fatalf("chunk North.South mismatch, expected: %s got: %s",
				chunk.ID, chunk.North.South.ID)
		}
		if chunk.North.East == nil {
			t.Fatalf("chunk North.East is not loaded")
		}
		if chunk.North.West == nil {
			t.Fatalf("chunk North.West is not loaded")
		}
	})
}
