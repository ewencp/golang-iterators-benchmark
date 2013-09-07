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
		for it, valid := IntClosureIterator(); valid; val, valid = it() {
			if !valid {
				break
			}
			sum += val
		}
	}
}

func BenchmarkDataClosureIterator(b *testing.B) {
	InitData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum, val int = 0, 0
		for it, valid := DataClosureIterator(); valid; val, valid = it() {
			if !valid {
				break
			}
			sum += val
		}
	}
}
