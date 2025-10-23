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
	gameWorld := world.CreateNewWorld()

	commandPrompt := promptui.Select{
		Label: "Select Command",
		Items: []string{
			"info",
			"move",
			"quit",
		},
	}

	movePrompt := promptui.Select{
		Label: "Move at 1 step",
		Items: []string{"up", "down", "left", "right"},
	}

	return &GameClient{
		player:         gameWorld.Players[0],
		activeChunk:    gameWorld.Chunks["0:0"],
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
		case "move":
			g.MovePlayerPrompt()
		case "quit":
			return
		default:
			fmt.Printf(" It seems %s command is unavailable\n", command)
		}
	}
}

func (g *GameClient) GetInfo() {
	x, y := g.player.GetPosition()
	fmt.Printf("player is at x:%d y:%d \n", x, y)
	x, y = g.activeChunk.GetPosition()
	fmt.Printf("current chunk is at x:%d y:%d \n", x, y)
}

func (g *GameClient) MovePlayerPrompt() {
	_, move, err := g.movementPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	playerX, playerY := g.player.GetPosition()

	switch move {
	case "up":
		playerY++
	case "down":
		playerY--
	case "right":
		playerX++
	case "left":
		playerX--
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
		g.player.ChunkID = &g.activeChunk.ID
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
		g.player.ChunkID = &g.activeChunk.ID
	}

	g.player.SetPosition(playerX, playerY)
}

func main() {
	gameClient := NewClient()
	gameClient.Run()
}
