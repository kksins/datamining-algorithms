package heap

import (
	//	"fmt"
	"sync"
)

const (
	MAXHEAP = "maxHeap"
	MINHEAP = "minHeap"
)

type Item interface {
	Less(i Item) bool
}

type SimpleItem struct {
	item  interface{}
	value float64
}

type Heap struct {
	sync.Mutex
	heap []Item
	attr string
}

func NewSimpleItem(item interface{}, value float64) SimpleItem {
	return SimpleItem{item: item, value: value}
}
func (simpleItem SimpleItem) Less(i Item) bool {
	ii := i.(SimpleItem)
	return simpleItem.value < ii.value
}
func (simpleItem SimpleItem) GetItem() interface{} {
	return simpleItem.item
}

func NewMaxHeap(heap []Item) Heap {
	tmpHeap := make([]Item, len(heap))
	copy(tmpHeap, heap)
	buildMaxHeap(tmpHeap)
	return Heap{heap: tmpHeap, attr: MAXHEAP}
}

func NewMinHeap(heap []Item) Heap {
	tmpHeap := make([]Item, len(heap))
	copy(tmpHeap, heap)
	buildMinHeap(tmpHeap)
	return Heap{heap: tmpHeap, attr: MINHEAP}
}

func (h Heap) GetSize() int {
	return len(h.heap)
}

func (h Heap) IsEmpty() bool {
	return len(h.heap) == 0
}
func (h Heap) Get(i int) Item {
	return h.heap[i]
}

func (h Heap) GetItems() []Item {
	size := len(h.heap)
	items := make([]Item, size)
	copy(items, h.heap)
	return items
}

func (h Heap) GetAttr() string {
	return h.attr
}

func (h Heap) FindBottom() (index int, value Item) {
	size := len(h.heap)
	if size == 0 {
		return -1, nil
	}
	startIndex := size/2 + 1
	max, min := h.heap[startIndex-1], h.heap[startIndex-1]
	maxIndex, minIndex := startIndex-1, startIndex-1
	for i := startIndex + 1; i <= size; i++ {
		cur := h.heap[i-1]
		if max.Less(cur) {
			max, maxIndex = cur, i-1
		}
		if cur.Less(min) {
			min, minIndex = cur, i-1
		}
	}
	var resultIndex int
	if h.attr == MAXHEAP {
		resultIndex = minIndex
	} else if h.attr == MINHEAP {
		resultIndex = maxIndex
	}
	return resultIndex, h.heap[resultIndex]
}

func (h *Heap) InsertValue(value Item) bool {
	h.Lock()
	defer h.Unlock()
	size := len(h.heap)
	if size == 0 {
		h.heap = append(h.heap, value)
		return true
	}
	if h.attr == MAXHEAP && value.Less(h.heap[0]) {
		h.heap[0] = value
		adjustMaxHeapDown(h.heap, 1, size)
		return true
	} else if h.attr == MINHEAP && h.heap[0].Less(value) {
		h.heap[0] = value
		adjustMinHeapDown(h.heap, 1, size)
		return true
	}
	return false
}
func (h *Heap) Delete(i int) {
	if len(h.heap) == 0 {
		return
	} else if len(h.heap) == 1 {
		if i == 0 {
			h.heap = make([]Item, 0)
			return
		}
		panic("index error in Delete()")
	}
	h.Lock()
	defer h.Unlock()
	h.heap = append(h.heap[:i], h.heap[i+1:]...)
	if h.attr == MAXHEAP {
		buildMaxHeap(h.heap)
	} else if h.attr == MINHEAP {
		buildMinHeap(h.heap)
	}
}

// add a value to the heap
func (h *Heap) Add(value Item) {
	h.Lock()
	defer h.Unlock()
	h.heap = append(h.heap, value)
	if h.attr == MAXHEAP {
		adjustMaxHeapUp(h.heap, len(h.heap), 1)
	} else if h.attr == MINHEAP {
		adjustMinHeapUp(h.heap, len(h.heap), 1)
	} else {
		panic("heap type error(not maxHeap/minHeap)")
	}
}

// TODD:merge buildMaxHeap and buildMinHeap
///////////////////////////////////////////////////////
/////////////// build max heap ////////////////////////
func buildMaxHeap(array []Item) {
	size := len(array)
	for i := size / 2; i > 0; i -= 1 {
		adjustMaxHeapDown(array, i, size)
	}
}

func adjustMaxHeapDown(array []Item, start, end int) {
	tmp := array[start-1]
	for i := 2 * start; i <= end; i *= 2 {
		if i < end && array[i-1].Less(array[i]) {
			i++
		}
		if array[i-1].Less(tmp) {
			break
		}
		array[start-1], start = array[i-1], i
	}
	array[start-1] = tmp
}
func adjustMaxHeapUp(array []Item, start, end int) {
	tmp := array[start-1]
	for i := start / 2; i >= end; i /= 2 {
		if tmp.Less(array[i-1]) {
			break
		}
		array[start-1] = array[i-1]
		start = i
	}
	array[start-1] = tmp
}

///////////////////////////////////////////////////////
/////////////// build min heap ////////////////////////
func buildMinHeap(array []Item) {
	size := len(array)
	for i := size / 2; i > 0; i -= 1 {
		adjustMinHeapDown(array, i, size)
	}
}

func adjustMinHeapDown(array []Item, start, end int) {
	tmp := array[start-1]
	for i := 2 * start; i <= end; i *= 2 {
		if i < end && array[i].Less(array[i-1]) {
			i++
		}
		if tmp.Less(array[i-1]) {
			break
		}
		array[start-1], start = array[i-1], i
	}
	array[start-1] = tmp
}
func adjustMinHeapUp(array []Item, start, end int) {
	tmp := array[start-1]
	for i := start / 2; i >= end; i /= 2 {
		if array[i-1].Less(tmp) {
			break
		}
		array[start-1] = array[i-1]
		start = i
	}
	array[start-1] = tmp
}
