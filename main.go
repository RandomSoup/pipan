package main

import (
	"fmt"
	"runtime"
  "sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/luminoso-256/pipan/libmcpi"
)

const (
	VERSION = "21MMDD"
)

var (
	tabs        = []string{"Play", "Profiles", "About"}
	currTab     = 0
	font        rl.Font
	profiles    []PipanProfile
	cSelProfile = 0
)

var wg *sync.WaitGroup

func main() {
  /* Wait until the command is done. */
  defer wg.Wait()
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "Pipan")
	rl.SetTargetFPS(60)
	font = rl.LoadFont("data/font.ttf")
	profiles = append(profiles, PipanProfile{
		"A Test Profile", true, "Tiny", "AUsername", make(map[string]bool),
	})
	profiles = append(profiles, PipanProfile{
		"Test 2: Electric Boogalo", false, "Tiny", "AUsername", make(map[string]bool),
	})
	profiles = append(profiles, PipanProfile{
		"Three's a crowd", true, "Tiny", "AUsername", make(map[string]bool),
	})
	for !rl.WindowShouldClose() {

		//= Core Input
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyOne) {
			currTab = 0
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyTwo) {
			currTab = 1
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyThree) {
			currTab = 2
		}

		if currTab == 0 {
			if rl.IsKeyDown(rl.KeyUp) && cSelProfile > 0 {
				cSelProfile -= 1
			}
			if rl.IsKeyDown(rl.KeyDown) && cSelProfile < len(profiles)-1 {
				cSelProfile += 1
			}
			if rl.IsKeyDown(rl.KeyEnter) {
				var ff []string
				for f, b := range profiles[cSelProfile].Features {
					if b {
						ff = append(ff, f)
					}
				}
				lp := libmcpi.LaunchProfile{
					ff,
					profiles[cSelProfile].Username,
					profiles[cSelProfile].RendDist,
					"minecraft-pi-reborn-client",
				} 
				var wg = lp.Launch()
        break
			}
		}

		//= Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		/* The three core pages of UI */
		switch currTab {
		case 0:
			pfX := float32(5)
			for i, profile := range profiles {
				name := ""
				if profile.Modded {
					name = fmt.Sprintf("[%d] %s (modded)", i, profile.Name)
				} else {
					name = fmt.Sprintf("[%d] %s (vanilla)", i, profile.Name)
				}
				if i == cSelProfile {
					rl.DrawTextEx(font, name, rl.Vector2{5, pfX}, 16, 3, rl.Orange)
				} else {
					rl.DrawTextEx(font, name, rl.Vector2{5, pfX}, 16, 3, rl.Red)
				}
				pfX += 18
			}
			rl.DrawTextEx(font, "Arrows to Sel | Enter to Launch", rl.Vector2{5, float32(rl.GetScreenHeight()) - 36}, 14, 3, rl.LightGray)
			break
		case 1:
			break
		case 2:
			if runtime.GOOS == "windows" {
				//if you're building on windows this isnt going to work. Such a build should appropriately be identified as a UI dummy
				rl.DrawTextEx(font, "Pipan ~ "+VERSION+" ("+runtime.GOOS+"-"+runtime.GOARCH+") [UIDummy]", rl.Vector2{5, 5}, 21, 3, rl.Black)
			} else {
				rl.DrawTextEx(font, "Pipan ~ "+VERSION+" ("+runtime.GOOS+"-"+runtime.GOARCH+")", rl.Vector2{5, 5}, 21, 3, rl.Black)
			}
			rl.DrawTextEx(font, "-----------------------------\nhttps://github.com/randomsoup/pipan", rl.Vector2{5, 30}, 15, 3, rl.Black)
			rl.DrawTextEx(font, "Pipan is (C) Luminoso 2021 / MIT \nHave suggestions? Open a Github Issue", rl.Vector2{5, 85}, 15, 3, rl.Black)
			break
		}

		/* Tab Bar */
		rl.DrawLine(0, int32(rl.GetScreenHeight()-20), int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()-20), rl.LightGray)
		tbX := float32(0)
		for i, tab := range tabs {
			if i == currTab {
				rl.DrawRectangleRounded(rl.Rectangle{tbX + 5, float32(rl.GetScreenHeight() - 18), 20, 16}, 1, 12, rl.Red)
			} else {
				rl.DrawRectangleRounded(rl.Rectangle{tbX + 5, float32(rl.GetScreenHeight() - 18), 20, 16}, 1, 12, rl.Maroon)
			}
			rl.DrawTextEx(font, fmt.Sprintf("^%d", i+1), rl.Vector2{tbX + 6, float32(rl.GetScreenHeight() - 16)}, 14, 3, rl.Black)

			rl.DrawTextEx(font, tab, rl.Vector2{tbX + 30, float32(rl.GetScreenHeight() - 16)}, 14, 3, rl.Black)

			tbX += rl.MeasureTextEx(font, tab, 14, 3).X + 32
		}

		rl.EndDrawing()

	}
}
