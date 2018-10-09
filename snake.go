package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	snakeColor = 0xffffff
)

type snake struct {
	x      int32
	y      int32
	xSpeed int32
	ySpeed int32

	w int32
	h int32

	length int

	color uint32
}

func newSnake() *snake {
	return &snake{
		x:      0,
		y:      0,
		xSpeed: 1,
		ySpeed: 0,
		w:      1 * cellSize,
		h:      1 * cellSize,
		length: 1,
		color:  snakeColor,
	}
}

func (s *snake) update() {
	s.x += s.xSpeed * cellSize
	s.y += s.ySpeed * cellSize

	if s.x == windowWidth {
		s.x = 0
	} else if s.x < 0 {
		s.x = windowWidth - s.w
	} else if s.y == windowHeight {
		s.y = 0
	} else if s.y < 0 {
		s.y = windowHeight - s.w
	}
}

func (s *snake) draw(surface *sdl.Surface) error {
	rect := sdl.Rect{X: s.x, Y: s.y, W: s.w, H: s.h}
	return surface.FillRect(&rect, s.color)
}

func (s *snake) move(key sdl.Keycode) {
	if key == sdl.K_RIGHT {
		s.moveRight()
	} else if key == sdl.K_LEFT {
		s.moveLeft()
	} else if key == sdl.K_UP {
		s.moveUp()
	} else if key == sdl.K_DOWN {
		s.moveDown()
	}
}

func (s *snake) moveRight() {
	if s.xSpeed == 1 || s.xSpeed == -1 {
		return
	}
	s.xSpeed = 1
	s.ySpeed = 0
}

func (s *snake) moveLeft() {
	if s.xSpeed == -1 || s.xSpeed == 1 {
		return
	}
	s.xSpeed = -1
	s.ySpeed = 0
}

func (s *snake) moveUp() {
	if s.ySpeed == -1 || s.ySpeed == 1 {
		return
	}
	s.xSpeed = 0
	s.ySpeed = -1
}

func (s *snake) moveDown() {
	if s.ySpeed == 1 || s.ySpeed == -1 {
		return
	}
	s.xSpeed = 0
	s.ySpeed = 1
}

func (s *snake) eat(f *food) bool {
	if s.x == f.x && s.y == f.y {
		s.length++
		return true
	}
	return false
}

func (s *snake) isDead() bool {
	return false
}
