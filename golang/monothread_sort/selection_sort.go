package monothread_sort

type selectionSort struct{}

func NewSelectionSort() *selectionSort {
	return &selectionSort{}
}

func (ss *selectionSort) Sort(array []int) {
	for i := 0; i < len(array); i++ {
		index := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[index] {
				index = j
			}
		}

		if index != i {
			array[i], array[index] = array[index], array[i]
		}
	}
}
