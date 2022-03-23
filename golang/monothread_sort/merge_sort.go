package monothread_sort

type mergeSort struct {
	array []int64
}

func NewMergeSort(array []int64) *mergeSort {
	return &mergeSort{
		array: array,
	}
}

func (ss *mergeSort) Sort() {
	ss.mergeSort(ss.array, 0, int64(len(ss.array)-1))
}

func (ss *mergeSort) mergeSort(arr []int64, low int64, high int64) {
	if low < high {
		mid := low + (high-low)/2

		ss.mergeSort(arr, low, mid)
		ss.mergeSort(arr, mid+1, high)

		ss.merge(arr, low, mid, high)
	}
}

func (ss *mergeSort) merge(arr []int64, low int64, mid int64, high int64) {
	var i, j, k int64
	n1 := mid - low + 1
	n2 := high - mid

	left := make([]int64, n1)
	right := make([]int64, n2)

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
