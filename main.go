package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PlayerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.X -= 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 228, 245, 7})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.X, g.Y)

	screen.DrawImage(
		g.PlayerImage.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image),
		op,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	img, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{PlayerImage: img, X: 100, Y: 100}); err != nil {
		log.Fatal(err)
	}
}
