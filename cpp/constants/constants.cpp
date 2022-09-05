namespace constants
{
  const bool DEBUG = true;
  const bool LOGS = true;
  const bool TOTAL_LOGS = DEBUG;
  const int ARRAY_SIZES_LENGTH = DEBUG ? 2 : 19;
  const int THREADS_NUMBER = 4;
  const int RANDOM_INT_MIN = 0;
  const int RANDOM_INT_MAX = DEBUG ? 100 : 2000000000;

  int *array_sizes()
  {
    int *array = new int[ARRAY_SIZES_LENGTH];

    if (DEBUG)
    {
      array[0] = 8;
      array[1] = 16;
    }

    array[0] = 8;
    array[1] = 16;
    array[2] = 100;
    array[3] = 1000;
    array[4] = 10000;
    array[5] = 25000;
    array[6] = 50000;
    array[7] = 75000;
    array[8] = 100000;
    array[9] = 250000;
    array[10] = 500000;
    array[11] = 750000;
    array[12] = 1000000;
    array[13] = 5000000;
    array[14] = 10000000;
    array[15] = 25000000;
    array[16] = 50000000;
    array[17] = 100000000;
    array[18] = 250000000;

    return array;
  }

  const int *ARRAY_SIZES = array_sizes();
}
