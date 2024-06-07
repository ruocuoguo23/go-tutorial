package main

import "fmt"

func main() {
	// 声明一个切片
	var slice []int
	fmt.Println("Empty slice:", slice)

	// 初始化切片
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("Initialized slice:", slice)

	// 访问和修改切片元素
	slice[0] = 10
	fmt.Println("Modified slice:", slice)

	// 切片长度和容量
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	// 切片是引用类型
	sliceCopy := slice
	sliceCopy[0] = 100
	fmt.Println("Original slice:", slice)
	fmt.Println("Copied slice:", sliceCopy)

	// 使用内置函数 append 动态增加元素
	slice = append(slice, 6)
	fmt.Println("Appended slice:", slice)

	// 基于数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArr := arr[1:4]
	fmt.Println("Slice from array:", sliceFromArr)

	// 基于切片创建切片
	newSlice := slice[1:3]
	fmt.Println("New slice from slice:", newSlice)

	slice = make([]int, 5)
	fmt.Println("Initialized slice with make:", slice)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	// append to slice
	slice = append(slice, 1, 2, 3)
	fmt.Println("Appended slice:", slice)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))
}
