// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

type Game struct {
	physicsSpace *cp.Space
	playerBody   *cp.Body
}

func (g *Game) Initialize() error {
	g.physicsSpace = cp.NewSpace()
	g.physicsSpace.Iterations = 1
	g.physicsSpace.SetGravity(cp.Vector{X: 0, Y: 100})
	g.playerBody = cp.NewBody(1, cp.INFINITY)
	g.playerBody.SetPosition(cp.Vector{X: 600, Y: 300})
	playerShape := cp.NewBox(g.playerBody, 32, 64, 0)
	playerShape.SetFriction(0.0)
	playerShape.SetElasticity(0.0)
	playerShape.SetCollisionType(1)

	g.physicsSpace.AddBody(g.playerBody)
	g.physicsSpace.AddShape(playerShape)
	return nil
}

func (g *Game) Update() error {
	g.physicsSpace.Step(1.0 / 60.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	// vector.DrawFilledRect(screen, 0, 0, 32, 32, color.White, false)

	// render the player bounding box
	g.playerBody.EachShape(func(s *cp.Shape) {
		bb := s.BB()
		x, y, width, height := bb.L, bb.B, bb.R-bb.L, bb.T-bb.B
		vector.StrokeRect(screen, float32(x), float32(y), float32(width), float32(height), 2, color.White, true)
	})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten")
	game := Game{}
	game.Initialize()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
