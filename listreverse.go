package piscine

/*
	type NodeL struct {
		Data interface{}
		Next *NodeL
	}

	type List struct {
		Head *NodeL
		Tail *NodeL
	}
*/
func ListReverse(l *List) {
	if l.Head == nil {
		return
	}

	temp := l.Head
	current := l.Head
	var prev *NodeL

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail = temp
	l.Tail.Next = nil
}
