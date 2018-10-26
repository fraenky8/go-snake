package main

import (
	"fmt"
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

func (f *food) String() string {
	return fmt.Sprintf("Food{x:%v, y:%v}", f.x, f.y)
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

	col := windowWidth / cellSize
	row := windowHeight / cellSize

	f.x = int32(rand.Intn(col)) * cellSize
	f.y = int32(rand.Intn(row)) * cellSize

	// is nil at the very beginning
	if snake != nil {
		for _, cell := range snake.body {
			if f.x == cell.x && f.y == cell.y {
				f.update(snake)
			}
		}
	}
}

func (f *food) draw(surface *sdl.Surface) error {
	rect := sdl.Rect{X: f.x, Y: f.y, W: f.w, H: f.h}
	return surface.FillRect(&rect, f.color)
}
