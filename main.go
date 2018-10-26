package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle = "The Snake Game"

	windowHeight = 600
	windowWidth  = 800
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer window.Destroy()

	scene, err := newScene(window)
	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}

	events := make(chan sdl.Event)
	errc := scene.run(events)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}

	return nil
}
