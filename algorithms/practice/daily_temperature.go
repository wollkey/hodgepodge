package practice

// DailyTemperatures
// O(n^2)
func DailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	answer := make([]int, length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if temperatures[i] < temperatures[j] {
				answer[i] = j - i
				break
			}
		}
	}

	return answer
}

// DailyTemperatureOptimized
// O(n)
func DailyTemperatureOptimized(temperatures []int) []int {
	length := len(temperatures)
	answer := make([]int, length)
	var stack []int

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			answer[index] = i - index
		}

		stack = append(stack, i)
	}

	return answer
}
