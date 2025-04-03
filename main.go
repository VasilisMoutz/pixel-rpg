package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

type Enemy struct {
	*Sprite
	followsPlayer bool
}

type Player struct {
	*Sprite
	health uint
}

type Potion struct {
	*Sprite
	healthAmt uint
}

type Game struct {
	player  *Player
	enemies *[]Enemy
	potions *[]Potion
}

func (g *Game) Update() error {

	// Player moving functionality
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.X -= 2
	}

	// Enemies functionality
	for _, enemy := range *g.enemies {
		if enemy.followsPlayer {

			if g.player.X > enemy.X {
				enemy.X += 1
			} else if g.player.X < enemy.X {
				enemy.X -= 1
			}

			if g.player.Y > enemy.Y {
				enemy.Y += 1
			} else if g.player.Y < enemy.Y {
				enemy.Y -= 1
			}
		}
	}

	// Potion functionality
	for _, potion := range *g.potions {
		if g.player.X > potion.X {
			g.player.health += potion.healthAmt
			fmt.Printf("Picked up potion! Health: %d\n", g.player.health)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Screen Color
	screen.Fill(color.RGBA{100, 228, 245, 7})

	// settting for placing
	opts := &ebiten.DrawImageOptions{}

	// Make player
	opts.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(
		g.player.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image),
		opts,
	)
	opts.GeoM.Reset()

	// Make enemies
	for _, enemy := range *g.enemies {
		opts.GeoM.Translate(enemy.X, enemy.Y)
		screen.DrawImage(
			enemy.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image),
			opts,
		)
		opts.GeoM.Reset()
	}

	// Make potions
	for _, potion := range *g.potions {
		opts.GeoM.Translate(potion.X, potion.Y)
		screen.DrawImage(
			potion.Image,
			opts,
		)
		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	// window dimentions
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// get ninja image
	ninja, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	// get squirrel image
	squirrel, _, err := ebitenutil.NewImageFromFile("assets/images/squirrel.png")
	if err != nil {
		log.Fatal(err)
	}

	// get mouse image
	mouse, _, err := ebitenutil.NewImageFromFile("assets/images/mouse.png")
	if err != nil {
		log.Fatal(err)
	}

	// Health potion
	potion, _, err := ebitenutil.NewImageFromFile("assets/images/life-pot.png")
	if err != nil {
		log.Fatal(err)
	}

	// construct
	game := &Game{
		player: &Player{
			&Sprite{
				Image: ninja,
				X:     50,
				Y:     50,
			},
			50,
		},
		enemies: &[]Enemy{
			{
				&Sprite{
					Image: squirrel,
					X:     100,
					Y:     150,
				},
				false,
			},
			{
				&Sprite{
					Image: mouse,
					X:     200,
					Y:     200,
				},
				true,
			},
		},
		potions: &[]Potion{
			{
				&Sprite{
					Image: potion,
					X:     220,
					Y:     170,
				},
				5,
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
