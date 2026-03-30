func maxArea(height []int) int {
    
    maxArea := 0
    if height[0] <= height[1] {
        maxArea = height[0]
    } else {
        maxArea = height[1]
    }

    for i, j := 0, len(height)-1; i < j; {
        
        minHeight := 0
        if height[i] < height[j] {
            minHeight = height[i]
            i++
        } else {
            minHeight = height[j]
            j--
        }
        lenght := j - i + 1
        area := minHeight * lenght
        if area > maxArea{
            maxArea = area
        }
    }

    return maxArea
}