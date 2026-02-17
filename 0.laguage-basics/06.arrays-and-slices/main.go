package main

import (
	"fmt"
	"slices"
)

func main() {

	/** Arrays are fixed-size collections of elements of the same type.
	* Initialize an array with values
	* int -> all elements are initialized to zero (0)
	* bool -> all elements are initialized to false
	* string -> all elements are initialized to empty string ("")
	 */

	var nums [4]int
	fmt.Println("Array: ", nums)

	array_size := len(nums)
	fmt.Println("Array Size: ", array_size)

	nums[0] = 10
	nums[3] = 40
	fmt.Println("3rd Element:", nums[3])

	names := [4]string{"John", "Jane", "Jack", "Jill"}
	fmt.Println("String Array: ", names)

	array_2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("2D Array: ", array_2d)

	/**
	 	* Slices are dynamic collections of elements of the same type.
	 	* They are more flexible than arrays and can grow or shrink in size.
		* Uninitialized slices are nil and have a length of zero.
	*/

	var slice []int
	fmt.Println("Slice: ", slice, "Length: ", len(slice), "Nill: ", slice == nil)

	// With make
	slice = make([]int, 2, 4)

	fmt.Println("Slice: ", slice, "Length: ", len(slice), "Capacity: ", cap(slice), "Nill: ", slice == nil)

	slice[0] = 0
	slice[1] = 1
	// slice[2] = 2 // Will panic: runtime error: index out of range [2] with length 2
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println("Slice: ", slice, "Length: ", len(slice), "Capacity: ", cap(slice), "Nill: ", slice == nil)

	// Note - Capacity vs Length: Capacity is the maximum number of elements that a slice can hold, while length is the current number of elements in the slice. Even though capacity can be increased dynamically using the append function.

	numbers := make([]int, len(slice), cap(slice))
	copy(numbers, slice)
	fmt.Println("Numbers - Copied from slice: ", numbers)

	// Equality Check
	is_equal := slices.Equal(slice, numbers)
	fmt.Println("Equality Check: ", is_equal)

	// Other Operations

	// Slice Length Operations
	numbers = numbers[:5]
	fmt.Println("Numbers: ", numbers, "Length: ", len(numbers))

	// Delete Element at Index 2
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("Numbers: ", numbers, "Length: ", len(numbers))

	// Insert Element at Index 2
	numbers = append(numbers[:2], append([]int{0}, numbers[2:]...)...)
	fmt.Println("Numbers: ", numbers, "Length: ", len(numbers))

	// Update Element at Index 2
	numbers[2] = 2
	fmt.Println("Numbers: ", numbers, "Length: ", len(numbers))

}
