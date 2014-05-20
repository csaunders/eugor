package sprites

import (
	"eugor/algebra"
	"eugor/dungeon"
	"eugor/termboxext"
	"fmt"
	"github.com/nsf/termbox-go"
)

type Interactable struct {
	Test   func(point algebra.Point, tileMap *dungeon.TileMap) bool
	Action func(point algebra.Point, tileMap *dungeon.TileMap)
	Name   string
}

type MapContext struct {
	TileMap             *dungeon.TileMap
	render              bool
	cursor              int
	currentInteractions map[string]algebra.Point
	interactions        []Interactable
}

func DefaultMapContext(d *dungeon.TileMap) *MapContext {
	context := &MapContext{TileMap: d}
	context.AddInteraction(closeDoorHandler())
	context.AddInteraction(openDoorHandler())
	return context
}

func (m *MapContext) AddInteraction(i Interactable) {
	m.interactions = append(m.interactions, i)
}

func (m *MapContext) Toggle(point algebra.Point) {
	m.render = !m.render
	m.cursor = 0
	if m.render {
		m.currentInteractions = m.Interactions(point)
	}
}

func (m *MapContext) IsFocused() bool {
	return m.render
}

func (m *MapContext) HandleInput(point algebra.Point, event termbox.Event) {
	switch event.Key {
	case termbox.KeyArrowUp:
		m.cursor = (m.cursor - 1) % len(m.interactions)
	case termbox.KeyArrowDown:
		m.cursor = (m.cursor + 1) % len(m.interactions)
	case termbox.KeyEnter:
		i := 0
		for _, p := range m.currentInteractions {
			if m.cursor == i {
				m.PerformInteraction(p)
				break
			}
			i++
		}
		m.Toggle(point)
	}
}

func (m *MapContext) Draw() {
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

func (m *MapContext) Interactions(point algebra.Point) map[string]algebra.Point {
	availableActions := m.interactables(point)
	result := make(map[string]algebra.Point)
	for p, interaction := range availableActions {
		name := buildNameFromPoint(point, p, interaction.Name)
		result[name] = p
	}
	return result
}

func (m *MapContext) PerformInteraction(point algebra.Point) {
	for _, i := range m.interactions {
		if i.Test(point, m.TileMap) {
			i.Action(point, m.TileMap)
			break
		}
	}
}

func (m *MapContext) interactables(point algebra.Point) map[algebra.Point]Interactable {
	result := make(map[algebra.Point]Interactable)
	points := algebra.MakePoints(point, []string{"up", "down", "left", "right"})
	for _, interactable := range m.interactions {
		for _, p := range points {
			if interactable.Test(p, m.TileMap) {
				result[p] = interactable
			}
		}
	}
	return result
}

func buildNameFromPoint(source, dest algebra.Point, name string) string {
	var destination string
	switch algebra.DetermineDirection(source, dest) {
	case algebra.North:
		destination = "North"
	case algebra.South:
		destination = "South"
	case algebra.East:
		destination = "East"
	case algebra.West:
		destination = "West"
	default:
		destination = "¯\\(°_o)/¯"
	}
	return fmt.Sprintf("%s - %s", name, destination)
}

func openDoorHandler() Interactable {
	return Interactable{
		Name: "Open Door",
		Test: func(p algebra.Point, d *dungeon.TileMap) bool {
			tile := d.FetchTile(p.X, p.Y)
			return d.CanInteractWith(p.X, p.Y) && tile.Name == "door"
		},
		Action: func(p algebra.Point, d *dungeon.TileMap) {
			d.Interact(p.X, p.Y)
		},
	}
}

func closeDoorHandler() Interactable {
	return Interactable{
		Name: "Close Door",
		Test: func(p algebra.Point, d *dungeon.TileMap) bool {
			tile := d.FetchTile(p.X, p.Y)
			return d.CanInteractWith(p.X, p.Y) && tile.Name == "opendoor"
		},
		Action: func(p algebra.Point, d *dungeon.TileMap) {
			d.Interact(p.X, p.Y)
		},
	}
}
