package sprites

import (
	"eugor"
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

func (dm *DungeonMaster) Tick(playerPosition eugor.Point) {
	for _, m := range dm.monsters {
		m.Tick(playerPosition)
	}
}

func (dm *DungeonMaster) Drawables() (drawables []eugor.Drawable) {
	drawables = make([]eugor.Drawable, len(dm.monsters))
	for i, m := range dm.monsters {
		drawables[i] = m
	}
	return
}

func (dm *DungeonMaster) Occupied(x, y int) bool {
	m, _ := dm.retrieveMonster(x, y)
	return m != nil
}

func (dm *DungeonMaster) Interact(x, y int, char *Character) bool {
	m, i := dm.retrieveMonster(x, y)
	if m == nil {
		return false
	}

	attack := char.AttackAttribute()
	defense := Ability{Name: "dodge", Modifier: "dex", SuccessRate: 50}
	isHit := attack.Challenge(defense)
	if isHit {
		dm.monsters[i] = dm.monsters[len(dm.monsters)-1]
		dm.monsters = dm.monsters[0 : len(dm.monsters)-1]
		return true
	}
	return false
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

func (dm *DungeonMaster) retrieveMonster(x, y int) (*Creature, int) {
	for i, m := range dm.monsters {
		if m.X() == x && m.Y() == y {
			return m, i
		}
	}
	return nil, -1
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
