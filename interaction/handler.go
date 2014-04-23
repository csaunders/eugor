package interaction

import (
	"eugor/dungeon"
	"eugor/sprite"
	"github.com/nsf/termbox-go"
)

type Handler interface {
	CanHandle(e termbox.Event) bool
	Handle(e termbox.Event) (rune, func(rune, dungeon.TileMap, sprites.Sprite))
}
