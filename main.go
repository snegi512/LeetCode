package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var stt, sed []int
	var t1, t2 *int
	i1, i2 := 0, 0
	c := (len(nums1) + len(nums2)) % 2
	center := (len(nums1)+len(nums2))/2 + c

	switch {
	case len(nums1) == 0:
		fmt.Println(1)
		stt, sed = nums2, nums1
		t1, t2 = &i2, &i1
	case len(nums2) == 0:
		fmt.Println(2)
		stt, sed = nums1, nums2
		t1, t2 = &i1, &i2
	case nums1[0] <= nums2[0]:
		fmt.Println(3)
		stt, sed = nums1, nums2
		t1, t2 = &i1, &i2
	case nums1[0] > nums2[0]:
		fmt.Println(4)
		stt, sed = nums2, nums1
		t1, t2 = &i2, &i1
	}

	for i := 0; i < center; i++ {

		if *t1 >= len(stt) {
			stt, sed = sed, stt
			t1, t2 = t2, t1
			continue
		}
		if stt[*t1] > sed[*t2] {
			stt, sed = sed, stt
			t1, t2 = t2, t1
		}
		*t1++
	}
	if c == 0 && len(sed) != 0 {
		return float64((stt[*t1-1] + sed[*t2])) / 2.0
	} else {
		return float64(stt[*t1-1])
	}

	// if nums1
	// struct = nums1
	// for i:= 0; i < (len(nums1)+len(nums2))/2;{

	// }
	return 0
}
func main() {
	fmt.Print(findMedianSortedArrays([]int{1, 2}, []int{}))
}
