from constants import THREADS_NUMBER
import threading


class QuickSort:
    def sort(self, array: list):
        self.threads_count = 0

        self.quickSort(array, 0, len(array)-1)

    def quickSort(self, array: list, low: int, high: int):
        if low < high:
            pi = self.partition(array, low, high)

            threads = []

            if self.threads_count < THREADS_NUMBER:
                self.threads_count += 1

                thread = threading.Thread(target=self.quickSort,
                                          args=(array, low, pi-1))
                threads.append(thread)
                thread.start()
            else:
                self.quickSort(array, low, pi-1)

            if self.threads_count < THREADS_NUMBER:
                self.threads_count += 1

                thread = threading.Thread(target=self.quickSort,
                                          args=(array, pi+1, high))
                threads.append(thread)
                thread.start()
            else:
                self.quickSort(array, pi+1, high)

            for thread in threads:
                thread.join()

            self.threads_count -= len(threads)

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
