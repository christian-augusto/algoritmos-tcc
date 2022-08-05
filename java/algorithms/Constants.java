package algorithms;

/**
 * Constants
 */
public class Constants {
    public static boolean DEBUG = true;
    public static boolean LOGS = true;
    public static boolean TOTAL_LOGS = DEBUG;
    public static int[] ARRAY_SIZES = arraySizes();
    public static int THREADS_NUMBER = 4;
    public static int RANDOM_INT_MIN = 0;
    public static int RANDOM_INT_MAX = DEBUG ? 100 : 2_000_000_000;

    private static int[] arraySizes() {
        int[] array = null;

        if (DEBUG) {
            array = new int[2];
            array[0] = 8;
            array[1] = 16;
        } else {
            array = new int[14];
            array[0] = 8;
            array[1] = 16;
            array[2] = 100;
            array[3] = 1_000;
            array[4] = 10_000;
            array[5] = 25_000;
            array[6] = 50_000;
            array[7] = 75_000;
            array[8] = 100_000;
            array[9] = 250_000;
            array[10] = 500_000;
            array[11] = 750_000;
            array[12] = 1_000_000;
            array[13] = 5_000_000;
        }

        return array;
    }
}
