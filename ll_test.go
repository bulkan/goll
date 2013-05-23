package goll

import "testing"

// Test that we can initialize LinkedList struct
func TestLinkListInit(t *testing.T) {
    if ll := new(LinkedList); ll.head != nil {
        t.Errorf("ll.head is not nil\n")
    }
}

func TestAddToHead(t *testing.T) {
    ll := new(LinkedList)

    // adding data a new node to empty list
    ll.AddToHead(-1)

    if ll.head == nil {
        t.Errorf("ll.head should not be nil\n")
    }

    if ll.head.data != -1 {
        t.Errorf("ll.head.data is %v should not be nil\n", ll.head.data)
    }

    // add another node to list
    ll.AddToHead(-2)

    if ll.head.data != -2 {
        t.Errorf("ll.head.data is %v should not be nil\n", ll.head.data)
    }

    // make sure that the tail data is still -1
    if ll.head.next.data != -1 {
        t.Errorf("ll.head.next.data is %v should not be nil\n", ll.head.next.data)
    }

}
