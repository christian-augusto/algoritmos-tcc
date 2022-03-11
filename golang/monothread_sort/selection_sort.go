package monothread_sort

type selectionSort struct {
	array []int64
}

func NewSelectionSort(array []int64) *selectionSort {
	return &selectionSort{
		array: array,
	}
}

func (ss *selectionSort) Sort() {
	for i := int64(0); i < int64(len(ss.array)); i++ {
		index := i

		for j := i + 1; j < int64(len(ss.array)); j++ {
			if ss.array[j] < ss.array[index] {
				index = j
			}
		}

		if index != i {
			ss.array[i], ss.array[index] = ss.array[index], ss.array[i]
		}
	}
}
