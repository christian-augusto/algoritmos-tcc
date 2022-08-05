package algorithms.sort;

import algorithms.Constants;

/**
 * MultithreadMergeSort
 */
public class MultithreadMergeSort {
    private int[] array = null;

    public MultithreadMergeSort(int[] array) {
        this.array = array;
    }

    public void sort() throws Exception {
        int arrayLen = array.length;
        Thread[] threads = new Thread[Constants.THREADS_NUMBER];

        for (int i = 0; i < Constants.THREADS_NUMBER; i++) {
            int low = i * (arrayLen / 4);
            int high = (i + 1) * (arrayLen / 4) - 1;
            int mid = low + (high - low) / 2;

            Runnable runnable = new Runnable() {
                public void run() {
                    mergeSort(low, mid);
                    mergeSort(mid + 1, high);

                    merge(low, mid, high);
                }
            };

            Thread thread = new Thread(runnable);

            threads[i] = thread;

            thread.start();
        }

        for (Thread thread : threads) {
            thread.join();
        }

        merge(0, (arrayLen / 2 - 1) / 2, arrayLen / 2 - 1);
        merge(arrayLen / 2, arrayLen / 2 + (arrayLen - 1 - arrayLen / 2) / 2, arrayLen - 1);
        merge(0, (arrayLen - 1) / 2, arrayLen - 1);
    }

    private void mergeSort(int low, int high) {
        if (low < high) {
            int mid = low + (high - low) / 2;

            mergeSort(low, mid);
            mergeSort(mid + 1, high);

            merge(low, mid, high);
        }
    }

    private void merge(int low, int mid, int high) {
        int i, j, k;
        int n1 = mid - low + 1;
        int n2 = high - mid;

        int[] left = new int[n1];
        int[] right = new int[n2];

        for (i = 0; i < n1; i++) {
            left[i] = array[low + i];
        }

        for (i = 0; i < n2; i++) {
            right[i] = array[mid + 1 + i];
        }

        i = 0;
        j = 0;
        k = low;
        while (i < n1 && j < n2) {
            if (left[i] <= right[j]) {
                array[k] = left[i];
                i++;
            } else {
                array[k] = right[j];
                j++;
            }
            k++;
        }

        while (i < n1) {
            array[k] = left[i];
            k++;
            i++;
        }

        while (j < n2) {
            array[k] = right[j];
            k++;
            j++;
        }
    }
}
