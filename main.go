package main

import (
	_ "github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"log"
	"strconv"
	"math/rand"
	"time"
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)
var (
	newFile *os.File
	err     error
)
func main() {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "Time")
	xlsx.SetCellValue("Sheet1", "B1", "Array")
for i:=0;i<500;i++{
	arr:=make([]int,10000000)
	s:=strconv.Itoa(i)
	filename:=s+".txt"
	xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2),filename)
	newFile, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	for i := 0; i < 10000000; i++ {
		swap:=rand.Intn(10000000)
		_, err=newFile.WriteString(strconv.Itoa(swap)+"\n")
		arr[i]=swap
	}
	newFile.Close()
	start:=time.Now()
	mergesort_iterative_cpu(arr,10000000)
	end:=time.Now()
	fmt.Println(end.Nanosecond()-start.Nanosecond())
	xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2),end.Nanosecond()-start.Nanosecond())
}
	err := xlsx.SaveAs("./Benchmark_cpu.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
func mergesort_iterative_cpu(arr[]int,size int){
	temparr:= make([]int,size)
		var right int
		var rend int
		var i int
		var j int
		var m int

		for k:= 1; k < size; k *= 2 {
	//at each partition size, sort and merge
			for  left := 0; left + k < size; left += k*2 {
	//store the start of the right partition and its end
	right = left + k
	rend = right + k

	//if the partitions are uneven, readjust the end
	if rend > size{
		rend = size
		}
			m = left
			i = left
			j = right

	//merge
	for i < right && j < rend {
		if arr[i] <= arr[j] {
			temparr[m] = arr[i]
			i++
	} else {
			temparr[m] = arr[j]
			j++
	}
		m++
	}
	for i < right {
		temparr[m] = arr[i]
		i++
		m++
	}
	for j < rend {
		temparr[m] = arr[j]
		j++
		m++
	}
	//copy from temp array into initial array
	for m = left; m < rend; m++ {
		arr[m] = temparr[m]
	}
		}
	}
}