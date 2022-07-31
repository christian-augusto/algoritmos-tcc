from constants import THREADS_NUMBER
import threading


class MergeSort:
    def sort(self, array: list):
        threads = []
        arrayLen = len(array)

        for i in range(0, THREADS_NUMBER):
            low = i * (arrayLen // 4)
            high = (i + 1) * (arrayLen // 4) - 1
            mid = low + (high-low)//2

            threads.append(
                threading.Thread(target=self.sort_thread,
                                 args=(array, low, mid, high))
            )
            threads[i].start()

        for thread in threads:
            thread.join()

        self.merge(array, 0, (arrayLen//2-1)//2, arrayLen//2-1)
        self.merge(array, arrayLen//2, arrayLen//2 +
                   (arrayLen-1-arrayLen//2)//2, arrayLen-1)
        self.merge(array, 0, (arrayLen-1)//2, arrayLen-1)

    def sort_thread(self, array: list, low: int, mid: int, high: int):
        self.mergeSort(array, low, mid)
        self.mergeSort(array, mid + 1, high)

        self.merge(array, low, mid, high)

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
