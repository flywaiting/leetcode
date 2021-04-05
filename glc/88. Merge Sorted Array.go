package glc

func Merge(n1, n2 []int, m, n int) {
	merge(n1, m, n2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
