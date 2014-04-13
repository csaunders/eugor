package sprites

import (
	"eugor/dungeon"
	"eugor/termboxext"
	"fmt"
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

func (m MapContext) IsFocused() bool {
	return m.render
}

func (m MapContext) HandleInput(point dungeon.Point, event termbox.Event) MapContext {
	switch event.Key {
	case termbox.KeyArrowUp:
		m.cursor = (m.cursor - 1) % len(m.interactions)
	case termbox.KeyArrowDown:
		m.cursor = (m.cursor + 1) % len(m.interactions)
	case termbox.KeyEnter:
		i := 0
		for _, p := range m.currentInteractions {
			if m.cursor == i {
				m.TileMap = m.PerformInteraction(p)
				break
			}
			i++
		}
		m = m.Toggle(point)
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
		if m.cursor+1 == y {
			termboxext.DrawString(1, y, "*", termbox.ColorMagenta, termbox.ColorBlack)
		}
		termboxext.DrawString(2, y, name, termbox.ColorGreen, termbox.ColorBlack)
		y++
	}
}

func (m MapContext) Interactions(point dungeon.Point) map[string]dungeon.Point {
	availableActions := m.interactables(point)
	result := make(map[string]dungeon.Point)
	for p, interaction := range availableActions {
		name := buildNameFromPoint(point, p, interaction.Name)
		result[name] = p
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

func buildNameFromPoint(source, dest dungeon.Point, name string) string {
	var destination string
	switch dungeon.DetermineDirection(source, dest) {
	case dungeon.North:
		destination = "North"
	case dungeon.South:
		destination = "South"
	case dungeon.East:
		destination = "East"
	case dungeon.West:
		destination = "West"
	default:
		destination = "¯\\(°_o)/¯"
	}
	return fmt.Sprintf("%s - %s", name, destination)
}
