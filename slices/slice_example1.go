package main

import "fmt"

func remove(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s // return unchanged if index is out of range
	}
	return append(s[:index], s[index+1:]...)
}

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 7}
	arr1 = remove(arr1, 2) // removes element at index 2 (value 3)
	fmt.Println(arr1)      // Output: [1 2 4 5 7]

	fmt.Println("************slice example1************")
	arr := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(arr) // // [1 2 3 4 , 5]
	sl := arr[2:4]
	fmt.Println(sl)                                            // [3 4]
	fmt.Printf("len: %v, cap: %v, %v\n", len(sl), cap(sl), sl) // len: 2 cap: 3 [0 4]
	fmt.Printf("%p\n", &sl[0])

	fmt.Println("************slice example2************")
	sl[0] = 0
	fmt.Println(sl)                                            // [0 4]
	fmt.Println(arr)                                           // [1 2 0 4 5]
	fmt.Printf("len: %v, cap: %v, %v\n", len(sl), cap(sl), sl) // len: 2 cap: 3 [0 4]
	fmt.Printf("%p\n", &sl[0])

	fmt.Println("************slice example3************")
	sl = append(sl, 9)
	fmt.Println(sl)                                            // [0 4 9]
	fmt.Println(arr)                                           // [1 2 0 4 9]
	fmt.Printf("len: %v, cap: %v, %v\n", len(sl), cap(sl), sl) // len: 2 cap: 3 [0 4]
	fmt.Printf("%p\n", &sl[0])
	fmt.Printf("%p\n", &arr[0])

	fmt.Println("************slice example4************")

	sl = append(sl, 10)
	fmt.Println(sl)                                            // [0 4 9]
	fmt.Println(arr)                                           // [1 2 0 4 9]
	fmt.Printf("len: %v, cap: %v, %v\n", len(sl), cap(sl), sl) // len: 2 cap: 3 [0 4]
	fmt.Printf("%p\n", &sl[0])
	fmt.Printf("%p\n", &arr[0])

}
