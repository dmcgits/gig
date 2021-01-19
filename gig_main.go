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
	bgColor := sdl.Color{R: 0, G: 155, B: 0, A: 255}
	renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A)

	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		fmt.Println("Attempted Sdl.LoadBmp:", err)
		return
	}

	playerTex, err := renderer.CreateTextureFromSurface(img)

	if err != nil {
		fmt.Println("Attempted texture from surface:", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
			// Our game loop
			renderer.Clear()

			renderer.Copy(playerTex,
				&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
				&sdl.Rect{X: 100, Y: 100, W: 105, H: 105})
			renderer.Present()
		}
	}
}
