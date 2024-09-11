package piscine

func ListPushFront(l *List, data interface{}) {
	newnode := &NodeL{Data: data, Next: l.Head}

	if l.Head == nil {
		l.Head = newnode
	} else {
		newnode.Next = l.Head
		l.Head = newnode
	}
}
