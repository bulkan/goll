package goll

// the basic node struct
type node struct {
    next *node
    data int
}

/*
 struct which we will add methods on that contains the
 head pointer
*/
type LinkedList struct {
    head *node
}

func (ll *LinkedList) AddToHead(data int) bool {
    tmp := &node{data:data}
    if ll.head == nil {
        ll.head = tmp
    } else {
        tmp.next = ll.head
        ll.head = tmp
    }
    return true
}
