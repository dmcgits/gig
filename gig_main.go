//--------------------------------------------------------------------------------------
// Gaming In Go Tutorial
//
// Following along with this series: https://www.youtube.com/watch?v=5HCdqV4nQkQ
//
// Use the composition/component pattern and the SDL2 libraries to make a basic game.
//
//--------------------------------------------------------------------------------------
package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	// SDL, an external simple media library, needs initialization.
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Attempting sdl.INIT_EVERYTHING: ", err)
		return
	}
	// Ask SDL to create and show us a window
	window, err := sdl.CreateWindow(
		"GIG follow along!",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, // location
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Attempted sdl.CreateWindow: ", err)
		return
	}
	defer window.Destroy() // when main ends, clean up the window

	// make a renderer to draw our sprites to screen using hardware acceleration
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Attempting sdl.createRenderer: ", err)
		return
	}
	defer renderer.Destroy() // SDL uses c and needs plenty of memory cleanup
	renderer.SetDrawColor(0, 155, 0, 255)
	for {
		// Our game loop
		renderer.Clear()
		renderer.Present()
	}
}
