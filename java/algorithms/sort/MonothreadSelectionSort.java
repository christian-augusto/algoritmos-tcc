package algorithms.sort;

/**
 * MonothreadSelectionSort
 */
public class MonothreadSelectionSort {
    private int[] array = null;

    public MonothreadSelectionSort(int[] array) {
        this.array = array;
    }

    public void sort() {
        for (int i = 0; i < array.length; i++) {
            int index = i;

            for (int j = i + 1; j < array.length; j++) {
                if (array[j] < array[index]) {
                    index = j;
                }
            }

            if (index != i) {
                int aux = array[i];

                array[i] = array[index];
                array[index] = aux;
            }
        }
    }
}
