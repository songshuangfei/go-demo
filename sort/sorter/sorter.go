package main

import (
	"flag"
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
)

import (
	"go-demo/sort/algorithms/bubblesort"
	"go-demo/sort/algorithms/qsort"
)

var infile *string = flag.String("i","infile","File contains values for sorting")
var outfile *string = flag.String("o","outfile","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","sort algorithm")

func readValues(infile string)(values []int,err error){//返回值初始化为默认值
	file,err := os.Open(infile)//只要:=有一个变量名是未定义的则不会报错，入这里err没报错
	if err != nil{
		fmt.Println("failed to open the input file ",infile)
		return
	}

	defer file.Close()//函数return前会执行这里的语句关闭文件
	
	br := bufio.NewReader(file)
	values = make([]int,0)
	
	for {
		line,isPrefix,err1 := br.ReadLine()
		if err1 !=nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line.seems unexpected")
			return
		}
		str := string(line)
		value,err1 := strconv.Atoi(str)//字符串转整形
		if err != nil{
			err = err1
			return
		}
		values = append(values,value)
	}
	return
}

func writeValues(values []int,outfile string) error{
	file,err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ",outfile)
		return err
	}

	defer file.Close()

	for _,value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main(){
	flag.Parse()//解析命令
	if infile != nil{
		fmt.Println("infile = ",*infile,"outfile = ",*outfile,"algorithm = ",*algorithm)
	}

	values,err := readValues(*infile)//读取文件中的值

	if err == nil {
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)//数组和数组切片不一样，数组是传值，而切片是传入一个原生指针，函数内修改后外面也会变
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm ",*algorithm," unsupported")
		}
		writeValues(values,*outfile)
	}else{
		fmt.Println(err)
	}

}