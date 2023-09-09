package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"vimjam/gameObjects"
)

type Game struct{
	Width int32
	Height int32
	Tile string
	Player gameObjects.Player 
	Tiles []gameObjects.Tile
}

func NewGame () *Game{
	return &Game{
		Width: 1280,
		Height: 768,
		Tile: "Unnamed Platform",
	}
}

func (g *Game) CreateMap(){
	tileMap := [12]string{
		"####################",
		"#   #              #",
		"#   #              #",
		"#   #              #",
		"#                  #",
		"#                  #",
		"#  ##              #",
		"#                  #",
		"#         ##       #",
		"#                  #",
		"#     ####         #",
		"####################",
	}
	for y := 0; y < len(tileMap); y++{
		for x := 0; x < len(tileMap[y]); x++{
			if tileMap[y][x] == '#'{
				tile := gameObjects.NewTile(float32(x * 64), float32(y * 64), rl.Black, int32(64))
				g.Tiles = append(g.Tiles, *tile)
			}
		}
		
	}	
}


func (g *Game) RenderMap(){
	for _,tile := range g.Tiles{
		tile.Draw()	
	}
}

func (g *Game) HandleHorizontalCollisions(){
	for _, tile := range g.Tiles {
		g.Player.CheckHorizontalCollision(&tile)
	}
}

func (g *Game) HandleVerticalollisions(){
	for _, tile := range g.Tiles {
		g.Player.CheckVerticalCollision(&tile)
	}
}

func (g *Game) Run(){
	rl.InitWindow(g.Width, g.Height,g.Tile)	
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose(){
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		g.RenderMap()
		g.Player.Update(g.Tiles)	
		g.Player.Draw()
		rl.EndDrawing()
	}
}

func main(){
	game := NewGame()
	player := gameObjects.NewPlayer()
	game.Player = *player
	game.CreateMap()
	game.Run()
}

