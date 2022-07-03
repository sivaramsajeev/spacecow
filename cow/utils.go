package cow

import (
	"log"
	"math/rand"

	gc "github.com/gbin/goncurses"
)

func genStarfield(pl, pc int) *gc.Pad {
	pad, err := gc.NewPad(pl, pc)
	if err != nil {
		log.Fatal(err)
	}
	stars := int(float64(pc*pl) * density)
	planets := int(float64(pc*pl) * planet_density)
	for i := 0; i < stars; i++ {
		y, x := rand.Intn(pl), rand.Intn(pc)
		c := int16(rand.Intn(4) + 1)
		pad.AttrOn(gc.A_BOLD | gc.ColorPair(c))
		pad.MovePrint(y, x, ".")
		pad.AttrOff(gc.A_BOLD | gc.ColorPair(c))
	}
	for i := 0; i < planets; i++ {
		y, x := rand.Intn(pl), rand.Intn(pc)
		c := int16(rand.Intn(2) + 5)
		pad.ColorOn(c)
		if i%2 == 0 {
			pad.MoveAddChar(y, x, 'O')
		}
		pad.MoveAddChar(y, x, 'o')
		pad.ColorOff(c)
	}
	return pad
}

func lifeToText(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "*"
	}
	return s
}

func drawObjects(s *gc.Window) {
	for _, ob := range objects {
		ob.Draw(s)
	}
}

func handleInput(stdscr *gc.Window, snake *Snake) bool {
	lines, cols := stdscr.MaxYX()
	y, x := snake.YX()
	k := stdscr.GetChar()

	switch byte(k) {
	case 0:
		break
	case 'a':
		x--
		if x < 2 {
			x = 2
		}
		snake.direction = left
	case 'd':
		x++
		if x > cols-3 {
			x = cols - 3
		}
		snake.direction = right
	case 's':
		y++
		if y > lines-4 {
			y = lines - 4
		}
		snake.direction = down
	case 'w':
		y--
		if y < 2 {
			y = 2
		}
		snake.direction = up
	default:
		return false
	}
	snake.MoveWindow(y, x)
	return true
}

func updateObjects(my, mx int) {
	end := len(objects)
	tmp := make([]Object, 0, end)
	for _, ob := range objects {
		ob.Update()
	}
	for i, ob := range objects {
		ob.Collide(i)
		if ob.Expired(my, mx) {
			ob.Cleanup()
		} else {
			tmp = append(tmp, ob)
		}
	}
	if len(objects) > end {
		objects = append(tmp, objects[end:]...)
	} else {
		objects = tmp
	}
}

var speeds = []int{-75, -50, -25, -10, 0, 10, 25, 50, 75}

func spawnFood(my, mx int) {
	var y, x, sy, sx int
	switch rand.Intn(4) {
	case 0:
		y, x = 1, rand.Intn(mx-2)+1
		sy, sx = speeds[5:][rand.Intn(4)], speeds[rand.Intn(9)]
	case 1:
		y, x = rand.Intn(my-2)+1, 1
		sy, sx = speeds[rand.Intn(9)], speeds[5:][rand.Intn(4)]
	case 2:
		y, x = rand.Intn(my-2)+1, mx-2
		sy, sx = speeds[rand.Intn(9)], speeds[rand.Intn(4)]
	case 3:
		y, x = my-2, rand.Intn(mx-2)+1
		sy, sx = speeds[rand.Intn(4)], speeds[rand.Intn(9)]
	}
	w, err := gc.NewWindow(1, 1, y, x)
	if err != nil {
		log.Println("spawnFood:", err)
	}
	a := &Food{Window: w, alive: true, sy: sy, sx: sx, y: y * 100,
		x: x * 100}
	a.ColorOn(2)
	a.Print("@")
	objects = append(objects, a)
}
