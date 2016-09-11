package main

import "fmt"

func main() {
	l := &List{}

	l.Push(1)
	l.Push(2)
	l.Push(3)

	l.Status()

	l.Reverse()

	l.Status()
}

type Node struct {
	val  int
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(element int) *List {
	n := &Node{val: element}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}

	l.tail = n
	return l
}

func (l *List) Status() {
	for n := l.First(); n != nil; n = n.Next() {
		fmt.Printf("%#v ", n.val)
	}
	fmt.Printf("\n")
}

func (l *List) Delete(element int) bool {
	success := false
	var prevNode = l.First()
	for n := l.First(); n != nil && !success; n = n.Next() {
		if n.val == element {
			if n == prevNode {
				l.head = nil
			}
			prevNode.next = n.next
			success = true
		}

		prevNode = n
	}

	return success
}

func (l *List) Reverse() {
	var prevNode *Node
	var nextNode = l.First()
	var n = l.First()

	for {
		if n == nil {
			break
		}
		nextNode = n.next
		n.next = prevNode
		prevNode = n
		n = nextNode
	}

	l.tail = l.head
	l.head = prevNode
}

func (l *List) Find(element int) *Node {
	found := false
	var foundNode *Node
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.val == element {
			found = true
			foundNode = n
		}
	}
	return foundNode
}
