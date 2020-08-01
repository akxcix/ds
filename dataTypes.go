package ds

// Item is an interface which is used so that data structures which
// are based on order such as Binary Search Trees can be implemented easily. Any struct that
// implements the Less method can be stored in the Nodes of various data structures
type Item interface {
	Less(than Item) bool
}

// Int provides Less interface to int
type Int int

// Less returns true if int a < int b
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}
