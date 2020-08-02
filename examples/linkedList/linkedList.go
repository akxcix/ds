package main

import (
	"github.com/iamadarshk/ds"
)

func main() {
	var ll ds.LinkedList
	ll.Print()                       // []
	ll.Insert(ds.Int(1))             // [1]
	ll.Insert(ds.Int(2))             // [2 1]
	ll.Insert(ds.Int(-1))            // [-1 2 1]
	ll.Insert(ds.Int(4) + ds.Int(5)) // [9 -1 2 1]
	ll.Print()                       // Prints [9 -1 2 1]
	ll.Delete(ds.Int(1))             // [9 -1 2]
	ll.Insert(ds.Int(7))             // [7 9 -1 2]
	ll.Delete(ds.Int(7))             // [9 -1 2]
	ll.Delete(ds.Int(-1))            // [9 2]
	_ = ll.Search(ds.Int(9))         // returns &{9 ADDRESS}
	ll.IsEmpty()                     // false
	ll.Traverse()                    // [9 2]
	ll.Print()
}
