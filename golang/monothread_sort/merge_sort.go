package monothread_sort

type mergeSort struct{}

func NewMergeSort() *mergeSort {
	return &mergeSort{}
}

func (ss *mergeSort) Sort(array []int) {
	ss.mergeSort(array, 0, len(array)-1)
}

func (ss *mergeSort) mergeSort(array []int, low int, high int) {
	if low < high {
		mid := low + (high-low)/2

		ss.mergeSort(array, low, mid)
		ss.mergeSort(array, mid+1, high)

		ss.merge(array, low, mid, high)
	}
}

func (ss *mergeSort) merge(array []int, low int, mid int, high int) {
	var i, j, k int
	n1 := mid - low + 1
	n2 := high - mid

	left := make([]int, n1)
	right := make([]int, n2)

	for i = 0; i < n1; i++ {
		left[i] = array[low+i]
	}

	for i = 0; i < n2; i++ {
		right[i] = array[mid+1+i]
	}

	i = 0
	j = 0
	k = low
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		array[k] = left[i]
		k++
		i++
	}

	for j < n2 {
		array[k] = right[j]
		k++
		j++
	}
}
