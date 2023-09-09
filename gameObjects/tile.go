package gameObjects

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct{
	Pos rl.Vector2
	Color rl.Color
	Size int32
	hitbox rl.Rectangle
}


func NewTile(x float32, y float32, color rl.Color, size int32) *Tile{
	return &Tile{
		Pos: rl.NewVector2(x,y),
		Color: color,
		Size: size,
		hitbox: rl.NewRectangle(x,y,64,64),
	}
}


func (t *Tile) Draw(){
	rl.DrawRectangle(int32(t.Pos.X), int32(t.Pos.Y), t.Size, t.Size, t.Color)
}
