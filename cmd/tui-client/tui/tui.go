// Package tui contains TUI definition
package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dialytica/enjoku/cmd/tui-client/scene"
	"github.com/dialytica/enjoku/world"
)

// KeyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit},                // second column
	}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc", "quit"),
	),
}

type TUIModel struct {
	playerName  string
	playerID    string
	worldID     string
	lastPPos    *world.PlayerPosition
	keys        KeyMap
	help        help.Model
	borderStyle lipgloss.Style
	inputStyle  lipgloss.Style
	lastKey     string
	gameScene   scene.Scene
	player      *world.Player
	activeChunk *world.ChunkGraph
	gameWorld   *world.World
	quitting    bool
}

func NewTUIModel(playerName, playerID, worldID string) *TUIModel {
	tuiModel := TUIModel{
		playerName: playerName,
		playerID:   playerID,
		worldID:    worldID,
		keys:       Keys,
		help:       help.New(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
		gameScene:  scene.New(world.ChunkLenght+1, world.ChunkWidth+1),
		borderStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("61")).
			Width(world.ChunkLenght + 1).
			Height(world.ChunkWidth + 1),
	}

	return &tuiModel
}

func (m *TUIModel) Init() tea.Cmd {
	m.gameWorld = world.LoadWorld(m.playerName, m.playerID, m.worldID)
	m.player = m.gameWorld.Players[m.playerID]
	m.activeChunk = m.gameWorld.Chunks[m.player.ChunkID]
	return nil
}

func (m *TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.KeyMsgHandler(msg)
	case tea.WindowSizeMsg:
		return m.WindowSizeMsgHandler(msg)
	}

	return m, nil
}

func (m *TUIModel) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	var status string
	if m.lastKey == "" {
		status = "Waiting for input..."
	} else {
		status = "You chose: " + m.inputStyle.Render(m.lastKey)
		status = status + "\n"
		status = status + "chunk Pos: " + fmt.Sprintf("%+v", *m.activeChunk.Position) + "\n"
		status = status + "player Pos: " + fmt.Sprintf("%+v", *m.player.Position)
	}

	playerSprite := lipgloss.NewStyle().Foreground(lipgloss.Color("#75FFB7")).Render("@")
	enemySprite := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")).Render("X")

	for pos, playerID := range m.activeChunk.PlayerIDsPosition {
		var spriteX, spriteY int
		playerX, playerY := pos.GetPosition()
		if m.lastPPos == nil {
			m.lastPPos = &world.PlayerPosition{}
			m.lastPPos.SetPosition(playerX, playerY)
		}
		spriteX, spriteY = m.lastPPos.GetPosition()
		m.gameScene.RemoveSprite(spriteX+32, -spriteY+12)
		if playerID == m.playerID {
			m.gameScene.UpdateSprite(playerSprite, playerX+32, -playerY+12)
		} else {
			m.gameScene.UpdateSprite(enemySprite, playerX+32, -playerY+12)
		}
		m.lastPPos.SetPosition(playerX, playerY)
	}

	inBorder := m.borderStyle.MaxHeight(48).Render(m.gameScene.Render())

	helpView := m.help.View(m.keys)

	return "\n" + inBorder + "\n" + status + strings.Repeat("\n", 4) + helpView
}

func (m *TUIModel) KeyMsgHandler(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	x, y := m.player.GetPosition()
	switch {
	case key.Matches(msg, m.keys.Up):
		m.gameWorld.MovePlayer(m.playerID, x, y+1)
		m.lastKey = "↑"
	case key.Matches(msg, m.keys.Down):
		m.gameWorld.MovePlayer(m.playerID, x, y-1)
		m.lastKey = "↓"
	case key.Matches(msg, m.keys.Left):
		m.gameWorld.MovePlayer(m.playerID, x-1, y)
		m.lastKey = "←"
	case key.Matches(msg, m.keys.Right):
		m.gameWorld.MovePlayer(m.playerID, x+1, y)
		m.lastKey = "→"
	case key.Matches(msg, m.keys.Help):
		m.help.ShowAll = !m.help.ShowAll
	case key.Matches(msg, m.keys.Quit):
		m.quitting = true
		return m, tea.Quit
	}
	m.activeChunk = m.gameWorld.Chunks[m.player.ChunkID]
	return m, nil
}

func (m *TUIModel) WindowSizeMsgHandler(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.help.Width = msg.Width
	return m, nil
}
