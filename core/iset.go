package core

// ISet interface definition
type ISet interface {
	// Insert an element to the set.
	Insert(T)
	// Remove the dedicated element.
	Remove(T)
	// Check if there is a dedicated element in the set.
	Has(T) bool
	// Return if the set is empty.
	Empty() bool
	// Return how many elements contained in the set.
	Size() int
}
