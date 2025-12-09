package main

import (
	"fmt"
	"math"

	"github.com/manifoldco/promptui"

	"github.com/dialytica/enjoku/world"
)

type GameClient struct {
	player         *world.Player
	activeChunk    *world.ChunkGraph
	gameWorld      *world.World
	commandPrompt  promptui.Select
	movementPrompt promptui.Select
}

func NewClient() *GameClient {
	commandPrompt := promptui.Select{
		Label: "Select Command",
		Items: []string{
			"info",
			"skip",
			"move",
			"quit",
		},
	}

	movePrompt := promptui.Select{
		Label: "Move",
		Items: []string{"up", "down", "left", "right"},
	}

	playerID := "2b4d54f5-d6c0-4346-af41-469a6794adae"
	worldID := "37ba110f-ac1c-4609-b15c-50e7c192bd04"
	gameWorld := world.LoadWorld("NJQ", playerID, worldID)
	player := gameWorld.Players[playerID]

	var activeChunk *world.ChunkGraph
	if player.ChunkID == "" {
		// TODO: this is only workaround, devise a better implementaion later
		for _, c := range gameWorld.Chunks {
			activeChunk = c
			break
		}
		player.ChunkID = activeChunk.ID
	} else {
		activeChunk = gameWorld.Chunks[player.ChunkID]
	}

	return &GameClient{
		player:         player,
		activeChunk:    activeChunk,
		gameWorld:      gameWorld,
		commandPrompt:  commandPrompt,
		movementPrompt: movePrompt,
	}
}

func (g *GameClient) Run() {
	for {
		_, command, err := g.commandPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch command {
		case "info":
			g.GetInfo()
		case "skip":
			g.MovePlayerPrompt(3)
		case "move":
			g.MovePlayerPrompt(1)
		case "quit":
			return
		default:
			fmt.Printf(" It seems %s command is unavailable\n", command)
		}
	}
}

func (g *GameClient) GetInfo() {
	x, y := g.player.GetPosition()
	fmt.Printf("player %s is at x:%d y:%d \n", g.player.Name, x, y)
	x, y = g.activeChunk.GetPosition()
	fmt.Printf("current chunk is at x:%d y:%d \n", x, y)
}

func (g *GameClient) MovePlayerPrompt(distance int) {
	_, move, err := g.movementPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	playerX, playerY := g.player.GetPosition()

	switch move {
	case "up":
		playerY += distance
	case "down":
		playerY -= distance
	case "right":
		playerX += distance
	case "left":
		playerX -= distance
	}

	if math.Abs(float64(playerX)) > float64(g.activeChunk.Length)/2 {
		g.gameWorld.LoadAdjacentChunks(g.activeChunk)
		if playerX > 0 {
			g.activeChunk = g.activeChunk.East
			playerX = -g.activeChunk.Length / 2
		} else {
			g.activeChunk = g.activeChunk.West
			playerX = g.activeChunk.Length / 2
		}
		g.player.ChunkID = g.activeChunk.ID
	}

	if math.Abs(float64(playerY)) > float64(g.activeChunk.Width)/2 {
		g.gameWorld.LoadAdjacentChunks(g.activeChunk)
		if playerY > 0 {
			g.activeChunk = g.activeChunk.North
			playerY = -g.activeChunk.Width / 2
		} else {
			g.activeChunk = g.activeChunk.South
			playerY = g.activeChunk.Width / 2
		}
		g.player.ChunkID = g.activeChunk.ID
	}

	g.player.SetPosition(playerX, playerY)
}

func main() {
	gameClient := NewClient()
	gameClient.Run()
}
