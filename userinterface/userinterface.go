package userinterface

import "eugor/handler"

type Widget interface {
	Draw()
	SetPosition(x, y int)
	SetSize(w, h int)
	Show()
	Hide()
	Toggle()
	SetHandler(h handler.Handler)
	Handler() handler.Handler
}
