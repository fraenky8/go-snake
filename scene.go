package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	backgroundColor = 0x24292e

	cellSize = 20
)

type scene struct {
	window  *sdl.Window
	surface *sdl.Surface

	snake *snake
	food  *food
}

func newScene(window *sdl.Window) (*scene, error) {

	surface, err := window.GetSurface()
	if err != nil {
		return nil, fmt.Errorf("could not create surface: %v", err)
	}

	snake := newSnake()
	food := newFood()

	return &scene{
		window:  window,
		surface: surface,
		snake:   snake,
		food:    food,
	}, nil
}

func (s *scene) run(events <-chan sdl.Event) <-chan error {
	errc := make(chan error)

	// TODO smoothen the rendering

	go func() {
		defer close(errc)
		tick := time.Tick(65 * time.Millisecond)
		for {
			select {
			case e := <-events:
				if done := s.handleEvent(e); done {
					return
				}
			case <-tick:
				s.update()

				if s.snake.isDead() {
					fmt.Println("GAME OVER!")
					return
				}

				if err := s.draw(); err != nil {
					errc <- err
				}

				if s.snake.eat(s.food) {
					s.food.update(s.snake)
				}
			}
		}
	}()

	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		if t.Type == sdl.KEYDOWN && t.State == sdl.PRESSED {
			s.snake.move(t.Keysym.Sym)
		}
	default:
	}
	return false
}

func (s *scene) update() {
	s.snake.update()
}

func (s *scene) draw() error {
	if err := s.surface.FillRect(nil, backgroundColor); err != nil {
		return fmt.Errorf("could not draw scene: %v", err)
	}
	if err := s.food.draw(s.surface); err != nil {
		return fmt.Errorf("could not draw food: %v", err)
	}
	if err := s.snake.draw(s.surface); err != nil {
		return fmt.Errorf("could not draw snake: %v", err)
	}
	return s.window.UpdateSurface()
}
