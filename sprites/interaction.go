package sprites

import (
	"eugor/dungeon"
	"eugor/termboxext"
	"github.com/nsf/termbox-go"
)

type Interactable struct {
	Test   func(point dungeon.Point, tileMap dungeon.TileMap) bool
	Action func(point dungeon.Point, tileMap dungeon.TileMap) dungeon.TileMap
	Name   string
}

type MapContext struct {
	TileMap             dungeon.TileMap
	render              bool
	cursor              int
	currentInteractions map[string]dungeon.Point
	interactions        []Interactable
}

func (m MapContext) AddInteraction(i Interactable) MapContext {
	m.interactions = append(m.interactions, i)
	return m
}

func (m MapContext) Toggle(point dungeon.Point) MapContext {
	m.render = !m.render
	m.cursor = 0
	if m.render {
		m.currentInteractions = m.Interactions(point)
	}
	return m
}

func (m MapContext) Draw() {
	if !m.render {
		return
	}

	termboxext.Fill(0, 0, 50, 20, ' ', termbox.ColorBlack, termbox.ColorBlack)
	termboxext.DrawSimpleBox(0, 0, 50, 20, termbox.ColorGreen, termbox.ColorBlack)
	y := 1
	for name, _ := range m.currentInteractions {
		termboxext.DrawString(1, y, name, termbox.ColorGreen, termbox.ColorBlack)
		y++
	}
}

func (m MapContext) Interactions(point dungeon.Point) map[string]dungeon.Point {
	availableActions := m.interactables(point)
	result := make(map[string]dungeon.Point)
	for p, interaction := range availableActions {
		result[interaction.Name] = p
	}
	return result
}

func (m MapContext) PerformInteraction(point dungeon.Point) dungeon.TileMap {
	for _, i := range m.interactions {
		if i.Test(point, m.TileMap) {
			m.TileMap = i.Action(point, m.TileMap)
			break
		}
	}
	return m.TileMap
}

func (m MapContext) interactables(point dungeon.Point) map[dungeon.Point]Interactable {
	result := make(map[dungeon.Point]Interactable)
	points := dungeon.MakePoints(point, []string{"up", "down", "left", "right"})
	for _, interactable := range m.interactions {
		for _, p := range points {
			if interactable.Test(p, m.TileMap) {
				result[p] = interactable
			}
		}
	}
	return result
}