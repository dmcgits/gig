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

}
