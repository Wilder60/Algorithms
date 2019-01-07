package Sorts

//------------------------------------------INSERTION-SORT--------------------------------------------------------------
func InsertionSort(_Array []int) []int {
	for i := 1; i < len(_Array); i++ {

		var key int = _Array[i]
		var j int = i - 1

		for j > -1 && _Array[j] > key {
			_Array[j+1] = _Array[j]
			j--
		}
		_Array[j+1] = key
	}
	return _Array
}

//------------------------------------------BUBBLE-SORT-----------------------------------------------------------------
func BubbleSort(_Array []int) []int {
	for i := 0; i < len(_Array)-1; i++ {
		for j := 0; j < len(_Array)-i-1; j++ {
			if _Array[j] > _Array[j+1] {
				_Array[j], _Array[j+1] = _Array[j+1], _Array[j]
			}
		}
	}
	return _Array
}

//------------------------------------------SELECTION-SORT-------------------------------------------------------------
func SelectionSort(_Array []int) []int {

	for i := 0; i < len(_Array); i++ {
		lowestIndex := i
		for j := i + 1; j < len(_Array); j++ {
			if _Array[lowestIndex] > _Array[j] {
				lowestIndex = j
			}
		}
		_Array[lowestIndex], _Array[i] = _Array[i], _Array[lowestIndex]
	}
	return _Array
}

//------------------------------------------QUICK-SORT------------------------------------------------------------------
func QuickSort(_Array []int) []int {
	return quickSortRecursive(_Array, 0, len(_Array)-1)
}

func quickSortRecursive(_Array []int, low int, high int) []int {
	if low < high {
		pivot := parition(_Array, low, high)
		quickSortRecursive(_Array, low, pivot-1)
		quickSortRecursive(_Array, pivot+1, high)
	}
	return _Array
}

func parition(_Array []int, low int, high int) int {
	pivot := _Array[high]

	i := low - 1

	for j := low; j <= high-1; j++ {

		if _Array[j] <= pivot {
			i++
			_Array[i], _Array[j] = _Array[j], _Array[i]
		}

	}
	_Array[i+1], _Array[high] = _Array[high], _Array[i+1]
	return (i + 1)
}
