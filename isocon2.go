package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var ( // MARK: var ███████████████████████████████

	//game
	mapvelocity           = 1
	gamelevel             = 1
	scorelevel, scoregame int
	//enemies
	enemytimer, enemytimercount, enemycount int
	enemyon, enemycreated                   bool
	enemies                                 = make([]enemy, 30)
	enemylvlhp                              = 1
	//weapons
	bullets     = make([]bullet, 30)
	bulletcount int
	//imgs
	hpimg    = rl.NewRectangle(0, 34, 16, 14)
	ship1    = rl.NewRectangle(0, 0, 14, 14)
	target1  = rl.NewRectangle(0, 16, 16, 16)
	target2  = rl.NewRectangle(16, 16, 16, 16)
	target3  = rl.NewRectangle(32, 16, 16, 16)
	target4  = rl.NewRectangle(48, 16, 16, 16)
	powerup1 = rl.NewRectangle(3, 82, 16, 16)
	powerup2 = rl.NewRectangle(27, 82, 17, 17)
	powerup3 = rl.NewRectangle(54, 83, 16, 16)
	meteor   = rl.NewRectangle(8, 109, 31, 31)
	enemy1   = rl.NewRectangle(2, 56, 17, 17)
	enemy2   = rl.NewRectangle(26, 57, 17, 17)
	enemy3   = rl.NewRectangle(51, 56, 17, 17)
	enemy4   = rl.NewRectangle(75, 56, 18, 18)
	enemy5   = rl.NewRectangle(103, 60, 16, 16)
	enemy6   = rl.NewRectangle(127, 59, 16, 16)
	enemy7   = rl.NewRectangle(151, 60, 16, 16)
	enemy8   = rl.NewRectangle(172, 59, 18, 18)
	enemy9   = rl.NewRectangle(197, 60, 17, 17)
	enemy10  = rl.NewRectangle(220, 60, 18, 18)
	enemy11  = rl.NewRectangle(243, 61, 16, 16)
	enemy12  = rl.NewRectangle(267, 61, 16, 16)
	enemy13  = rl.NewRectangle(291, 61, 16, 16)
	enemy14  = rl.NewRectangle(316, 61, 16, 16)
	enemy15  = rl.NewRectangle(338, 59, 17, 17)
	enemy16  = rl.NewRectangle(363, 62, 16, 16)
	//player
	playercollisiontimer int
	playercollisionpause bool
	player               = playerblock{}
	trailtimer           int
	attackhp             = 1
	//map
	topx, topy       int
	maxh             = 30
	tilew            = 16
	tileh            = 8
	level            = make([]blok, drawa)
	mapa, mapw, mapl int
	gridbloks        = make([]gridblok, drawa)
	// core
	optionslist                                                        = make([]option, 40)
	optionscount                                                       int
	options, paused, scanlines, pixelnoise, ghosting, blur             bool
	mouseblok                                                          int
	mousepos                                                           rl.Vector2
	centerlines, grid, debug, fadeblinkon, fadeblink2on                bool
	monw, monh                                                         int
	fps                                                                = 60
	framecount                                                         int
	imgs                                                               rl.Texture2D
	camera, camerabackg                                                rl.Camera2D
	fadeblink                                                          = float32(0.2)
	fadeblink2                                                         = float32(0.1)
	onoff1, onoff2, onoff3, onoff6, onoff10, onoff15, onoff30, onoff60 bool

	xcent, ycent, ytop, optionselect, tilesize, centerblok, drawblok, drawbloknext, draww, drawh, drawa int
)

