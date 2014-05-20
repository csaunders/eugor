package sprites

import (
	"eugor/algebra"
	"eugor/camera"
	"eugor/dungeon"
	"math/rand"
	"time"
)

type DungeonMaster struct {
	character *Character
	monsters  []*Creature
}

func MakeDungeonMaster(character *Character, d *dungeon.TileMap) *DungeonMaster {
	dm := &DungeonMaster{character: character}
	dm.monsters = populateDungeon(d, 5)
	return dm
}

func (dm *DungeonMaster) Tick(playerPosition algebra.Point) {
	for _, m := range dm.monsters {
		m.Tick(playerPosition)
	}
}

func (dm *DungeonMaster) Drawables() (drawables []camera.Drawable) {
	drawables = make([]camera.Drawable, len(dm.monsters))
	for i, m := range dm.monsters {
		drawables[i] = m
	}
	return
}

func populateDungeon(d *dungeon.TileMap, difficulty int) []*Creature {
	prng := rand.New(rand.NewSource(time.Now().UnixNano()))
	names := MonsterNames()
	monsters := make([]*Creature, difficulty)
	for i := 0; i < difficulty; i++ {
		monsterName := names[prng.Intn(len(names))]
		m := genMonster(monsterName, d, prng)
		monsters[i] = m
	}
	return monsters
}

func genMonster(name string, d *dungeon.TileMap, prng *rand.Rand) *Creature {
	var x, y = -1, -1
	for true {
		tile := d.FetchTile(x, y)
		if tile != nil && tile.Walkable {
			break
		}
		x = prng.Intn(d.Width)
		y = prng.Intn(d.Height)
	}
	return Monsters[name](x, y, d)
}
