package main

import (
	"iter"
	"math/rand"
)

const NumItems int = 1000000

// Plain int array tests how quickly we can iterate over cache-coherent data
var int_data []int = make([]int, NumItems)

func InitInts() {
	for i := 0; i < NumItems; i++ {
		int_data[i] = rand.Int()
	}
}

// Iterating over the data in a list of pointers tests data that isn't
// cache coherent (following the references ruins the coherency)
type Data struct {
	// This intentionally contains more than just an int so it may
	// not be as conveniently size if allocations end up in order.
	foo int
	bar *Data
}

var struct_data []*Data = make([]*Data, NumItems)

func InitData() {
	for i := 0; i < NumItems; i++ {
		struct_data[i] = &Data{foo: rand.Int(), bar: nil}
	}
}

// These are the different types of iterators:
// Callbacks:
func IntCallbackIterator(cb func(int)) {
	for _, val := range int_data {
		cb(val)
	}
}
func DataCallbackIterator(cb func(int)) {
	for _, val := range struct_data {
		cb(val.foo)
	}
}

// Builtin "generic" push iterator (Golang 1.23+):
func IntBuiltinPushIterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, val := range int_data {
			if !yield(val) {
				break
			}
		}
	}
}
func DataBuiltinPushIterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, val := range struct_data {
			if !yield(val.foo) {
				break
			}
		}
	}
}

// Builtin "generic" pull iterator (Golang 1.23+):
func IntBuiltinPullIterator() (next func() (int, bool), stop func()) {
	return iter.Pull(IntBuiltinPushIterator())
}
func DataBuiltinPullIterator() (next func() (int, bool), stop func()) {
	return iter.Pull(DataBuiltinPushIterator())
}

// Builtin "generic" pull iterator, implemented manually to avoid channels in generic iter.Pull implementation (Golang 1.23+):
func IntBuiltinPullManualIterator() (next func() (int, bool), stop func()) {
	var idx = 0
	var data_len = len(int_data)
	next = func() (int, bool) {
		if idx >= data_len {
			return 0, false
		}
		prev_idx := idx
		idx++
		return int_data[prev_idx], true
	}
	stop = func() {
		idx = data_len
	}
	return
}
func DataBuiltinPullManualIterator() (next func() (int, bool), stop func()) {
	var idx = 0
	var data_len = len(struct_data)
	next = func() (int, bool) {
		if idx >= data_len {
			return 0, false
		}
		prev_idx := idx
		idx++
		return struct_data[prev_idx].foo, true
	}
	stop = func() {
		idx = data_len
	}
	return
}

// Channels:
func IntChannelIterator() <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range int_data {
			ch <- val
		}
		close(ch)
	}()
	return ch
}
func DataChannelIterator() <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range struct_data {
			ch <- val.foo
		}
		close(ch)
	}()
	return ch
}

// Buffered Channels:
const ChannelBuffer int = 50

func IntBufferedChannelIterator() <-chan int {
	ch := make(chan int, ChannelBuffer)
	go func() {
		for _, val := range int_data {
			ch <- val
		}
		close(ch)
	}()
	return ch
}
func DataBufferedChannelIterator() <-chan int {
	ch := make(chan int, ChannelBuffer)
	go func() {
		for _, val := range struct_data {
			ch <- val.foo
		}
		close(ch)
	}()
	return ch
}

// Closures: Return (next(), has_next), where next() returns (val, has_next)
func IntClosureIterator() (func() (int, bool), bool) {
	var idx int = 0
	var data_len = len(int_data)
	return func() (int, bool) {
		prev_idx := idx
		idx++
		return int_data[prev_idx], (idx < data_len)
	}, (idx < data_len)
}

func DataClosureIterator() (func() (int, bool), bool) {
	var idx int = 0
	var data_len = len(struct_data)
	return func() (int, bool) {
		prev_idx := idx
		idx++
		return struct_data[prev_idx].foo, (idx < data_len)
	}, (idx < data_len)
}

type StatefulIterator interface {
	Value() int
	Next() bool
}

type dataStatefulIterator struct {
	current int
	data    []*Data
}

func NewDataStatefulIterator(data []*Data) *dataStatefulIterator {
	return &dataStatefulIterator{data: data, current: -1}
}

func NewDataStatefulIteratorInterface(data []*Data) StatefulIterator {
	return &dataStatefulIterator{data: data, current: -1}
}

func (it *dataStatefulIterator) Value() int {
	return it.data[it.current].foo
}
func (it *dataStatefulIterator) Next() bool {
	it.current++
	if it.current >= len(it.data) {
		return false
	}
	return true
}

type intStatefulIterator struct {
	current int
	data    []int
}

func (it *intStatefulIterator) Value() int {
	return it.data[it.current]
}
func (it *intStatefulIterator) Next() bool {
	it.current++

	if it.current >= len(it.data) {
		return false
	}
	return true
}

func NewIntStatefulIterator(data []int) *intStatefulIterator {
	return &intStatefulIterator{data: data, current: -1}
}

func NewIntStatefulIteratorInterface(data []int) StatefulIterator {
	return &intStatefulIterator{data: data, current: -1}
}
