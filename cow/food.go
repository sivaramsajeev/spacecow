package cow

import (
	gc "github.com/gbin/goncurses"
)

type Food struct {
	*gc.Window
	alive  bool
	y, x   int
	sy, sx int
}

func (a *Food) Cleanup() {
	a.Delete()
}

func (a *Food) Collide(i int) {
}

func (a *Food) Draw(w *gc.Window) {
	w.Overlay(a.Window)
}

func (a *Food) Expired(my, mx int) bool {
	y, x := a.YX()
	if x <= 0 || x >= mx-1 || y <= 0 || y >= my-1 || !a.alive {
		return true
	}
	return false
}

func (a *Food) Update() {
	a.y += a.sy
	a.x += a.sx
	a.MoveWindow(a.y/100, a.x/100)
}
