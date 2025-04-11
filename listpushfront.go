package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Tail == nil {
		l.Tail = newNode
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
