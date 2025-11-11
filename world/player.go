package world

import "github.com/google/uuid"

type PlayerPosition struct {
	X int
	Y int
}

func (p *PlayerPosition) GetPosition() (int, int) {
	return p.X, p.Y
}

func (p *PlayerPosition) SetPosition(x, y int) {
	p.X = x
	p.Y = y
}

type Player struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	ChunkID  string          `json:"chunk_id"`
	Position *PlayerPosition `json:"position"`
}

func CreateNewPlayer(name string) *Player {
	playerPosition := &PlayerPosition{
		X: 0,
		Y: 0,
	}
	return &Player{
		ID:       uuid.NewString(),
		Name:     name,
		Position: playerPosition,
	}
}

func (p *Player) GetPosition() (int, int) {
	return p.Position.X, p.Position.Y
}

func (p *Player) SetPosition(x, y int) {
	p.Position.SetPosition(x, y)
}
