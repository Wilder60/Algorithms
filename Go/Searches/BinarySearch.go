package searches

func BinarySearch(target int, arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	return binarySearchHelper(target, 0, len(arr), arr)
}

func binarySearchHelper(target, low, high int, arr []int) bool {
	mid := (low + high) / 2
	if low >= high {
		return false
	}
	if arr[mid] < target {
		return binarySearchHelper(target, mid, high, arr)
	} else if arr[mid] > target {
		return binarySearchHelper(target, low, mid, arr)
	}
	return true
}
