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
	"fmt"
	"image/color"
	"log"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	// g.playerBody.SetVelocityUpdateFunc(func(body *cp.Body, gravity cp.Vector, damping, dt float64) {
	// 	slog.Info("Update I have no clue")
	// })

	playerShape := cp.NewBox(g.playerBody, 32, 64, 0)
	playerShape.SetFriction(0.0)
	playerShape.SetElasticity(0.0)
	playerShape.SetCollisionType(1)

	g.physicsSpace.AddBody(g.playerBody)
	g.physicsSpace.AddShape(playerShape)

	// static ground
	staticBody := cp.NewBody(cp.INFINITY, cp.INFINITY)
	staticBody.SetPosition(cp.Vector{X: 600, Y: 600})
	groundShape := cp.NewBox(staticBody, 680, 20, 0)
	groundShape.SetFriction(0.0)
	groundShape.SetElasticity(0.0)
	groundShape.SetCollisionType(2)
	g.physicsSpace.AddBody(staticBody)
	g.physicsSpace.AddShape(groundShape)

	return nil
}

const (
	WALKING_DIRECTION_NONE  = 0
	WALKING_DIRECTION_LEFT  = -1
	WALKING_DIRECTION_RIGHT = 1
)

var walkingDirection int = WALKING_DIRECTION_NONE

func (g *Game) Update() error {
	g.physicsSpace.Step(1.0 / 60.0) // FIXME: Use ebiten's ActualTPS
	wantsToJump := inpututil.IsKeyJustPressed(ebiten.KeySpace)
	if wantsToJump {
		g.playerBody.ApplyImpulseAtLocalPoint(cp.Vector{X: 0, Y: -100}, g.playerBody.CenterOfGravity())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		walkingDirection = WALKING_DIRECTION_LEFT
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyA) && walkingDirection == WALKING_DIRECTION_LEFT {
		walkingDirection = WALKING_DIRECTION_NONE
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		walkingDirection = WALKING_DIRECTION_RIGHT
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyD) && walkingDirection == WALKING_DIRECTION_RIGHT {
		walkingDirection = WALKING_DIRECTION_NONE
	}

	WALKING_ACCELERATION := 30.0
	WALKING_SPEED_LIMIT := 100.0
	// SPEED_LIMIT := 500.0
	slog.Debug("Update I have no clue")
	g.playerBody.ApplyForceAtWorldPoint(cp.Vector{X: WALKING_ACCELERATION * 2000 * float64(walkingDirection), Y: 0}, g.playerBody.CenterOfGravity())
	vel := g.playerBody.Velocity()
	vel.Y = 0
	g.playerBody.SetVelocity(vel.Clamp(WALKING_SPEED_LIMIT).X, g.playerBody.Velocity().Y)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	RENDER_BOUNDING_BOXES := true
	RENDER_FPS := true

	screen.Fill(color.Black)
	// vector.DrawFilledRect(screen, 0, 0, 32, 32, color.White, false)

	// render the player bounding box
	if RENDER_BOUNDING_BOXES {
		g.physicsSpace.EachBody(func(body *cp.Body) {
			body.EachShape(func(s *cp.Shape) {
				bb := s.BB()
				x, y, width, height := bb.L, bb.B, bb.R-bb.L, bb.T-bb.B
				vector.StrokeRect(screen, float32(x), float32(y), float32(width), float32(height), 2, color.White, true)
			})
		})
	}
	if RENDER_FPS {
		ebitenutil.DebugPrint(screen, fmt.Sprintln("TPS:", ebiten.ActualTPS()))
	}
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
