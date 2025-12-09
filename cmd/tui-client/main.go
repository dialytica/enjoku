package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/dialytica/enjoku/cmd/tui-client/tui"
)

func main() {
	if os.Getenv("HELP_DEBUG") != "" {
		f, err := tea.LogToFile("debug.log", "help")
		if err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		}
		defer f.Close() // nolint:errcheck
	}
	playerName := "Hero"
	playerID := "79bc4a8a-c23a-49e9-975d-f2c6b5637060"
	worldID := "e79fb1f4-b38a-4cd0-9077-333f9a2765b9"

	if _, err := tea.NewProgram(tui.NewModel(playerName, playerID, worldID), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
