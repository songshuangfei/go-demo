package main

import (
	"fmt"
	"go-demo/MergeSort/pipeline"
)

func main() {
	p := pipeline.Merge(
		pipeline.Insort(pipeline.ArraySource(4, 7, 2, 14, 0)),
		pipeline.Insort(pipeline.ArraySource(8, 3, 2, 12, 3)))

	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}
