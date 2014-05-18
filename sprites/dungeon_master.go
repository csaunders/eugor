package sprites

import (
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

func populateDungeon(d *dungeon.TileMap, difficulty int) []*Creature {
	prng := rand.New(rand.NewSource(time.Now().UnixNano()))
	names := MonsterNames()
	monsters := make([]*Creature, difficulty)
	for i := 0; i < difficulty; i++ {
		monsterName := names[prng.Intn(len(names))]
		m := genMonster(monsterName, d, prng)
		monsters = append(monsters, m)
	}
	return monsters
}

func genMonster(name string, d *dungeon.TileMap, prng *rand.Rand) *Creature {
	var x, y = -1, -1
	for !d.FetchTile(x, y).Walkable {
		x = prng.Intn(d.Width)
		y = prng.Intn(d.Height)
	}
	return Monsters[name](x, y, d)
}
