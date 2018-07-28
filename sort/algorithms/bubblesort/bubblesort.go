package bubblesort

func BubbleSort(values []int){
	flag := true //flag是为了在顺序情况下节约循环次数

	for i := 0;i < len(values)-1;i ++ {
		flag = true
		for j := 0;j < len(values)-i-1;j ++ {
			if values[j] > values[j+1] {
				values[j],values[j+1] = values[j+1],values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}