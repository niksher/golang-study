package lesson

import (
	"sort"
)

func Lesson4(A []int) int {
	sort.Ints(A)
	oldElement := 0
	for _, a := range A {
		if oldElement == 0 {
			oldElement = a
			continue
		}
		if a != oldElement+1 {
			return oldElement + 1
		}
		oldElement = a
	}

	return 0
}

func Lesson3(A []int) int {
	sort.Ints(A)
	oldElement := 0
	for _, a := range A {
		if oldElement == 0 {
			oldElement = a
			continue
		}
		if a != oldElement+1 {
			return 0
		}
		oldElement = a
	}

	return 1
}

func Lesson2(A []int) int {
	sort.Ints(A)
	oldElement := 0
	oldElementIsPair := false
	for _, currentElement := range A {
		if oldElement == 0 {
			oldElement = currentElement
			continue
		}

		if currentElement != oldElement && !oldElementIsPair {
			return oldElement
		}

		if currentElement == oldElement && !oldElementIsPair {
			oldElementIsPair = true
		} else {
			oldElementIsPair = false
		}

		oldElement = currentElement
	}

	if !oldElementIsPair {
		return oldElement
	}

	return 0
}

func Lesson1(A []int, K int) []int {
	l := len(A)
	B := make([]int, l)
	for i := 0; i < l; i++ {
		moved := (i + K) % l
		B[moved] = A[i]
	}

	return B
}
