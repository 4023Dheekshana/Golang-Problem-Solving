package main

import (
	"container/list"
	"fmt"
)

// queue can be implemented by using list and slice
// Implemented by list
type customQueue struct {
	queue *list.List
}

func (c *customQueue) enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("pop error")
}

func (c *customQueue) Size() int16 {
	return int16(c.queue.Len())
}

func (c *customQueue) Front() string {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val
		}
		return "query datatype is invalid"
	}
	return "queue is empty"
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

func main() {
	c := &customQueue{
		queue: list.New(),
	}
	c.enqueue("A")
	c.enqueue("B")
	c.enqueue("C")
	c.enqueue("D")

	fmt.Println(c.Size())
	if c.Size() > 0 {
		fmt.Println(c.Front())
	}
	c.Dequeue()
	fmt.Println(c.Size())
	if c.Size() > 0 {
		fmt.Println(c.Front())
	}

}
