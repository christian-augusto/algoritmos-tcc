package monothread_sort

type selectionSort struct {
	array []int
}

func NewSelectionSort(array []int) *selectionSort {
	return &selectionSort{
		array: array,
	}
}

func (ss *selectionSort) Sort() {
	for i := int(0); i < int(len(ss.array)); i++ {
		index := i

		for j := i + 1; j < int(len(ss.array)); j++ {
			if ss.array[j] < ss.array[index] {
				index = j
			}
		}

		if index != i {
			ss.array[i], ss.array[index] = ss.array[index], ss.array[i]
		}
	}
}
