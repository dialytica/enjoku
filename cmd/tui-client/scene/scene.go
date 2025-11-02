// Package scene render game scene into terminal
package scene

import "log"

type Scene struct {
	Width   int
	Height  int
	sprites map[int]map[int]string
}

func New(Width, Height int) Scene {
	return Scene{
		Width:   Width,
		Height:  Height,
		sprites: make(map[int]map[int]string),
	}
}

// Render the Scene, currently is using exhaustive algorithm. Might Consider
// to improve it later
func (s Scene) Render() string {
	res := ""
	for iy := range s.Height {
		for ix := range s.Width {
			if entity, ok := s.sprites[iy][ix]; ok {
				res = res + entity
				continue
			}
			res = res + " "
		}
		if iy != s.Height-1 {
			res = res + "\n"
		}
	}
	return res
}

// UpdateSprite insert or update entity in the defined position. It will
// replace the entity if the position was currently not empty
func (s Scene) UpdateSprite(sprite string, x, y int) {
	if _, ok := s.sprites[y]; !ok {
		s.sprites[y] = make(map[int]string)
	}
	s.sprites[y][x] = sprite
}

// RemoveSprite remove entity and return the removed entity. If the x,y position
// input found that it's empty, an empty string will be returned
func (s Scene) RemoveSprite(x, y int) string {
	log.Printf("sprites: %+v\n", s.sprites)
	if sprite, ok := s.sprites[y][x]; ok {
		delete(s.sprites[y], x)
		return sprite
	}
	return ""
}
