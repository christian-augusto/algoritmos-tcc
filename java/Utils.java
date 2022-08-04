import java.util.Arrays;
import java.util.concurrent.ThreadLocalRandom;

/**
 * Utils
 */
public class Utils {
    public static int[] generateRandomIntArray(int size, int min, int max) {
        int[] array = new int[size];

        for (int i = 0; i < size; i++) {
            array[i] = ThreadLocalRandom.current().nextInt(min, max + 1);
        }

        return array;
    }

    public static int[] copyArray(int[] array) {
        return Arrays.copyOf(array, array.length);
    }

    public static long getCurrentUnixTime() {
        return System.currentTimeMillis() / 1000L;
    }
}
