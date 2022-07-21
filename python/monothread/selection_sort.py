class SelectionSort:
    def sort(self, array: list):
        for i in range(0, len(array)):
            index = i

            for j in range(i+1, len(array)):
                if array[j] < array[index]:
                    index = j

            if index != i:
                array[i], array[index] = array[index], array[i]