//MARK: struct
type blok struct {
	activ                          bool
	h                              float32
	v1, v2, v3, v4, v5, v6, v7, v8 rl.Vector2
	color                          rl.Color
}
type option struct {
	name, txt                           string
	onoff, activ, switchvalue, intvalue bool
	value                               int
}
type gridblok struct {
	v1, v2, v3, v4 rl.Vector2
}
type playerblock struct {
	targeton, rotate               bool
	blocknumber, h, hp, timer, lr  int
	x, y, targetrotat, playerrotat float32
	target, playerrec              rl.Rectangle
}
type bullet struct {
	activ bool
	rec   rl.Rectangle
	hp    int
}
type enemy struct {
	activ, lastenemy, destroy, switch1, switch2, switch3 bool
	hp, score, movetype, timer, movetimer                int
	color                                                rl.Color
	img, rec                                             rl.Rectangle
	rotat                                                float32

	circv2   []rl.Vector2
	circl    []float32
	circdir  []int
	circfade []float32
}

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "GAME TITLE")
	//rl.ToggleFullscreen()
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	// MARK: load images
	imgs = rl.LoadTexture("imgs.png") // load images
	//	createimgs()

	rl.SetTargetFPS(fps)
	//rl.HideCursor()
	//	rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {
		framecount++

		mousepos = rl.GetMousePosition()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		if !paused {
			//	drawnocameraback()
			drawlayers()
		}
		if grid {
			drawgrid()
		}
		if options {
			drawoptions()
		}

		rl.EndMode2D()
		drawnocamera()

		if debug {
			drawdebug()
		}
		update()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func update() { // MARK: update

	input()
	timers()

	drawbloknext += mapw * mapvelocity
	player.blocknumber += mapw * mapvelocity
	drawbloknexth := drawbloknext / mapw
	playerh := player.blocknumber / mapw

	if drawbloknexth-playerh > (draww - 4) {
		player.blocknumber += mapw
	}

}
func drawnocameraback() { // MARK: drawnocameraback

}
func drawlayers() { // MARK: drawlayers ███████████████████████████████

	//terrain layer
	drawblok = drawbloknext
	count := 0
	for a := 0; a < len(gridbloks); a++ {

		if level[drawblok].activ {

			if level[drawblok].h == 0 { // ground level
				rl.DrawTriangle(gridbloks[a].v3, gridbloks[a].v4, gridbloks[a].v1, level[drawblok].color)
				rl.DrawTriangle(gridbloks[a].v1, gridbloks[a].v2, gridbloks[a].v3, level[drawblok].color)
				/*	if ghosting {
						changex := rFloat32(-2, 3)
						changey := rFloat32(-2, 3)
						v1g := rl.NewVector2(gridbloks[a].v1.X+changex, gridbloks[a].v1.Y+changey)
						v2g := rl.NewVector2(gridbloks[a].v2.X+changex, gridbloks[a].v2.Y+changey)
						v3g := rl.NewVector2(gridbloks[a].v3.X+changex, gridbloks[a].v3.Y+changey)
						v4g := rl.NewVector2(gridbloks[a].v4.X+changex, gridbloks[a].v4.Y+changey)

						rl.DrawTriangle(v3g, v4g, v1g, rl.Fade(level[drawblok].color, 0.4))
						rl.DrawTriangle(v1g, v2g, v3g, rl.Fade(level[drawblok].color, 0.4))
					}
				*/

			} else { // elevated blocks
				v5 := rl.NewVector2(gridbloks[a].v1.X, gridbloks[a].v1.Y-level[drawblok].h)
				v6 := rl.NewVector2(gridbloks[a].v2.X, gridbloks[a].v2.Y-level[drawblok].h)
				v7 := rl.NewVector2(gridbloks[a].v3.X, gridbloks[a].v3.Y-level[drawblok].h)
				v8 := rl.NewVector2(gridbloks[a].v4.X, gridbloks[a].v4.Y-level[drawblok].h)

				//top triangles
				rl.DrawTriangle(v7, v8, v5, level[drawblok].color)
				rl.DrawTriangle(v5, v6, v7, level[drawblok].color)
				//right vert triangles
				rl.DrawTriangle(gridbloks[a].v3, gridbloks[a].v4, v7, level[drawblok].color)
				rl.DrawTriangle(gridbloks[a].v4, v8, v7, level[drawblok].color)
				//right vert shadow
				rl.DrawTriangle(gridbloks[a].v3, gridbloks[a].v4, v7, rl.Fade(rl.Black, 0.5))
				rl.DrawTriangle(gridbloks[a].v4, v8, v7, rl.Fade(rl.Black, 0.5))
				//left vert triangles
				rl.DrawTriangle(v7, v6, gridbloks[a].v3, level[drawblok].color)
				rl.DrawTriangle(v6, gridbloks[a].v2, gridbloks[a].v3, level[drawblok].color)
				//left vert shadows
				rl.DrawTriangle(v7, v6, gridbloks[a].v3, rl.Fade(rl.Black, 0.3))
				rl.DrawTriangle(v6, gridbloks[a].v2, gridbloks[a].v3, rl.Fade(rl.Black, 0.3))
				if grid {
					//top
					rl.DrawLineV(v5, v6, rl.Magenta)
					rl.DrawLineV(v6, v7, rl.Magenta)
					rl.DrawLineV(v7, v8, rl.Magenta)
					rl.DrawLineV(v8, v5, rl.Magenta)
					//vert
					rl.DrawLineV(gridbloks[a].v3, v5, rl.Magenta)
					rl.DrawLineV(gridbloks[a].v4, v8, rl.Magenta)
					rl.DrawLineV(gridbloks[a].v2, v6, rl.Magenta)
				}

			}
		}

		drawblok -= mapw
		count++
		if count == draww {

			count = 0
			drawblok += mapw * draww
			drawblok--
		}
	}

	//bullets
	updatebullets()
	//enemies
	updateenemies()

	// score
	scorelvltxt := strconv.Itoa(scorelevel)
	rl.DrawText(scorelvltxt, int(gridbloks[drawa-1].v1.X+97), int(gridbloks[drawa-1].v1.Y-112), 80, randomcolor())
	rl.DrawText(scorelvltxt, int(gridbloks[drawa-1].v1.X+99), int(gridbloks[drawa-1].v1.Y-114), 80, rl.Black)
	rl.DrawText(scorelvltxt, int(gridbloks[drawa-1].v1.X+100), int(gridbloks[drawa-1].v1.Y-115), 80, randomyellow())

	// hp
	hpv2 := rl.NewVector2(gridbloks[drawa-1].v1.X-97, gridbloks[drawa-1].v1.Y-90)

	for a := 0; a < player.hp; a++ {
		destrec := rl.NewRectangle(hpv2.X, hpv2.Y, 48, 42)
		origin := rl.NewVector2(0, 0)

		rl.DrawTexturePro(imgs, hpimg, destrec, origin, 0.0, randomred())
		//	rl.DrawTextureRec(imgs, hpimg, hpv2, randomred())
		hpv2.X -= 50
	}

	//player layer
	drawblok = drawbloknext
	count = 0
	for a := 0; a < len(gridbloks); a++ {

		if drawblok == player.blocknumber {
			player.x = gridbloks[a].v1.X
			player.y = gridbloks[a].v1.Y - float32(player.h)
			// collision rotation & movemmbet
			if player.rotate {
				player.playerrotat += 30.0
				player.x += rFloat32(-10, 11)
				player.y += rFloat32(-10, 11)
			} else {
				player.playerrotat = 45.0
			}
			// player image rec
			destrec := rl.NewRectangle(player.x, player.y, 42, 42)
			origin := rl.NewVector2(21, 21)
			player.playerrec = rl.NewRectangle(destrec.X-18, destrec.Y-14, 32, 32)
			// player collision rec
			//		rl.DrawRectangleRec(player.playerrec, rl.Fade(brightyellow(), 0.2))
			rl.DrawTexturePro(imgs, ship1, destrec, origin, player.playerrotat, brightred())
			if ghosting {
				destrec2 := rl.NewRectangle(player.x+rFloat32(-3, 4), player.y+rFloat32(-3, 4), 42, 42)
				rl.DrawTexturePro(imgs, ship1, destrec2, origin, player.playerrotat, rl.Fade(brightred(), 0.3))
			}
			//player shadow
			destrec = rl.NewRectangle(player.x, player.y+float32(player.h-2), 42, 42)
			rl.DrawTexturePro(imgs, ship1, destrec, origin, 45.0, rl.Fade(rl.Black, 0.3))
			//crosshair
			if player.targeton {
				//	targetv2 := rl.NewVector2(player.x+100, player.y-100)
				origin := rl.NewVector2(player.target.Width, player.target.Height)
				destrec := rl.NewRectangle(player.x+288, player.y-200, player.target.Width*2, player.target.Height*2)
				if player.x > float32(monw/2) {
					destrec = rl.NewRectangle(player.x+140, player.y-100, player.target.Width*2, player.target.Height*2)
				}
				if player.y < float32(monh/2) {
					destrec = rl.NewRectangle(player.x+100, player.y-80, player.target.Width*2, player.target.Height*2)
				}
				rl.DrawTexturePro(imgs, player.target, destrec, origin, player.targetrotat, randomgreen())

				player.targetrotat += 1.0
			}
			//ship trail
			if !player.rotate {
				if trailtimer == 0 {
					rl.DrawCircle(int(player.playerrec.X), int(player.playerrec.Y)+int(player.playerrec.Height), 5, rl.Fade(rl.White, rF32(0.3, 0.9)))
				} else if trailtimer == 1 {
					rl.DrawCircle(int(player.playerrec.X-5), int(player.playerrec.Y+5)+int(player.playerrec.Height), 4, rl.Fade(rl.White, rF32(0.3, 0.9)))
				} else if trailtimer == 2 {
					rl.DrawCircle(int(player.playerrec.X-10), int(player.playerrec.Y+10)+int(player.playerrec.Height), 3, rl.Fade(rl.White, rF32(0.3, 0.9)))
				}
			}
		}

		drawblok -= mapw
		count++
		if count == draww {
			count = 0
			drawblok += mapw * draww
			drawblok--
		}
	}

	// grid
	if grid {
		for a := 0; a < len(gridbloks); a++ {
			rl.DrawLineV(gridbloks[a].v1, gridbloks[a].v2, rl.Magenta)
			rl.DrawLineV(gridbloks[a].v2, gridbloks[a].v3, rl.Magenta)
			rl.DrawLineV(gridbloks[a].v3, gridbloks[a].v4, rl.Magenta)
			rl.DrawLineV(gridbloks[a].v4, gridbloks[a].v1, rl.Magenta)

		}
	}

}
func drawnocamera() { // MARK: drawnocamera

	// screen fx
	if scanlines {
		for a := 0; a < monh; a++ {
			rl.DrawLine(0, a, monw, a, rl.Fade(rl.Black, 0.2))
			a += 2
		}
	}
	if pixelnoise {
		for a := 0; a < 100; a++ {
			width := rFloat32(1, 3)
			rec := rl.NewRectangle(rFloat32(0, monw), rFloat32(0, monh), width, width)
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, rF32(0.4, 1.1)))
		}
	}
	//centerlines
	if centerlines {
		rl.DrawLine(monw/2, 0, monw/2, monh, rl.Magenta)
		rl.DrawLine(0, monh/2, monw, monh/2, rl.Magenta)
	}

}
func drawoptions() { // MARK: drawoptions

	x := 100
	y := 100
	optionscount = 0
	for a := 0; a < len(optionslist); a++ {

		if a == optionselect {
			rl.DrawRectangle(x-12, y-4, monw/6, 28, darkred())
		}

		if optionslist[a].activ {
			rl.DrawText(optionslist[a].name, x, y, 20, rl.White)
			if optionslist[a].switchvalue {
				rl.DrawRectangle(x+((monw/6)-50), y-2, 24, 24, rl.White)
				rl.DrawRectangle(x+((monw/6)-50)+4, y+2, 16, 16, rl.Black)
				if optionslist[a].onoff {
					rl.DrawRectangle(x+((monw/6)-50)+8, y+6, 8, 8, rl.White)
				}
			}

			y += 30

			if y > monh-100 {
				y = 100
				x += monw / 3
			}
			optionscount++
		}

	}
	if optionselect == optionscount {
		optionselect = optionscount - 1
	}

}
func updatebullets() { // MARK: updatebullets
	//bullets
	for a := 0; a < len(bullets); a++ {
		if bullets[a].activ {
			//	rl.DrawRectangleRec(bullets[a].rec, rl.Magenta)
			v2 := rl.NewVector2(bullets[a].rec.X+4, bullets[a].rec.Y+4)
			bulletcolor := randomorange()
			rl.DrawCircleV(v2, 6, bulletcolor)
			bullets[a].rec.X += 9
			bullets[a].rec.Y -= 6

			//bullet enemy collisions
			for b := 0; b < len(enemies); b++ {
				if enemies[b].activ {
					if rl.CheckCollisionRecs(bullets[a].rec, enemies[b].rec) {
						//testcollision = true

						enemies[b].hp -= bullets[a].hp

						if enemies[b].hp <= 0 && !enemies[b].destroy {
							scorelevel += enemies[b].score
							enemies[b].destroy = true
							circv2 := make([]rl.Vector2, 40)
							circl := make([]float32, 40)
							circdir := make([]int, 40)
							circfade := make([]float32, 40)
							for c := 0; c < len(circv2); c++ {
								circv2[c] = rl.NewVector2(bullets[a].rec.X+rFloat32(-3, 4), bullets[a].rec.Y+rFloat32(-3, 4))
								circl[c] = rFloat32(7, 15)
								circfade[c] = 0.7
								for {
									circdir[c] = rInt(1, 10)
									if circdir[c] != 5 {
										break
									}
								}
							}
							enemies[b].circv2 = circv2
							enemies[b].circl = circl
							enemies[b].circdir = circdir
							enemies[b].circfade = circfade
						}

					}
				}
			}
		}
	}
}
func updateenemies() { // MARK: updateenemies

	// create next row of enemies
	if enemyon && !enemycreated {

		number := rInt(2, 6)
		x := rInt(monw/2, monw+200)
		y := rInt(-150, -80)

		for a := 0; a < number; a++ {

			createenemies()
			enemies[enemycount].rec = rl.NewRectangle(float32(x), float32(y), enemies[a].img.Width*3, enemies[a].img.Width*3)
			enemycount++

			x += rInt(80, 160)
			y += rInt(0, 150)
		}
		enemycreated = true
		// draw enemies
	} else if enemyon && enemycreated {

		for a := 0; a < len(enemies); a++ {

			if enemies[a].activ && !enemies[a].destroy {

				origin := rl.NewVector2(0, 0)
				destrec := rl.NewRectangle(enemies[a].rec.X, enemies[a].rec.Y, enemies[a].img.Width*3, enemies[a].img.Width*3)
				//enemy img
				if enemies[a].img == enemy2 {
					rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, randomcolor())
					enemies[a].rotat += rFloat32(-3, 4)
				} else if enemies[a].img == enemy4 {
					destrec = rl.NewRectangle(enemies[a].rec.X, enemies[a].rec.Y, enemies[a].img.Width*4, enemies[a].img.Width*4)
					rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, enemies[a].color)
					enemies[a].rotat += rFloat32(1, 4)
				} else if enemies[a].img == enemy7 {
					destrec = rl.NewRectangle(enemies[a].rec.X, enemies[a].rec.Y, enemies[a].img.Width*4, enemies[a].img.Width*4)
					rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, randomyellow())
				} else if enemies[a].img == enemy16 {
					destrec = rl.NewRectangle(enemies[a].rec.X, enemies[a].rec.Y, enemies[a].img.Width*4, enemies[a].img.Width*4)
					rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, randomred())
				} else if enemies[a].img == enemy8 || enemies[a].img == enemy9 || enemies[a].img == enemy10 || enemies[a].img == enemy11 || enemies[a].img == enemy12 || enemies[a].img == enemy13 {
					destrec = rl.NewRectangle(enemies[a].rec.X, enemies[a].rec.Y, enemies[a].img.Width*4, enemies[a].img.Width*4)
					rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, enemies[a].color)
				} else {
					rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, enemies[a].color)
				}
				//enemy shadow
				destrec = rl.NewRectangle(enemies[a].rec.X, enemies[a].rec.Y+110, enemies[a].img.Width*3, enemies[a].img.Width*3)
				rl.DrawTexturePro(imgs, enemies[a].img, destrec, origin, enemies[a].rotat, rl.Fade(rl.Black, 0.3))
				//rl.DrawRectangleRec(enemies[a].rec, rl.Fade(enemies[a].color, 0.2))

				// move enemies
				switch enemies[a].movetype {
				case 9:
					if !enemies[a].switch2 {
						if !enemies[a].switch1 {
							enemies[a].rec.X -= 2
							enemies[a].rec.Y += 12
							if enemies[a].rec.Y > float32(monh) {
								enemies[a].switch1 = true
							}
						} else {
							enemies[a].rec.X -= 2
							enemies[a].rec.Y -= 12
							if enemies[a].rec.Y < 0 {
								enemies[a].switch1 = false
							}
						}

						if enemies[a].rec.Y > float32(monh/3) && enemies[a].rec.Y < float32((monh/3)*2) && enemies[a].rec.X < float32(monw-(monw/3)) {

							if rolldice()+rolldice()+rolldice() > 16 {
								enemies[a].switch2 = true
							}
						}
					} else {

						enemies[a].rec.X -= 24
					}
				case 8:
					if !enemies[a].switch1 {
						enemies[a].rec.X -= 2
						enemies[a].rec.Y += 12
						if enemies[a].rec.Y > float32(monh) {
							enemies[a].switch1 = true
						}
					} else {
						enemies[a].rec.X -= 2
						enemies[a].rec.Y -= 12
						if enemies[a].rec.Y < 0 {
							enemies[a].switch1 = false
						}
					}
				case 7:
					if !enemies[a].switch1 {
						enemies[a].rec.X += 2
						enemies[a].rec.Y += 2
						enemies[a].movetimer++
						if enemies[a].movetimer >= fps/2 {
							enemies[a].movetimer = 0
							enemies[a].switch1 = true
						}
					} else {
						enemies[a].rec.X -= 8
						enemies[a].rec.Y += 2
						enemies[a].movetimer++
						if enemies[a].movetimer >= fps/2 {
							enemies[a].movetimer = 0
							enemies[a].switch1 = false
						}
					}

				case 6:
					enemies[a].rec.X -= 4
					enemies[a].rec.Y += 4
					if rolldice() == 6 {
						enemies[a].rec.Y -= 12
					}

				case 5:
					if enemies[a].rec.X < float32(monw/2) && !enemies[a].switch1 && !enemies[a].switch2 {
						if flipcoin() {
							enemies[a].switch1 = true
						} else {
							enemies[a].switch2 = true
						}
					} else if enemies[a].switch1 {
						enemies[a].rec.X -= 12
						enemies[a].rec.Y += 18

					} else if enemies[a].switch2 {
						enemies[a].rec.X -= 12
						enemies[a].rec.Y -= 18

					} else {
						enemies[a].rec.X -= 12
						enemies[a].rec.Y += 6
					}
				case 4:
					if onoff6 {
						switch rolldice() {
						case 1:
							enemies[a].rec.X -= 6
							enemies[a].rec.Y += 6
						case 2:
							enemies[a].rec.X -= 12
							enemies[a].rec.Y += 12
						case 3:
							enemies[a].rec.X -= 12
						case 4:
							enemies[a].rec.Y += 12
						case 5:
							enemies[a].rec.Y -= 12
						case 6:
							enemies[a].rec.X -= 3
						}
					}
				case 3:
					enemies[a].rec.X -= 6
					enemies[a].rec.Y += 6
					if onoff10 {
						switch rolldice() {
						case 1, 2:
							enemies[a].rec.X -= 12
						case 3, 4:
							enemies[a].rec.Y += 12
						case 5, 6:
							enemies[a].rec.Y -= 12
						}
					}
				case 2:
					enemies[a].rec.X -= 6
					enemies[a].rec.Y += 6
					if player.y < enemies[a].rec.Y {
						enemies[a].rec.Y -= 6
					} else if player.y > enemies[a].rec.Y {
						enemies[a].rec.Y += 6
					}
				case 1:
					if enemies[a].rec.X < float32(monw/2) {
						enemies[a].rec.X -= 6
						enemies[a].rec.Y += 6
						if rolldice() == 6 {
							enemies[a].rec.X += 12
						}
						if rolldice() == 6 {
							enemies[a].rec.Y += 6
						}
					} else {
						enemies[a].rec.X -= 6
						enemies[a].rec.Y += 3
					}
				}
				if rl.CheckCollisionRecs(enemies[a].rec, player.playerrec) {
					playercollisionenemy()
				}
				// end enemies
				enemies[a].timer--
				if enemies[a].timer <= 0 && enemies[a].rec.X < 0 || enemies[a].timer <= 0 && enemies[a].rec.X > float32(monw) || enemies[a].timer <= 0 && enemies[a].rec.Y < 0 || enemies[a].timer <= 0 && enemies[a].rec.Y > float32(monh) {
					enemyon = false
					enemycreated = false
					for b := 0; b < len(enemies); b++ {
						enemies[b] = enemy{}
					}
					enemycount = 0
				}
				// draw enemy explode
			} else if enemies[a].activ && enemies[a].destroy {

				//	rl.DrawRectangleRec(enemies[a].rec, randomorange())

				for c := 0; c < len(enemies[a].circdir); c++ {
					rl.DrawCircleV(enemies[a].circv2[c], enemies[a].circl[c], rl.Fade(randomorange(), enemies[a].circfade[c]))
					if ghosting && enemies[a].circfade[c] > 0.1 {
						v2 := rl.NewVector2(enemies[a].circv2[c].X+rFloat32(-3, 4), enemies[a].circv2[c].Y+rFloat32(-3, 4))
						rl.DrawCircleV(v2, float32(enemies[a].circl[c]), rl.Fade(randomorange(), 0.3))
					}

					enemies[a].circfade[c] -= 0.05
					enemies[a].circl[c] -= 0.5

					switch enemies[a].circdir[c] {
					case 1:
						enemies[a].circv2[c].X -= 10
						enemies[a].circv2[c].Y += 10
					case 2:
						enemies[a].circv2[c].Y += 10
					case 3:
						enemies[a].circv2[c].X += 10
						enemies[a].circv2[c].Y += 10
					case 4:
						enemies[a].circv2[c].X -= 10
					case 6:
						enemies[a].circv2[c].X += 10
					case 7:
						enemies[a].circv2[c].X -= 10
						enemies[a].circv2[c].Y -= 10
					case 8:
						enemies[a].circv2[c].Y -= 10
					case 9:
						enemies[a].circv2[c].X += 10
						enemies[a].circv2[c].Y -= 10
					}
				}

				enemies[a].rec.X -= 10
				enemies[a].rec.Y += 7
				if enemies[a].lastenemy {
					if enemies[a].rec.X < 0 || enemies[a].rec.Y > float32(monh) {
						enemyon = false
						enemycreated = false
						for b := 0; b < len(enemies); b++ {
							enemies[b] = enemy{}
						}
						enemycount = 0
					}
				}
			}

		}

	}

}
func createenemies() { // MARK: updateenemies

	enemies[enemycount].activ = true

	chooseenemy()

	enemies[enemycount].timer = fps * rInt(3, 7)
	enemies[enemycount].lastenemy = true
	enemies[enemycount].hp = enemylvlhp
	enemies[enemycount].score = (enemies[enemycount].hp * 5) * 2

}
func chooseenemy() { // MARK: updateenemies

	choose := rInt(1, 17)

	switch choose {
	case 1:
		enemies[enemycount].img = enemy1
		enemies[enemycount].color = rl.White
		enemies[enemycount].movetype = rInt(1, 10)
	case 2:
		enemies[enemycount].img = enemy2
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].movetype = rInt(1, 10)
	case 3:
		enemies[enemycount].img = enemy3
		enemies[enemycount].rotat = 180.0
		enemies[enemycount].color = brightyellow()
		enemies[enemycount].movetype = rInt(1, 10)
	case 4:
		enemies[enemycount].img = enemy4
		enemies[enemycount].color = rl.Pink
		enemies[enemycount].rotat = rFloat32(0, 360)
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 5:
		enemies[enemycount].img = enemy5
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].movetype = rInt(1, 10)
	case 6:
		enemies[enemycount].img = enemy6
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].movetype = rInt(1, 10)
	case 7:
		enemies[enemycount].img = enemy7
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 8:
		enemies[enemycount].img = enemy8
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 9:
		enemies[enemycount].img = enemy9
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 10:
		enemies[enemycount].img = enemy10
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 11:
		enemies[enemycount].img = enemy11
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 12:
		enemies[enemycount].img = enemy12
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 13:
		enemies[enemycount].img = enemy13
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	case 14:
		enemies[enemycount].img = enemy14
		enemies[enemycount].color = rl.White
		enemies[enemycount].movetype = rInt(1, 10)
	case 15:
		enemies[enemycount].img = enemy15
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 180.0
		enemies[enemycount].movetype = rInt(1, 10)
	case 16:
		enemies[enemycount].img = enemy16
		enemies[enemycount].color = rl.White
		enemies[enemycount].rotat = 220.0
		enemies[enemycount].rec.Height = enemies[enemycount].img.Width * 4
		enemies[enemycount].rec.Width = enemies[enemycount].img.Width * 4
		enemies[enemycount].movetype = rInt(1, 10)
	}

}
func playercollisionenemy() { // MARK: playercollisionenemy
	if !playercollisionpause {
		player.hp--
		player.rotate = true
		playercollisionpause = true
	}

}
func createtopv2(blocknumnber int) { // MARK: createmap

	level[blocknumnber].v5 = rl.NewVector2(level[blocknumnber].v1.X, level[blocknumnber].v1.Y-level[blocknumnber].h)
	level[blocknumnber].v6 = rl.NewVector2(level[blocknumnber].v2.X, level[blocknumnber].v2.Y-level[blocknumnber].h)
	level[blocknumnber].v7 = rl.NewVector2(level[blocknumnber].v3.X, level[blocknumnber].v3.Y-level[blocknumnber].h)
	level[blocknumnber].v8 = rl.NewVector2(level[blocknumnber].v4.X, level[blocknumnber].v4.Y-level[blocknumnber].h)

}
func createmap() { // MARK: createmap

	for a := 0; a < len(level); a++ {
		level[a].activ = true
		level[a].color = randombluedark()
	}
	//left border
	block := mapw
	for a := 0; a < mapl-1; a++ {
		for b := 0; b < 4; b++ {
			level[block-b].h = rFloat32(30, 150)
		}
		block += mapw
	}
	//left contour
	block = mapw - 5
	for a := 0; a < mapl-1; a++ {
		starth := rFloat32(80, 120)
		changeblock := block
		for {
			level[changeblock].h = starth
			starth -= rFloat32(10, 80)
			changeblock--
			if starth <= 0 {
				break
			}
		}
		block += mapw
	}

	addfeatures()

}
func addfeatures() { // MARK: addfeatures

}
func creategridv2() {

	topx := float32(monw / 2)
	topyy := float32((monh - (tileh * draww)) / 2)

	xorig := topx
	yorig := topyy

	count := 0

	for a := 0; a < len(gridbloks); a++ {

		gridbloks[a].v1.X = topx
		gridbloks[a].v1.Y = topyy
		gridbloks[a].v2.X = topx - float32(tilew/2)
		gridbloks[a].v2.Y = topyy + float32(tileh/2)
		gridbloks[a].v3.X = topx
		gridbloks[a].v3.Y = topyy + float32(tileh)
		gridbloks[a].v4.X = topx + float32(tilew/2)
		gridbloks[a].v4.Y = topyy + float32(tileh/2)

		topx -= float32(tilew / 2)
		topyy += float32(tileh / 2)

		count++

		if count == draww {
			count = 0
			topx = xorig
			topyy = yorig

			topx += float32(tilew / 2)
			topyy += float32(tileh / 2)

			xorig = topx
			yorig = topyy
		}

	}

}

