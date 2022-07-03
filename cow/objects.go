package cow

import (
	"log"
	"math/rand"
	"os"
	"time"

	gc "github.com/gbin/goncurses"
)

type Object interface {
	Cleanup()
	Collide(int)
	Draw(*gc.Window)
	Expired(int, int) bool
	Update()
}

var objects = make([]Object, 0, 16)

const density = 0.05
const planet_density = 0.001

const (
	up int = iota
	down
	left
	right
)

func Start() {
	f, err := os.Create("err.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	var stdscr *gc.Window
	stdscr, err = gc.Init()
	if err != nil {
		log.Println("Init:", err)
	}
	defer gc.End()

	rand.Seed(time.Now().Unix())
	gc.StartColor()
	gc.Cursor(0)
	gc.Echo(false)
	gc.HalfDelay(1)

	gc.InitPair(1, gc.C_WHITE, gc.C_BLACK)
	gc.InitPair(2, gc.C_YELLOW, gc.C_BLACK)
	gc.InitPair(3, gc.C_MAGENTA, gc.C_BLACK)
	gc.InitPair(4, gc.C_RED, gc.C_BLACK)

	gc.InitPair(5, gc.C_BLUE, gc.C_BLACK)
	gc.InitPair(6, gc.C_GREEN, gc.C_BLACK)

	lines, cols := stdscr.MaxYX()
	pl, pc := lines, cols*3

	snake := newSnake(lines/2, 5)
	objects = append(objects, snake)

	field := genStarfield(pl, pc)
	text := stdscr.Duplicate()

	c := time.NewTicker(time.Second / 2)
	c2 := time.NewTicker(time.Second / 16)
	px := 0

loop:
	for {
		text.MovePrintf(0, 0, "Life: [%-5s]", lifeToText(snake.life))
		stdscr.Erase()
		stdscr.Copy(field.Window, 0, px, 0, 0, lines-1, cols-1, true)
		drawObjects(stdscr)
		stdscr.Overlay(text)
		stdscr.Refresh()
		select {
		case <-c.C:
			spawnFood(stdscr.MaxYX())
			if px+cols >= pc {
				break loop
			}
			px++
		case <-c2.C:
			updateObjects(stdscr.MaxYX())
			drawObjects(stdscr)
		default:
			if !handleInput(stdscr, snake) || snake.Expired(-1, -1) {
				break loop
			}
		}
	}
	msg := "GAME OVER!!!"
	end, err := gc.NewWindow(5, len(msg)+4, (lines/2)-2, (cols-len(msg))/2)
	if err != nil {
		log.Fatal("game over:", err)
	}
	end.MovePrint(2, 2, msg)
	end.Box(gc.ACS_VLINE, gc.ACS_HLINE)
	end.Refresh()
	gc.Nap(2000)
}
