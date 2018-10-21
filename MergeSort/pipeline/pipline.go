package pipeline

import (
	"sort"
)

//ArraySource ...
func ArraySource(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

//Insort ...
func Insort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for n := range in {
			a = append(a, n)
		}
		sort.Ints(a)
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//Merge ..
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok1 || (ok2 && v2 <= v1) {
				out <- v2
				v2, ok2 = <-in2
			} else {
				out <- v1
				v1, ok1 = <-in1
			}
		}
		close(out)
	}()
	return out
}
