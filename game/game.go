package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	//"strconv"
	"fmt"
)

var (
	square *ebiten.Image
	x      float64 = 0
	y      float64 = 0
	opts           = &ebiten.DrawImageOptions{}
	step float64 = 1
)

func initStep() {
	step = 0.1
}

func update(screen *ebiten.Image) error {
	// Fill the screen with #FF0000 color
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0x00})
	// Display the text though the debug function
	//ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")
	if square == nil {
		// Create an 16x16 image
		square, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)

	}

	// Fill the square with the white color
	square.Fill(color.White)

	// Create an empty option struct

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		opts.GeoM.Translate(0, step*-1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		opts.GeoM.Translate(0, step)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		opts.GeoM.Translate(step*-1, step)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		opts.GeoM.Translate(step, 0)
	}

	fmt.Println(step)
	//dx := strconv.Itoa(square.Bounds().Dx())
	/*
	if (x < 0 || x > 320) {
		ebitenutil.DebugPrint(screen, "!")
	} else {
		ebitenutil.DebugPrint(screen, "")
	}
	*/
	//opts.GeoM.Scale(x, y)

	//opts.GeoM.Scale(x, y)

	// Draw the square image to the screen with an empty option
	screen.DrawImage(square, opts)
	return nil
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
