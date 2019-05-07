func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    
    for i := 0; i < len(nums); i++ {
        if _, ok := m[nums[i]]; ok {
            return true
        }
        m[nums[i]] = struct{}{}
    }
    
    return false
}
