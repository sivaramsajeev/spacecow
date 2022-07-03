package cow

import (
	"log"

	gc "github.com/gbin/goncurses"
)

type Explosion struct {
	*gc.Window
	life int
}

func newExplosion(y, x int) *Explosion {
	w, err := gc.NewWindow(3, 3, y-1, x-1)
	if err != nil {
		log.Println("newExplosion:", err)
	}
	w.ColorOn(4)
	w.MovePrint(0, 0, `\ /`)
	w.AttrOn(gc.A_BOLD)
	w.MovePrint(1, 0, ` X `)
	w.AttrOn(gc.A_DIM)
	w.MovePrint(2, 0, `/ \`)
	return &Explosion{w, 5}
}

func (e *Explosion) Cleanup() {
	e.Delete()
}

func (e *Explosion) Collide(i int) {}

func (e *Explosion) Draw(w *gc.Window) {
	w.Overlay(e.Window)
}

func (e *Explosion) Expired(y, x int) bool {
	return e.life <= 0
}

func (e *Explosion) Update() {
	e.life--
}
