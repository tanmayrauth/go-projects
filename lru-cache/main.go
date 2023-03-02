package main

import (
	"fmt"
)

const SIZE = 2

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	head.Left = head

	return Queue{Head: head, Tail: tail}
}

// Struct methods
func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node

}

func (c *Cache) Remove(node *Node) *Node {
	fmt.Printf("Removing %s\n", node.Val)
	left := node.Left
	right := node.Right

	right.Left = left
	left.Right = right

	c.Queue.Length -= 1
	delete(c.Hash, node.Val)
	return node

}

func (c *Cache) Add(node *Node) {
	fmt.Printf("Adding Node %s\n", node.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Left = c.Queue.Head
	node.Right = tmp
	tmp.Left = node

	c.Queue.Length += 1

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Starting Cache...")
	cache := NewCache()
	for _, node := range []string{"Mango", "Orange", "Avocado"} {
		cache.Check(node)
		cache.Display()
	}

}
