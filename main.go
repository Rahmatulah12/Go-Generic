package main

import "fmt"

// it's used struct
func Sum(numbers []int) int {
	var total int

	for _, e := range numbers {
		total += e
	}
	return total
}

// it's used go generic
// it's can be used as comparable data types
func SumGeneric[T int | float64 | float32](numbers []T) T {
	var total T

	for _, e := range numbers {
		total += e
	}

	return total
}

/*
Keywords Comparable
*/
func SumNumberComparable[K comparable, T int | float64 | float32](m map[K]T) T {
	var s T
	for _, e := range m {
		s += e
	}
	return s
}

// Generic Type Constrant
type Number interface {
	int64 | float64
}

func SumNumberGenericTypeContrant[K comparable, T Number](m map[K]T) T {
	var n T
	for _, e := range m {
		n += e
	}
	return n
}

// generic pada pendefinisian struct
type UserValue[T Number] struct {
	Name    string
	Numbers []T
}

func main() {
	total := Sum([]int{10, 9, 8, 7, 6})
	fmt.Println("total:", total)

	total2 := SumGeneric([]int{10, 9, 8, 7, 6})
	fmt.Println("total:", total2)

	totalFloat := SumGeneric([]float64{30.25, 10.01, 2.6, 0.56})
	fmt.Println(totalFloat)

	// call comparable functions
	ints := map[string]int{
		"first":  1,
		"second": 2,
	}

	ints64 := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floats := map[string]float64{
		"one": 1.2,
		"two": 2.2,
	}

	totalComparableInt := SumNumberComparable(ints)
	fmt.Println(totalComparableInt)
	fmt.Println("Ini Generic Type Constrant :", SumNumberGenericTypeContrant(ints64))

	totalComparableFloat := SumNumberComparable(floats)
	fmt.Println(totalComparableFloat)
	fmt.Println("ini generic type contrant :", SumNumberGenericTypeContrant(floats))

	// generic pada struct
	user1 := UserValue[int64]{Name: "Rahmatulah Sidik", Numbers: []int64{1, 2, 3, 4, 5, 6}}
	fmt.Println(user1)

	var user2 UserValue[float64]
	user2.Name = "Rahmatulah"
	user2.Numbers = []float64{2.3, 3.4, 4.5, 5.6, 6.7}
	fmt.Println(user2)
}
