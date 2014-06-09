package logger

import (
	"eugor/termboxext"
	"fmt"
	"github.com/nsf/termbox-go"
)

var nilEvent Event = Event{}

type Logger struct {
	events []Event
	Render bool
}

var GlobalLog = &Logger{}

func (l *Logger) Append(e termbox.Event) {
	var message string
	switch e.Key {
	default:
		message = string(e.Ch)
	case termbox.KeyArrowUp:
		message = string('↑')
	case termbox.KeyArrowDown:
		message = string('↓')
	case termbox.KeyArrowLeft:
		message = string('←')
	case termbox.KeyArrowRight:
		message = string('→')
	}
	event := Event{LogLevel: Info, Message: fmt.Sprintf("Received event %s", message)}
	l.AppendEvent(event)
}

func (l *Logger) AppendEvent(event Event) {
	l.events = append(l.events, nilEvent)
	copy(l.events[1:], l.events[0:])
	l.events[0] = event
}

func (l *Logger) ToggleRender() {
	l.Render = !l.Render
}

func (l *Logger) Draw() {
	if l.Render == false {
		return
	}
	l.clear()
	l.drawLogEvents()
	l.drawBorder()
}

func (l *Logger) StartingY() int {
	_, height := termbox.Size()
	return height - (height / 4)
}

func (l *Logger) clear() {
	width, height := termbox.Size()
	termboxext.Fill(0, l.StartingY(), width, height/4, ' ', termbox.ColorCyan, termbox.ColorBlack)
}

func (l *Logger) drawBorder() {
	width, height := termbox.Size()
	termboxext.DrawSimpleBox(0, l.StartingY(), width, height/4, termbox.ColorCyan, termbox.ColorBlack)
}

func (l *Logger) drawLogEvents() {
	startingY := l.StartingY() + 1
	if len(l.events) == 0 {
		termboxext.DrawString(1, startingY, "There is nothing to log", termbox.ColorWhite, termbox.ColorBlack)
	}
	for index, event := range l.events {
		termboxext.DrawString(1, startingY+index, event.Message, termbox.ColorWhite, termbox.ColorBlack)
	}
}

type LogLevel uint8

const (
	Debug LogLevel = 0xFF - iota
	Info
	Warn
	Error
)

type Event struct {
	LogLevel LogLevel
	Message  string
}