// MARK: core  █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
func setinitialvalues() { // MARK: setinitialvalues

	player.timer = fps * 2
	playercollisiontimer = fps * 2
	enemytimer = 1

	ghosting = true
	scanlines = true
	pixelnoise = true

	draww = 80
	drawh = draww
	drawa = draww * draww
	gridbloks = make([]gridblok, drawa)

	mapw = draww + 2
	mapl = mapw * 100
	mapa = mapw * mapl
	level = make([]blok, mapa)

	centerblok = mapa / 2
	centerblok += mapw / 2

	drawbloknext = drawa + mapw*2

	xcent = monw / 2
	ycent = monh / 2

	player.blocknumber = mapw * 10
	player.blocknumber += mapw / 2
	player.h = 100
	player.hp = 3
	player.targeton = true
	player.target = target1

	creategridv2()

	createmap()

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "setres")
	setres(0, 0)
	rl.CloseWindow()
	setinitialvalues()
	raylib()

}
func input() { // MARK: input

	if rl.IsKeyPressed(rl.KeyEscape) {
		if options {
			options = false
			paused = false
		} else {
			options = true
			paused = true
			optionselect = 0
		}

	}

	if rl.IsKeyPressed(rl.KeyPause) {
		if paused {
			paused = false
		} else {
			paused = true
		}
	}

	if grid {
		if rl.IsKeyDown(rl.KeyLeft) {
			camera.Target.X -= 20
		}
		if rl.IsKeyDown(rl.KeyRight) {
			camera.Target.X += 20
		}
		if rl.IsKeyDown(rl.KeyUp) {
			camera.Target.Y -= 20
		}
		if rl.IsKeyDown(rl.KeyDown) {
			camera.Target.Y += 20
		}

	} else {
		if rl.IsKeyPressed(rl.KeySpace) {

			bullets[bulletcount].activ = true
			bullets[bulletcount].rec = rl.NewRectangle(player.x+20, player.y-30, 16, 16)
			bullets[bulletcount].hp = attackhp
			bulletcount++

			if bulletcount == len(bullets)-1 {
				bulletcount = 0
			}
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			player.blocknumber++
			player.lr = 1
		} else if rl.IsKeyReleased(rl.KeyLeft) {
			player.lr = 2
		}
		if rl.IsKeyDown(rl.KeyRight) {
			player.blocknumber--
			player.lr = 3
		} else if rl.IsKeyReleased(rl.KeyRight) {
			player.lr = 2
		}
		if rl.IsKeyDown(rl.KeyUp) {
			player.blocknumber += mapw * 2
		}
		if rl.IsKeyDown(rl.KeyDown) {
			player.blocknumber -= mapw
		}
		/*
			if rl.IsKeyDown(rl.KeyA) {
				player.h += 10
			}
			if rl.IsKeyDown(rl.KeyZ) {
				player.h -= 10
			}
		*/
	}

	if !options {

	} else if options {
		if rl.IsKeyPressed(rl.KeyUp) {
			optionselect--
			if optionselect < 0 {
				optionselect = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			optionselect++

		}
	}

	// DEV KEYS DELETE

	if rl.IsKeyPressed(rl.KeyF1) {

	}

	// DEV KEYS DELETE

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		camera.Zoom += 1.0
		//	camera.Target.X += 350
		//	camera.Target.Y += 350
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		camera.Zoom -= 1.0
	}
	if rl.IsKeyPressed(rl.KeyKpDivide) {
		if centerlines {
			centerlines = false
		} else {
			centerlines = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debug {
			debug = false
		} else {
			debug = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if grid {
			grid = false
		} else {
			grid = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp0) {
		if grid {
			grid = false
		} else {
			grid = true
		}
	}
}
func createimgs() { // MARK: createimgs

}
func createoptions() {

	optionslist[0].activ = true
	optionslist[0].name = "scanlines"
	optionslist[0].txt = "old tv lines - low/med/high performance impact"
	optionslist[0].onoff = true
	optionslist[0].switchvalue = true

	optionslist[1].activ = true
	optionslist[1].name = "pixel noise"
	optionslist[1].txt = "random dead pixels - low/med/high performance impact"
	optionslist[1].onoff = true
	optionslist[1].switchvalue = true

	optionslist[2].activ = true
	optionslist[2].name = "ghosting"
	optionslist[2].txt = "image blur - low/med/high performance impact"
	optionslist[2].onoff = true
	optionslist[2].switchvalue = true

	optionslist[3].activ = true
	optionslist[3].name = "lighting"
	optionslist[3].txt = "image blur - low/med/high performance impact"
	optionslist[3].onoff = true
	optionslist[3].switchvalue = true

}
func drawdebug() { // MARK: DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG

	rl.DrawRectangle(monw-300, 0, 300, monh, rl.Fade(rl.Black, 0.8))
	textx := monw - 290
	textx2 := monw - 145
	texty := 10

	//	camerazoomtext := fmt.Sprintf("%g", camera.Zoom)
	//	playermovingtext := strconv.FormatBool(player.moving)

	monwtxt := strconv.Itoa(monw)
	drawbloknexttxt := strconv.Itoa(drawbloknext)
	mapwtxt := strconv.Itoa(mapw)
	mapatxt := strconv.Itoa(mapa)
	zoomtxt := fmt.Sprintf("%g", camera.Zoom)
	drawbloknexthtxt := strconv.Itoa(drawbloknext / mapw)

	rl.DrawText("monwtxt", textx, texty, 10, rl.White)
	rl.DrawText(monwtxt, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("drawbloknexttxt", textx, texty, 10, rl.White)
	rl.DrawText(drawbloknexttxt, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("mapwtxt", textx, texty, 10, rl.White)
	rl.DrawText(mapwtxt, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("mapatxt", textx, texty, 10, rl.White)
	rl.DrawText(mapatxt, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("zoomtxt", textx, texty, 10, rl.White)
	rl.DrawText(zoomtxt, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("drawbloknexthtxt", textx, texty, 10, rl.White)
	rl.DrawText(drawbloknexthtxt, textx2, texty, 10, rl.White)
	texty += 12

	// fps
	rl.DrawRectangle(monw-110, monh-110, 100, 40, rl.Black)
	rl.DrawFPS(monw-100, monh-100)

}
func timers() { // MARK: timers

	if player.rotate {
		player.timer--
		if player.timer <= 0 {
			player.rotate = false
			player.timer = fps * 2
		}
	}

	if playercollisionpause {
		playercollisiontimer--
		if playercollisiontimer <= 0 {
			playercollisionpause = false
			playercollisiontimer = fps * 2
		}

	}

	if !enemyon {
		enemytimercount++
		if enemytimercount%fps == 0 {
			enemytimer--
		}
		if enemytimer <= 0 {
			enemyon = true
			enemytimercount = 0
			enemytimer = 1
		}
	}

	if trailtimer == 0 {
		trailtimer = 1
	} else if trailtimer == 1 {
		trailtimer = 2
	} else if trailtimer == 2 {
		trailtimer = 0
	}

	if framecount%1 == 0 {
		if onoff1 {
			onoff1 = false
		} else {
			onoff1 = true
		}
	}

	if framecount%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if framecount%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if framecount%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if framecount%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if framecount%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if framecount%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}
	if framecount%60 == 0 {
		if onoff60 {
			onoff60 = false
		} else {
			onoff60 = true
		}
	}
	if fadeblinkon {
		if fadeblink > 0.2 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.6 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
		}
	}
	if onoff3 {
		if fadeblink2on {
			if fadeblink2 > 0.1 {
				fadeblink2 -= 0.01
			} else {
				fadeblink2on = false
			}
		} else {
			if fadeblink2 < 0.2 {
				fadeblink2 += 0.01
			} else {
				fadeblink2on = true
			}
		}
	}
}

func setres(w, h int) { // MARK: setres

	if w == 0 {

		monw = rl.GetMonitorWidth(0)
		monh = rl.GetMonitorHeight(0)
		camera.Zoom = 2.0
		camera.Target.X += float32(monw / 4)
		camera.Target.Y += 270

	} else {
		monw = w
		monh = h
		camera.Zoom = 1.0

	}

}

func drawgrid() { // MARK: drawgrid

}

// MARK: colors
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(0, 255)))
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// random numbers
func rF32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
