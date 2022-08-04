import java.util.Arrays;

import monothread_sort.SelectionSort;

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

        if (Constants.LOGS && Constants.TOTAL_LOGS) {
            System.out.println(Arrays.toString(array_1));
            System.out.println(Arrays.toString(array_2));
        }

        // selection sort

        SelectionSort selectionSort = new SelectionSort();

        long start_1 = Utils.getCurrentUnixTime();

        selectionSort.sort(array_1);

        long elapsed_1 = Utils.getCurrentUnixTime() - start_1;

        if (Constants.LOGS) {
            System.out.printf("Selection sort execution time: %d ms\n", elapsed_1);

            if (Constants.TOTAL_LOGS) {
                System.out.println(Arrays.toString(array_1));
                System.out.println(Arrays.toString(array_2));
            }
        }
    }

    private static void separator() {
        System.out.println();
    }
}
