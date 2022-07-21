import utils
from contants import LOGS, TOTAL_LOGS, ARRAY_SIZES, RANDOM_INT_MIN, \
    RANDOM_INT_MAX
from monothread.selection_sort import SelectionSort as MonothreadSelectionSort


def main():
    for size in ARRAY_SIZES:
        print(f"size: {size}")

        sort(size)

        separator()


def sort(size: int):
    array_1 = utils.generate_random_int_array(
        size, RANDOM_INT_MIN, RANDOM_INT_MAX
    )

    if LOGS and TOTAL_LOGS:
        print(array_1)

    # selection sort

    selection_sort = MonothreadSelectionSort()

    start_1 = utils.current_milli_time()

    selection_sort.sort(array_1)

    elapsed_1 = utils.current_milli_time() - start_1

    if LOGS:
        print(f"Selection sort execution time: {elapsed_1} ms")

        if TOTAL_LOGS:
            print(array_1)


def separator():
	print("\n\n")



if __name__ == "__main__":
    main()
