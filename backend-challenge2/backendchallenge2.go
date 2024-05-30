package backendchallenge2

func Decode(encoded string) []int {
	n := len(encoded)
	if n == 0 {
		return nil
	}
	var array []int
	var result []int
	for i := 0; i < n; i++ {
		if encoded[i] == 'L' {
			array = append(array, 1, 0)
		} else if encoded[i] == 'R' {
			array = append(array, 0, 1)
		} else if encoded[i] == '=' {
			array = append(array, 0, 0)
		}
	}

	result = append(result, array[0])

	for i := 1; i < len(array)-2; i += 2 {
		sum := array[i] + array[i+1]
		result = append(result, sum)
	}

	result = append(result, array[9])

	for i := 0; i < len(result)-1; i++ {
		switch encoded[i] {
		case 'L':
			if result[i] == result[i+1] {
				result[i] = result[i] + 1
			} else if result[i] < result[i+1] {
				result[i] = result[i+1] + 1
			}
		case 'R':
			if result[i] == result[i+1] {
				result[i+1] = result[i+1] + 1
			} else if result[i] > result[i+1] {
				result[i+1] = result[i] + 1
			}
		case '=':
			result[i+1] = result[i]
		}
	}

	for i := 0; i < len(result)-1; i++ {
		if encoded[i] == '=' {
			if result[i] > result[i+1] {
				result[i+1] = result[i]
			} else if result[i+1] > result[i] {
				result[i] = result[i+1]
			}
		}
	}

	return result
}

func Sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}
