package gameObjects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PLAYERSIZE = 64 

type Player struct{
	pos rl.Vector2
	direction rl.Vector2
	velocity rl.Vector2
	moveSpeed float32
	color rl.Color
	hitbox rl.Rectangle
	gravity float32
}

func NewPlayer() *Player{
	return &Player{
		pos: rl.NewVector2(500,400),
		direction: rl.NewVector2(0,0),
		velocity: rl.NewVector2(0,0),
		moveSpeed: 4,
		color: rl.Red,
		hitbox: rl.NewRectangle(0,0,PLAYERSIZE,PLAYERSIZE),
		gravity: 0.32,
	}
}

func (p *Player) HandleInput(){
	p.direction.X = 0 
	p.direction.Y = 0 
	
	if(rl.IsKeyDown(rl.KeyA)){
		p.direction.X = -1
	}
	if(rl.IsKeyDown(rl.KeyD)){
		p.direction.X = 1
	}
	if(rl.IsKeyDown(rl.KeyW)){
		p.direction.Y = -1
	}
	if(rl.IsKeyDown(rl.KeyS)){
		p.direction.Y = 1
	}
}

func (p *Player) ApplyGravity(){
	p.pos.Y += p.gravity
}

func (p *Player) Move(tiles []Tile){	
	p.pos.X += p.direction.X * p.moveSpeed
	for _, tile := range tiles {
		p.CheckHorizontalCollision(&tile)
	}

	p.pos.Y += p.direction.Y * 2	
	for _, tile := range tiles {
		p.CheckVerticalCollision(&tile)
	}
}

func (p *Player) Update(tiles []Tile){
	p.HandleInput()
	p.Move(tiles)
}

func (p *Player) Draw(){
	rl.DrawRectangle(int32(p.pos.X),int32(p.pos.Y),PLAYERSIZE,PLAYERSIZE,p.color)
}

func (p *Player) CheckHorizontalCollision(tile *Tile){
	p.hitbox.X = p.pos.X
	p.hitbox.Y = p.pos.Y
	tile.hitbox.X = tile.Pos.X
	tile.hitbox.Y = tile.Pos.Y

	if(rl.CheckCollisionRecs(p.hitbox,tile.hitbox)){
		if(p.direction.X < 0){
			p.pos.X = tile.Pos.X + 64
			
		}
		if(p.direction.X > 0){
			p.pos.X = tile.Pos.X - 64
		}
	}
}

func (p *Player) CheckVerticalCollision(tile *Tile){
	p.hitbox.X = p.pos.X
	p.hitbox.Y = p.pos.Y
	tile.hitbox.X = tile.Pos.X
	tile.hitbox.Y = tile.Pos.Y

	if(rl.CheckCollisionRecs(p.hitbox,tile.hitbox)){
		if(p.direction.Y < 0){
			p.pos.Y = tile.Pos.Y + 64
		}
		if(p.direction.Y > 0){
			p.pos.Y = tile.Pos.Y - 64
		}
	}
}
