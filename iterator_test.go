package main

import "testing"

func BenchmarkIntsCallbackIterator(b *testing.B) {
	InitInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		cb := func(val int) {
			sum += val
		}
		IntCallbackIterator(cb)
	}
}

func BenchmarkDataCallbackIterator(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		cb := func(val int) {
			sum += val
		}
		DataCallbackIterator(cb)
	}
}

func BenchmarkIntsChannelIterator(b *testing.B) {
	InitInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		for val := range IntChannelIterator() {
			sum += val
		}
	}
}

func BenchmarkDataChannelIterator(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		for val := range DataChannelIterator() {
			sum += val
		}
	}
}

func BenchmarkIntsBufferedChannelIterator(b *testing.B) {
	InitInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		for val := range IntBufferedChannelIterator() {
			sum += val
		}
	}
}

func BenchmarkDataBufferedChannelIterator(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		for val := range DataBufferedChannelIterator() {
			sum += val
		}
	}
}

func BenchmarkIntsClosureIterator(b *testing.B) {
	InitInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum, val int = 0, 0
		for it, has_next := IntClosureIterator(); has_next; val, has_next = it() {
			sum += val
		}
	}
}

func BenchmarkDataClosureIterator(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum, val int = 0, 0
		for it, has_next := DataClosureIterator(); has_next; val, has_next = it() {
			sum += val
		}
	}
}

func BenchmarkIntStatefulIterator(b *testing.B) {
	InitInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		it := NewIntStatefulIterator(int_data)
		for it.Next() {
			sum += it.Value()
		}
	}
}

func BenchmarkDataStatefulIterator(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		it := NewDataStatefulIterator(struct_data)
		for it.Next() {
			sum += it.Value()
		}
	}
}

func BenchmarkIntStatefulIteratorInterface(b *testing.B) {
	InitInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		it := NewIntStatefulIteratorInterface(int_data)
		for it.Next() {
			sum += it.Value()
		}
	}
}

func BenchmarkDataStatefulIteratorInterface(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int = 0
		it := NewDataStatefulIteratorInterface(struct_data)
		for it.Next() {
			sum += it.Value()
		}
	}
}
