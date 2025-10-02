/* Package world shows how world are created */
package world

const (
	North = "North"
	South = "South"

	West = "West"
	East = "East"

	ChunkLenght = 128
	ChunkWidth  = 96

	DefaultName = "NJQ"
)

type IPosition interface {
	GetPosition() (int, int)
	SetPosition(x, y int)
}

type World struct {
	Chunks  []*ChunkGraph
	Players []*Player
}

func CreateNewWorld() *World {
	chunk := CreateNewChunk(0, 0)
	player := CreateNewPlayer(DefaultName)
	player.ChunkID = &chunk.ID
	world := &World{
		Chunks:  []*ChunkGraph{chunk},
		Players: []*Player{player},
	}
	return world
}

func LoadWorld() {
}

func LoadPlayers() {
}

func LoadChunks() {
}
