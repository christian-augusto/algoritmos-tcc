class QuickSort:
    def sort(self, array: list):
        self.quickSort(array, 0, len(array)-1)

    def quickSort(self, array: list, low: int, high: int):
        if low < high:
            pi = self.partition(array, low, high)

            self.quickSort(array, low, pi-1)
            self.quickSort(array, pi+1, high)

    def partition(self, array: list, low: int, high: int):
        pivot = array[high]

        i = low - 1
        j = low
        while j <= high-1:
            if array[j] < pivot:
                i += 1
                array[i], array[j] = array[j], array[i]

            j += 1

        array[i+1], array[high] = array[high], array[i+1]

        return i + 1
