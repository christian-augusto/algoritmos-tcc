package algorithms.sort;

/**
 * MonothreadMergeSort
 */
public class MonothreadMergeSort {
    private int[] array = null;

    public MonothreadMergeSort(int[] array) {
        this.array = array;
    }

    public void sort() {
        mergeSort(0, array.length - 1);
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
