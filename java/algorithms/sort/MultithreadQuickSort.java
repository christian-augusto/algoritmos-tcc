package algorithms.sort;

import algorithms.Constants;

/**
 * MultithreadQuickSort
 */
public class MultithreadQuickSort {
    private int[] array = null;
    private int threadsCount = 0;

    public MultithreadQuickSort(int[] array) {
        this.array = array;
    }

    public void sort() throws Exception {
        quickSort(0, array.length - 1);
    }

    private void quickSort(int low, int high) throws Exception {
        if (low < high) {
            int pi = partition(low, high);
            Thread thread_1 = null;
            Thread thread_2 = null;
            int localThreadsCount = 0;

            if (threadsCount < Constants.THREADS_NUMBER) {
                threadsCount++;
                localThreadsCount++;

                Runnable runnable = new Runnable() {
                    public void run() {
                        try {
                            quickSort(low, pi - 1);
                        } catch (Exception e) {
                            e.printStackTrace();
                        }
                    }
                };

                thread_1 = new Thread(runnable);

                thread_1.start();
            } else {
                quickSort(low, pi - 1);
            }

            if (threadsCount < Constants.THREADS_NUMBER) {
                threadsCount++;
                localThreadsCount++;

                Runnable runnable = new Runnable() {
                    public void run() {
                        try {
                            quickSort(pi + 1, high);
                        } catch (Exception e) {
                            e.printStackTrace();
                        }
                    }
                };

                thread_2 = new Thread(runnable);

                thread_2.start();
            } else {
                quickSort(pi + 1, high);
            }

            if (thread_1 != null) {
                thread_1.join();
            }

            if (thread_2 != null) {
                thread_2.join();
            }

            threadsCount -= localThreadsCount;
        }
    }

    private int partition(int low, int high) {
        int pivot = array[high];

        int i = low - 1;

        for (int j = low; j <= high - 1; j++) {
            if (array[j] < pivot) {
                i++;
                int aux = array[i];
                array[i] = array[j];
                array[j] = aux;
            }
        }

        int aux = array[i + 1];
        array[i + 1] = array[high];
        array[high] = aux;

        return i + 1;
    }
}
