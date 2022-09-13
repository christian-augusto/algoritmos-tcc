package algorithms;

import java.util.Arrays;
import algorithms.sort.MonothreadMergeSort;
import algorithms.sort.MonothreadQuickSort;
import algorithms.sort.MonothreadSelectionSort;
import algorithms.sort.MultithreadMergeSort;
import algorithms.sort.MultithreadQuickSort;

/**
 * Main
 */
public class Main {
    public static void main(String[] args) throws Exception {
        for (int size : Constants.ARRAY_SIZES) {
            System.out.printf("size: %d\n", size);

            sort(size);

            separator();
        }
    }

    private static void sort(int size) throws Exception {
        int[] array_1 = Utils.generateRandomIntArray(size, Constants.RANDOM_INT_MIN, Constants.RANDOM_INT_MAX);
        int[] array_2 = Utils.copyArray(array_1);
        int[] array_3 = Utils.copyArray(array_1);
        int[] array_4 = Utils.copyArray(array_1);
        int[] array_5 = Utils.copyArray(array_1);

        if (Constants.LOGS && Constants.TOTAL_LOGS) {
            System.out.println(Arrays.toString(array_1));
            System.out.println(Arrays.toString(array_2));
            System.out.println(Arrays.toString(array_3));
            System.out.println(Arrays.toString(array_4));
            System.out.println(Arrays.toString(array_5));
        }

        // selection sort

        MonothreadSelectionSort monothreadSelectionSort = new MonothreadSelectionSort(array_1);

        long start_1 = Utils.getCurrentUnixTime();

        monothreadSelectionSort.sort();

        long elapsed_1 = Utils.getCurrentUnixTime() - start_1;

        // merge sort

        MonothreadMergeSort monothreadMergeSort = new MonothreadMergeSort(array_2);

        long start_2 = Utils.getCurrentUnixTime();

        monothreadMergeSort.sort();

        long elapsed_2 = Utils.getCurrentUnixTime() - start_2;

        // quick sort

        MonothreadQuickSort monothreadQuickSort = new MonothreadQuickSort(array_3);

        long start_3 = Utils.getCurrentUnixTime();

        monothreadQuickSort.sort();

        long elapsed_3 = Utils.getCurrentUnixTime() - start_3;

        // multithread merge sort

        MultithreadMergeSort multithreadMergeSort = new MultithreadMergeSort(array_4);

        long start_4 = Utils.getCurrentUnixTime();

        multithreadMergeSort.sort();

        long elapsed_4 = Utils.getCurrentUnixTime() - start_4;

        // multithread quick sort

        MultithreadQuickSort multithreadQuickSort = new MultithreadQuickSort(array_5);

        long start_5 = Utils.getCurrentUnixTime();

        multithreadQuickSort.sort();

        long elapsed_5 = Utils.getCurrentUnixTime() - start_5;

        if (Constants.LOGS) {
            System.out.printf("Selection sort execution time: %d ms\n", elapsed_1);
            System.out.printf("Merge sort execution time: %d ms\n", elapsed_2);
            System.out.printf("Quick sort execution time: %d ms\n", elapsed_3);
            System.out.printf("Multithread merge sort execution time: %d ms\n", elapsed_4);
            System.out.printf("Multithread quick sort execution time: %d ms\n", elapsed_5);

            if (Constants.TOTAL_LOGS) {
                System.out.println(Arrays.toString(array_1));
                System.out.println(Arrays.toString(array_2));
                System.out.println(Arrays.toString(array_3));
                System.out.println(Arrays.toString(array_4));
                System.out.println(Arrays.toString(array_5));
            }
        }
    }

    private static void separator() {
        System.out.println();
    }
}
