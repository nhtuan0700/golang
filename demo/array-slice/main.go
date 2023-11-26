package main

import (
	"fmt"
)

func AddOneToEachElement(slice []int) {
	for i := range slice {
		slice[i]++
	}
}

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	fmt.Println("slice", slice, len(slice), slice[index+1:], len(slice[index+1:]))
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	fmt.Println("slice", slice)

	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// Append appends the items to the slice.
// First version: just loop calling Extend.
func Append(slice []int, items ...int) []int {
	// for _, item := range items {
	// 	slice = Extend(slice, item)
	// }
	// return slice
	n := len(slice)
	total := len(slice) + len(items)

	if total > cap(slice) {
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], items)
	return slice
}

func main() {
	// slice := buffer[10:20]
	// for i := 0; i < len(slice); i++ {
	// 	slice[i] = byte(i)
	// }
	// fmt.Println("before", slice)
	// AddOneToEachElement(slice)
	// fmt.Println("after", slice)

	// var buffer [256]byte
	// AddOneToEachElement(slice)
	// fmt.Println("after", slice)

	// Shift array
	// slice := make([]int, 10, 15) // Note capacity > length: room to add element.
	// for i := range slice {
	// 	slice[i] = i
	// }
	// fmt.Println(cap(slice))
	// newSlice := make([]int, len(slice), 2*cap(slice))
	// fmt.Println(copy(newSlice, slice))
	// slice = newSlice
	// fmt.Println(cap(slice))

	// Append
	// slice := make([]int, 0, 5)
	// for i := 0; i < 10; i++ {
	// 	slice = Extend(slice, i)
	// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	// 	fmt.Println("address of 0th element:", &slice[0])
	// }
	// slice := []int{0, 1, 2, 3, 4}
	// fmt.Println(slice)
	// slice = append(slice, 5, 6, 7, 8)
	// fmt.Println(slice)
	// str := string(slice)
	// slice := []byte(usr)

	// var buffer [10]byte
	// fmt.Println(buffer)
	// slice := make([]int, 10, 15)
	// copy(slice, slice)
	// for i := range slice {
	// 	slice[i] = int(buffer[i])
	// }

	// slice_cp := slice
	// var slice_cp [10]int
	// for i := range slice {
	// 	slice_cp[i] = int(slice[i])
	// }
	// var slice_cp [10]int
	// fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

}
