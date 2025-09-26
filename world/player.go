package world

type PlayerPosition struct {
	x int
	y int
}

func (p *PlayerPosition) GetPosition() (int, int) {
	return p.x, p.y
}

type Player struct {
	ID       string
	Name     string
	ChunkID  *string
	Position PlayerPosition
}

func (p *Player) GetPosition() (int, int) {
	return p.Position.x, p.Position.y
}

