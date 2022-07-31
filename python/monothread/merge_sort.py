class MergeSort:
    def sort(self, array: list):
        self.mergeSort(array, 0, len(array)-1)

    def mergeSort(self, array: list, low: int, high: int):
        if low < high:
            mid = low + (high-low)//2

            self.mergeSort(array, low, mid)
            self.mergeSort(array, mid+1, high)

            self.merge(array, low, mid, high)

    def merge(self, array: list, low: int, mid: int, high: int):
        n1 = mid - low + 1
        n2 = high - mid

        left = []
        right = []

        for i in range(0, n1):
            left.append(array[low+i])

        for i in range(0, n2):
            right.append(array[mid+1+i])

        i = 0
        j = 0
        k = low
        while i < n1 and j < n2:
            if left[i] <= right[j]:
                array[k] = left[i]
                i += 1
            else:
                array[k] = right[j]
                j += 1

            k += 1

        while i < n1:
            array[k] = left[i]
            k += 1
            i += 1

        while j < n2:
            array[k] = right[j]
            k += 1
            j += 1
