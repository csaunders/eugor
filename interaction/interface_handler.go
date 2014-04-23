package interaction

import (
	"eugor/dungeon"
	"eugor/sprite"
	"eugor/userinterface"
	"github.com/nsf/termbox-go"
)

type InterfaceHandler struct {
	widgets  map[rune]userinterface.Widget
	handlers map[userInterface.Widget]func(rune, dungeon.TileMap, sprites.Sprite)
}

func (i *InterfaceHandler) Handle(e termbox.Event) (rune, func(rune, dungeon.TileMap, sprites.Sprite)) {
	w := i.widgets[e.Ch]
	if w != nil {
		return e.Ch, i.generalizedHandler(w)
	}
}

func (i *InterfaceHandler) generalizedHandler(w userinterface.Widget) func(rune, dungeon.TileMap, sprites.Sprite) {
	var h func(rune, dungeon.TileMap, sprites.Sprite) = i.handlers[w]
	if h == nil {
		h = func(r rune, d dungeon.TileMap, s sprites.Sprite) {
			widgetHandler = w.Handler()
			if widgetHandler != nil && widgetHandler.CanHandle(r) {
				widgetHandler.Handle(r)(r, d, s)
			} else {
				w.Toggle()
			}
		}
		i.handlers[w] = h
	}
	return h
}

func (i *InterfaceHandler) Attach(r rune, w userinterface.Widget) {
	i.widgets[r] = w
}
