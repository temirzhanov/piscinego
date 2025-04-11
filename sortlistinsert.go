package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = listPushBackNode(l, data_ref)
	return ListSort(l)
}

func listPushBackNode(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data, Next: nil}

	if l == nil {
		return n
	}
	a := l
	for a.Next != nil {
		a = a.Next
	}
	a.Next = n
	return l
}
