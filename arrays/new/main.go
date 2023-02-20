package main

import (
	"fmt"
)

func main() {
	var nums1 [3]string
	nums1[0] = "First"
	nums1[1] = "Second"
	nums1[2] = "Third"
	fmt.Println(nums1)

	nums2 := [3]string{"First", "Second", "Third"}
	fmt.Println(nums2)
}
