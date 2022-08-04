import java.util.Arrays;

import monothread_sort.SelectionSort;
import monothread_sort.MergeSort;
import monothread_sort.QuickSort;

/**
 * Main
 */
public class Main {
    public static void main(String[] args) {
        for (int size : Constants.ARRAY_SIZES) {
            System.out.printf("size: %d\n", size);

            sort(size);

            separator();
        }
    }

    private static void sort(int size) {
        int[] array_1 = Utils.generateRandomIntArray(size, Constants.RANDOM_INT_MIN, Constants.RANDOM_INT_MAX);
        int[] array_2 = Utils.copyArray(array_1);
        int[] array_3 = Utils.copyArray(array_1);

        if (Constants.LOGS && Constants.TOTAL_LOGS) {
            System.out.println(Arrays.toString(array_1));
            System.out.println(Arrays.toString(array_2));
            System.out.println(Arrays.toString(array_3));
        }

        // selection sort

        SelectionSort selectionSort = new SelectionSort();

        long start_1 = Utils.getCurrentUnixTime();

        selectionSort.sort(array_1);

        long elapsed_1 = Utils.getCurrentUnixTime() - start_1;

        // merge sort

        MergeSort mergeSort = new MergeSort();

        long start_2 = Utils.getCurrentUnixTime();

        mergeSort.sort(array_2);

        long elapsed_2 = Utils.getCurrentUnixTime() - start_2;

        // quick sort

        QuickSort quickSort = new QuickSort();

        long start_3 = Utils.getCurrentUnixTime();

        quickSort.sort(array_3);

        long elapsed_3 = Utils.getCurrentUnixTime() - start_3;

        if (Constants.LOGS) {
            System.out.printf("Selection sort execution time: %d ms\n", elapsed_1);
            System.out.printf("Merge sort execution time: %d ms\n", elapsed_2);
            System.out.printf("Quick sort execution time: %d ms\n", elapsed_3);

            if (Constants.TOTAL_LOGS) {
                System.out.println(Arrays.toString(array_1));
                System.out.println(Arrays.toString(array_2));
                System.out.println(Arrays.toString(array_3));
            }
        }
    }

    private static void separator() {
        System.out.println();
    }
}
