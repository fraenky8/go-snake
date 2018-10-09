package main

import (
	"log"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	tests := [4]bool{}
	running := make(chan bool)

	go func(tests *[4]bool, running chan bool) {
		s := newSnake()
		f := newFood()
		f.x = windowWidth / 2
		f.y = windowHeight / 2

		for {
			f.update(s)

			if f.x < 0 {
				log.Println("f.x < 0")
				return
			} else if f.x >= windowWidth {
				log.Printf("f.x >= windowWidth: %v", f.x)
				return
			} else if f.y < 0 {
				log.Println("f.y < 0")
				return
			} else if f.y >= windowHeight {
				log.Printf("f.y >= windowHeight: %v", f.y)
				return
			}

			if f.x == 0 {
				tests[0] = true
			} else if f.x == windowWidth-cellSize {
				tests[1] = true
			} else if f.y == 0 {
				tests[2] = true
			} else if f.y == windowHeight-cellSize {
				tests[3] = true
			}

			if tests[0] && tests[1] && tests[2] && tests[3] {
				return
			}

			running <- true
		}
	}(&tests, running)

	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-running:
			// nop
			log.Printf("running: %v", tests)
		case <-timeout:
			log.Printf("evaluating: %v", tests)
			for n, b := range tests {
				if !b {
					t.Fatalf("Test %v failed!", n)
				}
			}
			return
		}
	}
}
