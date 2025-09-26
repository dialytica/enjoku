/* Package world shows how world are created */
package world

const (
	North = "North"
	South = "South"

	West = "West"
	East = "East"
)

type IPosition interface {
	GetPosition() (int, int)
	SetPosition(x, y int)
}

type World struct {
	Chunks  []ChunkGraph
	Players []Player
}
