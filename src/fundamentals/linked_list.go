package fundamentals

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) AddAtBegning(data int) {
	newNode := &Node{Value: data}
	newNode.Next = ll.Head
	ll.Head = newNode
}
