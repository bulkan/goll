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

func (ll *LinkedList) AddToHead(data int) {
    ll.head = &node{data:data, next:ll.head}
}
