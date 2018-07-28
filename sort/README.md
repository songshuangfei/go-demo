## 排序
一个读取文件中内容排序并输出到新文件的程序

* 项目放至 %GOPATH%/src 下
* `cd ./sorter`
* `go build` 编译
* 执行命令排序 `./sorter -i infile.dat -o outfile.dat -a qsort`(./sorter是编译好的二进制文件 infile.dat 是未排序的数据，outfile.dat 是数据输出文件，qsort是排序算法，只有qsort和bubblesort两种算法)
