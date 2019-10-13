package heap

import "math"

type Heap struct {
	items []int
}

func (h *Heap) Swap (i,j int) {
	temp := h.items[i]
	h.items[i] = h.items[j]
	h.items[j] = temp
}

func (h *Heap) ParentIndex (i int) int {
	return int(math.Floor(float64(i-1)/2))
}

func (h *Heap) LeftChildIndex (i int) int {
	return i * 2 + 1
}

func (h *Heap) RightChildIndex (i int) int {
	return i * 2 + 2
}

func (h *Heap) Parent (i int) int {
	return h.items[h.ParentIndex(i)]
}

func (h *Heap) LeftChild (i int) int {
	return h.items[h.LeftChildIndex(i)]
}

func (h *Heap) RightChild (i int) int {
	return h.items[h.RightChildIndex(i)]
}

// Go from root to lead
func (h *Heap) PercolateDown () {
	// Start from root
	i := 0
	// If left child is smaller that root
	for i != 0 && h.LeftChild(i) < h.items[i] {
		j := h.LeftChildIndex(i)
		// Check right just in case is smaller than left
		if i != 0 && h.RightChild(i) < h.items[j] {
			// if right is smaller, right swaps
			j = h.RightChildIndex(i)
		}
		h.Swap(j, i)
		i = j
	}
}

// Go from leaf to root
func (h *Heap) PercolateUp () {
	i := len(h.items) - 1
	for i != 0 && h.Parent(i) > h.items[i] {
		h.Swap(h.ParentIndex(i), i)
		i = h.ParentIndex(i)
	}
}

// Go from root to lead
func (h *Heap) Add (value int) {
	h.items = append(h.items, value)
	h.PercolateUp()
}

// Extracts the curret min/max
func (h *Heap) Extract (value int) int {
	item := h.items[0]
	// Put the last element first
	h.items[0] = h.items[len(h.items)-1]
	// Remove last element
	h.items = h.items[len(h.items)-1:]
	// Percolate current top
	h.PercolateDown()
	return item
}