package main

import (
	"github.com/iamadarshk/ds"
)

func main() {
	stack := ds.Stack{}
	stack.Print()         // []
	stack.Pop()           // returns <nil>, ErrEmptyStack. stack is []
	stack.Push(ds.Int(1)) // [1]
	stack.Push(ds.Int(2)) // [2 1]
	stack.Push(ds.Int(3)) // [3 2 1]
	stack.Traverse()      // returns [3 2 1]
	stack.Print()         // prints [3 2 1]
	stack.Pop()           // returns &{3 ADDRESS}, <nil>. stack is [2 1]
	stack.Pop()           // returns &{2 ADDRESS}, <nil>. stack is [1]
	stack.Pop()           // returns &{1 <nil>}, <nil>. stack is []
	stack.Print()         // prints []
}
