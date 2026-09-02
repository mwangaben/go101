package interators

import "iter"

func Count(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func CountWithIndex(start, end int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i, i*2) {
				return
			}
		}
	}
}
