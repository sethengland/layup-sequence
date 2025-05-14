package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 10},
	{input: 20},
	{input: 50},
	{input: 100},
	{input: 200},
	{input: 500},
	{input: 1000},
	{input: 2000},
	{input: 5000},
	{input: 10000},
	{input: 20000},
	{input: 50000},
	{input: 100000},
	{input: 200000},
}

func BenchmarkSequence(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				returnNthValue(v.input)
			}
		})
	}
}
