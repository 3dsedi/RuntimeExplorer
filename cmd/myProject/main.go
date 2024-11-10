package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func GenerateRandomArray(size int, max int) []int {
	array := make([]int, size)
	for i := range array {
		array[i] = rand.Intn(max)
	}
	return array
}

func InsertRandom(arr []int, value int) []int {
	pos := rand.Intn(len(arr) + 1)
	arr = append(arr[:pos], append([]int{value}, arr[pos:]...)...)
	return arr
}

func InsertFront(arr []int, value int) []int {
	arr = append([]int{value}, arr...)
	return arr
}

func InsertBack(arr []int, value int) []int {
	arr = append(arr, value)
	return arr
}

func RemoveFront(arr []int) []int {
	arr = arr[1:]
	return arr
}

func RemoveBack(arr []int) []int {
	pos := len(arr) - 1
	arr = arr[:pos]
	return arr
}

func SearchUnsorted(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func SearchSorted(arr []int, value int) int {

	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}

func main() {
	start := time.Now()
	array := GenerateRandomArray(500000000, 20000000000)
	fmt.Println("create the array", time.Since(start))

	start = time.Now()
	array = InsertRandom(array, 50)
	fmt.Println("Insert random time", time.Since(start))

	start = time.Now()
	array = InsertFront(array, 50)
	fmt.Println("insert in front time", time.Since(start))

	start = time.Now()
	array = InsertBack(array, 50)
	fmt.Println("insert in back time", time.Since(start))

	start = time.Now()
	array = RemoveFront(array)
	fmt.Println("Remove front time", time.Since(start))

	start = time.Now()
	array = RemoveBack(array)
	fmt.Println("remove back time", time.Since(start))

	start = time.Now()
	finde := SearchUnsorted(array, 50)
	fmt.Printf("value is item %d in unsorted", finde)
	fmt.Println("usorted time", time.Since(start))

	sort.Ints(array)
	start = time.Now()
	finde2 := SearchSorted(array, 50)
	fmt.Printf("value is item %d in sorted", finde2)
	fmt.Println("usorted time", time.Since(start))

}
