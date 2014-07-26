package eugor

type Performer interface {
	Perform(world interface{}, clock interface{})
}

type KeyboardHandler struct {
}
