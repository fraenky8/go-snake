package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	snakeColor = 0xffffff
)

type cell struct {
	x, y int32
}

type snake struct {
	x      int32
	y      int32
	xSpeed int32
	ySpeed int32

	w int32
	h int32

	body []cell

	isMoving bool
}

func newSnake() *snake {
	return &snake{
		x:      0,
		y:      0,
		xSpeed: 1,
		ySpeed: 0,
		w:      1 * cellSize,
		h:      1 * cellSize,
		body:   []cell{{x: 0, y: 0}},
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

	for i := len(s.body) - 1; i > 0; i-- {
		s.body[i] = s.body[i-1]
	}
	s.body[0].x, s.body[0].y = s.x, s.y
}

func (s *snake) draw(surface *sdl.Surface) error {
	for i := 0; i < len(s.body); i++ {
		rect := sdl.Rect{X: s.body[i].x, Y: s.body[i].y, W: s.w, H: s.h}
		if err := surface.FillRect(&rect, snakeColor); err != nil {
			return err
		}
	}
	return nil
}

func (s *snake) move(key sdl.Keycode) {
	if s.isMoving {
		return
	}
	s.isMoving = true
	if key == sdl.K_RIGHT {
		s.moveRight()
	} else if key == sdl.K_LEFT {
		s.moveLeft()
	} else if key == sdl.K_UP {
		s.moveUp()
	} else if key == sdl.K_DOWN {
		s.moveDown()
	}
	s.isMoving = false
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
		s.body = append(s.body, cell{f.x, f.y})
		return true
	}
	return false
}

func (s *snake) isDead() bool {
	for i := 1; i < len(s.body); i++ {
		if s.body[i].x == s.body[0].x && s.body[i].y == s.body[0].y {
			return true
		}
	}
	return false
}
