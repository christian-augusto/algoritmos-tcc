import utils
from constants import (
    LOGS, TOTAL_LOGS, ARRAY_SIZES, RANDOM_INT_MIN, RANDOM_INT_MAX
)
from monothread.selection_sort import SelectionSort as MonothreadSelectionSort
from monothread.merge_sort import MergeSort as MonothreadMergeSort
from monothread.quick_sort import QuickSort as MonothreadQuickSort


def main():
    for size in ARRAY_SIZES:
        print(f"size: {size}")

        sort(size)

        separator()


def sort(size: int):
    array_1 = utils.generate_random_int_array(
        size, RANDOM_INT_MIN, RANDOM_INT_MAX
    )
    array_2 = utils.copy_array(array_1)
    array_3 = utils.copy_array(array_1)

    if LOGS and TOTAL_LOGS:
        print(array_1)
        print(array_2)
        print(array_3)

    # selection sort

    selection_sort = MonothreadSelectionSort()

    start_1 = utils.current_milli_time()

    selection_sort.sort(array_1)

    elapsed_1 = utils.current_milli_time() - start_1

    # merge sort

    merge_sort = MonothreadMergeSort()

    start_2 = utils.current_milli_time()

    merge_sort.sort(array_2)

    elapsed_2 = utils.current_milli_time() - start_2

    # quick sort

    quick_sort = MonothreadQuickSort()

    start_3 = utils.current_milli_time()

    quick_sort.sort(array_3)

    elapsed_3 = utils.current_milli_time() - start_3

    if LOGS:
        print(f"Selection sort execution time: {elapsed_1} ms")
        print(f"Merge sort execution time: {elapsed_2} ms")
        print(f"Quick sort execution time: {elapsed_3} ms")

        if TOTAL_LOGS:
            print(array_1)
            print(array_2)
            print(array_3)


def separator():
    print()


if __name__ == "__main__":
    main()
