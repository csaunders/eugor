package sprites

type Sprite interface {
	PredictedMovement(r rune) (x, y int)
	Position() (x, y int)
	Move(x, y int)
}

func DefaultPredictedMovement(oX, oY int, r rune) (x, y int) {
	x, y = oX, oY
	switch r {
	case '↑':
		y -= 1
	case '↓':
		y += 1
	case '←':
		x -= 1
	case '→':
		x += 1
	}
	return
}
