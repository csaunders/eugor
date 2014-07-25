package eugor

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	ts "testing"
)

func TestActionQueue_Push_Peek(t *ts.T) {
	aq := make(ActionQueue, 0)
	a1 := &Action{happensAt: 1}
	a2 := &Action{happensAt: 3}
	heap.Push(&aq, a2)
	heap.Push(&aq, a1)
	assert.Equal(t, a1, aq.Peek(), "The item at the top of the queue should be the one with the lowest happensAt value")
}

func TestActionQueue_Push_Pop(t *ts.T) {
	aq := make(ActionQueue, 0)
	heap.Push(&aq, &Action{happensAt: 1})
	heap.Push(&aq, &Action{happensAt: 3})
	heap.Push(&aq, &Action{happensAt: 2})
	assert.Equal(t, 1, heap.Pop(&aq).(*Action).happensAt)
	assert.Equal(t, 2, heap.Pop(&aq).(*Action).happensAt)
}

func TestClock_ActsIn_Tick(t *ts.T) {
	clock := MakeClock()
	clock.ActsIn(1, 10)
	assert.Equal(t, 1, clock.Tick().(int))
}

func TestClock_ActsIn_Many_Tick(t *ts.T) {
	clock := MakeClock()
	clock.ActsIn(1, 10)
	clock.ActsIn(2, 5)
	assert.Equal(t, 2, clock.Tick().(int))
	assert.Equal(t, 5, clock.Time())
	clock.ActsIn(3, 2)
	assert.Equal(t, 3, clock.Tick().(int))
	assert.Equal(t, 7, clock.Time())
}
