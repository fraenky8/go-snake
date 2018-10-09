package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	foodColor = 0xfe2e2e
)

type food struct {
	x int32
	y int32

	w int32
	h int32

	color uint32
}

func newFood() *food {

	rand.Seed(time.Now().UTC().UnixNano())

	f := &food{
		x:     0,
		y:     0,
		w:     1 * cellSize,
		h:     1 * cellSize,
		color: foodColor,
	}

	f.update(nil)

	return f
}

func (f *food) update(snake *snake) {

	// FIXME food can appear outside of area!!

	col := windowWidth / cellSize
	row := windowHeight / cellSize

	f.x = int32(random(1, col+1)) * cellSize
	f.y = int32(random(1, row+1)) * cellSize

	if snake != nil {
		// FIXME exclude snake's location (incl length/tail)
	}
}

func (f *food) draw(surface *sdl.Surface) error {
	rect := sdl.Rect{X: f.x, Y: f.y, W: f.w, H: f.h}
	return surface.FillRect(&rect, f.color)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
