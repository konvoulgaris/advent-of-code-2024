package utils

func MoveArrayElement(slice []int, from, to int) []int {
	if from == to {
		return slice
	}

	element := slice[from]
	newSlice := append(slice[:from], slice[from+1:]...)

	if to > from {
		to--
	}

	newSlice = append(newSlice[:to], append([]int{element}, newSlice[to:]...)...)

	return newSlice
}

func FindIndex(slice []int, target int) int {
	for i, val := range slice {
		if val == target {
			return i
		}
	}

	return -1
}
