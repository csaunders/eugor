package sprites

import (
	"eugor/dungeon"
	"github.com/nsf/termbox-go"
)

type MonsterMaker func(x, y int, d *dungeon.TileMap) *Creature

var Monsters map[string]MonsterMaker = map[string]MonsterMaker{
	"kobold": func(x, y int, d *dungeon.TileMap) *Creature {
		c := MakeCreature(x, y, termbox.ColorGreen, 'k')
		c.Ai = MakeWalker(d)
		return c
	},
	"rat": func(x, y int, d *dungeon.TileMap) *Creature {
		c := MakeCreature(x, y, termbox.ColorWhite, 'r')
		c.Ai = MakeRunner(d)
		return c
	},
}

var MonsterKeys []string

func MonsterNames() []string {
	if len(MonsterKeys) > 0 {
		return MonsterKeys
	}
	for name := range Monsters {
		MonsterKeys = append(MonsterKeys, name)
	}
	return MonsterKeys
}
