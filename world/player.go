package world

import "github.com/google/uuid"

type PlayerPosition struct {
	x int
	y int
}

func (p *PlayerPosition) GetPosition() (int, int) {
	return p.x, p.y
}

func (p *PlayerPosition) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

type Player struct {
	ID       string
	Name     string
	ChunkID  string
	Position *PlayerPosition
}

func CreateNewPlayer(name string) *Player {
	playerPosition := &PlayerPosition{
		x: 0,
		y: 0,
	}
	return &Player{
		ID:       uuid.NewString(),
		Name:     name,
		Position: playerPosition,
	}
}

func (p *Player) GetPosition() (int, int) {
	return p.Position.x, p.Position.y
}

func (p *Player) SetPosition(x, y int) {
	p.Position.SetPosition(x, y)
}
