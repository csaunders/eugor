package interaction

import (
"eugor/sprites"
"eugor/dungeon"
"github.com/nsf/termbox-go"
)

type MovementHandler struct {}

func (m *MovementHandler) Handle(e termbox.Event) (rune, func(rune, dungeon.TileMap, sprites.Sprite)) {
  r := modifierKeysToRunes[e.Key]
  if r != nil {
    return r, moveToPosition
  }
}

func moveToPosition(r rune, d *dungeon.TileMap, s *Sprite) {
  x, y := s.PredictedMovement(r)
  switch {
  case d.CanMoveTo(x, y):
    s.Move(x, y)
  case d.CanInteractWith(x, y):
    d.Interact(x, y)
  }
}

modifierKeysToRunes := map[termbox.Key]rune {
  termbox.KeyArrowUp : '↑',
  termbox.KeyArrowDown : '↓',
  termbox.KeyArrowLeft : '←',
  termbox.KeyArrowRight : '→',
}
