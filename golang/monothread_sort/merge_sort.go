package monothread_sort

type mergeSort struct {
	array []int
}

func NewMergeSort(array []int) *mergeSort {
	return &mergeSort{
		array: array,
	}
}

func (ss *mergeSort) Sort() {
	ss.mergeSort(ss.array, 0, int(len(ss.array)-1))
}

func (ss *mergeSort) mergeSort(arr []int, low int, high int) {
	if low < high {
		mid := low + (high-low)/2

		ss.mergeSort(arr, low, mid)
		ss.mergeSort(arr, mid+1, high)

		ss.merge(arr, low, mid, high)
	}
}

func (ss *mergeSort) merge(arr []int, low int, mid int, high int) {
	var i, j, k int
	n1 := mid - low + 1
	n2 := high - mid

	left := make([]int, n1)
	right := make([]int, n2)

	for i = 0; i < n1; i++ {
		left[i] = arr[low+i]
	}

	for j = 0; j < n2; j++ {
		right[j] = arr[mid+1+j]
	}

	i = 0
	j = 0
	k = low
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = left[i]
		k++
		i++
	}

	for j < n2 {
		arr[k] = right[j]
		k++
		j++
	}
}
