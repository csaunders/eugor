package sprites

import (
	"eugor"
	"eugor/dungeon"
	"fmt"
	"github.com/csaunders/windeau"
	"github.com/nsf/termbox-go"
)

type Interactable struct {
	Test   func(point eugor.Point, tileMap *dungeon.TileMap) bool
	Action func(point eugor.Point, tileMap *dungeon.TileMap)
	Name   string
}

type MapContext struct {
	windeau.EventHandler
	TileMap             *dungeon.TileMap
	render              bool
	cursor              int
	currentInteractions map[string]eugor.Point
	interactions        []Interactable
	view                *windeau.Scrollview
}

func DefaultMapContext(d *dungeon.TileMap) *MapContext {
	context := &MapContext{TileMap: d, cursor: 0}
	context.AddInteraction(closeDoorHandler())
	context.AddInteraction(openDoorHandler())
	context.view = windeau.MakeScrollview(buildWindow(), []string{}, context)
	return context
}

func (m *MapContext) AddInteraction(i Interactable) {
	m.interactions = append(m.interactions, i)
}

func (m *MapContext) Toggle(point eugor.Point) {
	m.render = !m.render
	m.view.Parent.SetFocused(m.render)
	m.cursor = 0
	if m.render {
		m.currentInteractions = m.Interactions(point)
		m.updateView()
	}
}

func (m *MapContext) IsFocused() bool {
	return m.render
}

func (m *MapContext) HandleInput(point eugor.Point, event termbox.Event) {
	switch event.Key {
	case termbox.KeyArrowUp:
		m.cursor -= 1
	case termbox.KeyArrowDown:
		m.cursor += 1
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
	m.cursor = m.view.SetPosition(m.cursor)
}

func (m *MapContext) Draw() {
	if !m.render {
		return
	}
	m.view.Draw()
}

func (m *MapContext) Interactions(point eugor.Point) map[string]eugor.Point {
	availableActions := m.interactables(point)
	result := make(map[string]eugor.Point)
	for p, interaction := range availableActions {
		name := buildNameFromPoint(point, p, interaction.Name)
		result[name] = p
	}
	return result
}

func (m *MapContext) PerformInteraction(point eugor.Point) {
	for _, i := range m.interactions {
		if i.Test(point, m.TileMap) {
			i.Action(point, m.TileMap)
			break
		}
	}
}

func (m *MapContext) interactables(point eugor.Point) map[eugor.Point]Interactable {
	interactables := make(map[eugor.Point]Interactable)
	points := eugor.MakePoints(point, []string{"up", "down", "left", "right"})
	for _, interactable := range m.interactions {
		for _, p := range points {
			if interactable.Test(p, m.TileMap) {
				interactables[p] = interactable
			}
		}
	}
	return interactables
}

func (m *MapContext) updateView() {
	entries := make([]string, len(m.currentInteractions))
	i := 0
	for name, _ := range m.currentInteractions {
		entries[i] = name
		i++
	}
	m.cursor = 0
	m.view.SetPosition(m.cursor)
	m.view.Entries = entries
}

func buildNameFromPoint(source, dest eugor.Point, name string) string {
	var destination string
	switch eugor.DetermineDirection(source, dest) {
	case eugor.North:
		destination = "North"
	case eugor.South:
		destination = "South"
	case eugor.East:
		destination = "East"
	case eugor.West:
		destination = "West"
	default:
		destination = "¯\\(°_o)/¯"
	}
	return fmt.Sprintf("%s - %s", name, destination)
}

func openDoorHandler() Interactable {
	return Interactable{
		Name: "Open Door",
		Test: func(p eugor.Point, d *dungeon.TileMap) bool {
			tile := d.FetchTile(p.X, p.Y)
			return d.CanInteractWith(p.X, p.Y) && tile.Name == "door"
		},
		Action: func(p eugor.Point, d *dungeon.TileMap) {
			d.Interact(p.X, p.Y)
		},
	}
}

func closeDoorHandler() Interactable {
	return Interactable{
		Name: "Close Door",
		Test: func(p eugor.Point, d *dungeon.TileMap) bool {
			tile := d.FetchTile(p.X, p.Y)
			return d.CanInteractWith(p.X, p.Y) && tile.Name == "opendoor"
		},
		Action: func(p eugor.Point, d *dungeon.TileMap) {
			d.Interact(p.X, p.Y)
		},
	}
}

func buildWindow() *windeau.FocusableWindow {
	on := windeau.WindowState{termbox.ColorWhite, termbox.ColorDefault}
	border := windeau.MakeSimpleBorder('+', '|', '-')
	return windeau.MakeFocusableWindow(0, 0, 20, 6, on, on, border)
}
