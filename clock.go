package eugor

import (
	"container/heap"
	"fmt"
)

type Action struct {
	actor     interface{}
	happensAt int
}

type ActionQueue []*Action

func (aq ActionQueue) Peek() *Action {
	if len(aq) <= 0 {
		return nil
	}
	return aq[0]
}

func (aq ActionQueue) Len() int { return len(aq) }

func (aq ActionQueue) Less(i, j int) bool {
	return aq[i].happensAt < aq[j].happensAt
}

func (aq ActionQueue) Swap(i, j int) {
	aq[i], aq[j] = aq[j], aq[i]
}

func (aq *ActionQueue) Push(x interface{}) {
	action := x.(*Action)
	*aq = append(*aq, action)
}

func (aq *ActionQueue) Pop() interface{} {
	old := *aq
	n := len(old)
	action := old[n-1]
	*aq = old[0 : n-1]
	return action
}

func (aq ActionQueue) String() string {
	result := "["
	for _, item := range aq {
		result = result + fmt.Sprintf("%d, ", item.happensAt)
	}
	result = result + "]"
	return result
}

type Clock struct {
	currentTime int
	actions     ActionQueue
}

func MakeClock() *Clock {
	return &Clock{currentTime: 0, actions: make(ActionQueue, 0)}
}

func (c *Clock) ActsIn(actor interface{}, ticks int) {
	action := &Action{actor: actor, happensAt: c.currentTime + ticks}
	heap.Push(&c.actions, action)
}

func (c *Clock) Time() int {
	return c.currentTime
}

func (c *Clock) Tick() interface{} {
	action := heap.Pop(&c.actions).(*Action)
	c.currentTime = action.happensAt
	return action.actor
}
