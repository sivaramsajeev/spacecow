package cow

import (
	"log"

	gc "github.com/gbin/goncurses"
)

var Cow = []string{
	`^__^`,
	`(oo)`,
	`(__)`,
}

type Snake struct {
	*gc.Window
	life      int
	direction int
}

func newSnake(y, x int) *Snake {
	w, err := gc.NewWindow(5, 7, y, x)
	if err != nil {
		log.Fatal("newSnake:", err)
	}
	for i := 0; i < len(Cow); i++ {
		w.MovePrint(i, 0, Cow[i])
	}
	return &Snake{w, 5, right}
}

func (s *Snake) Cleanup() {
	s.Delete()
}

func (s *Snake) Collide(i int) {
	for k, v := range objects {
		if k == i {
			continue
		}
		switch f := v.(type) {
		case *Food:
			fy, fx := f.YX()
			sy, sx := s.YX()
			if fy == sy && fx == sx {
				objects = append(objects, newExplosion(f.YX()))
				f.alive = false
			}
		}
	}
}

func (s *Snake) Draw(w *gc.Window) {
	w.Overlay(s.Window)
}

func (s *Snake) Expired(y, x int) bool {
	return s.life <= 0
}

func (s *Snake) Update() {
	y, x := s.YX()
	switch s.direction {
	case up:
		s.MoveWindow(y-1, x)
	case down:
		s.MoveWindow(y+1, x)
	case left:
		s.MoveWindow(y, x-1)
	case right:
		s.MoveWindow(y, x+1)

	}
	return
}
