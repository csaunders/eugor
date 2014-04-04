package logger

import (
	"eugor/termboxext"
	"fmt"
	"github.com/nsf/termbox-go"
)

type Logger struct {
	events []Event
}

func (l Logger) Append(e termbox.Event) {
	event := Event{logLevel: Info, message: fmt.Sprintf("Received event %v", e.Ch)}
	l.events = append(l.events, event)
}

func (l Logger) Draw() {
	l.drawLogEvents()
	l.drawBorder()
}

func (l Logger) StartingY() int {
	_, height := termbox.Size()
	return height - (height / 4)
}

func (l Logger) drawBorder() {
	width, height := termbox.Size()
	termboxext.DrawSimpleBox(0, l.StartingY(), width, height, termbox.ColorCyan, termbox.ColorBlack)
}

func (l Logger) drawLogEvents() {
	_, height := termbox.Size()
	startingY := height - (height / 4) + 1
	termboxext.DrawString(0, startingY, "Hello World", termbox.ColorCyan, termbox.ColorRed)
}

type LogLevel uint8

const (
	Debug LogLevel = 0xFF - iota
	Info
	Warn
	Error
)

type Event struct {
	logLevel LogLevel
	message  string
}
