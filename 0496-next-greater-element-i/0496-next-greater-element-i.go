func nextGreaterElement(nums1 []int, nums2 []int) []int {
    
    m := make(map[int]int)

    for i := 0; i < len(nums2); i++{
        for j := i+1; j < len(nums2); j++{
            if nums2[i] < nums2[j] {
                m[nums2[i]] = nums2[j]
                break
            }
        }
        if _, exist := m[nums2[i]]; !exist{
            m[nums2[i]] = -1
        }
    }

    ans := make([]int, len(nums1))
    for i := range nums1 {
        ans[i] = m[nums1[i]]
    }

    return ans
}