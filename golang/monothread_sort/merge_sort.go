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

func (ss *mergeSort) mergeSort(arr []int64, l int64, r int64) {
	if l < r {
		m := l + (r-l)/2

		ss.mergeSort(arr, l, m)
		ss.mergeSort(arr, m+1, r)

		ss.merge(arr, l, m, r)
	}
}

func (ss *mergeSort) merge(arr []int64, l int64, m int64, r int64) {
	var i, j, k int64
	n1 := m - l + 1
	n2 := r - m

	L := make([]int64, n1)
	R := make([]int64, n2)

	for i = 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for j = 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i = 0
	j = 0
	k = l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
