#include <iostream>
#include <vector>
#include <thread>
#include <chrono>
#include "./constants/constants.cpp"
#include "./utils/utils.cpp"
#include "./monothread_sort/selection_sort.cpp"
#include "./monothread_sort/merge_sort.cpp"
#include "./monothread_sort/quick_sort.cpp"
#include "./multithread_sort/merge_sort.cpp"
#include "./multithread_sort/quick_sort.cpp"

using namespace std;
using namespace constants;

void sort(int size);
void separator();

int main()
{
  for (int i = 0; i < ARRAY_SIZES_LENGTH; i++)
  {
    int size = ARRAY_SIZES[i];

    cout << "size: " << size << endl;

    sort(size);

    separator();
  }

  return 0;
}

void sort(int size)
{
  vector<int> *array_1 = utils::generate_random_array(size, RANDOM_INT_MIN, RANDOM_INT_MAX);
  vector<int> *array_2 = utils::copy_array(array_1);
  vector<int> *array_3 = utils::copy_array(array_1);
  vector<int> *array_4 = utils::copy_array(array_1);
  vector<int> *array_5 = utils::copy_array(array_1);

  if (LOGS && TOTAL_LOGS)
  {
    utils::show_array(array_1);
    utils::show_array(array_2);
    utils::show_array(array_3);
    utils::show_array(array_4);
    utils::show_array(array_5);
  }

  // selection sort

  const long start_1 = utils::get_current_unix_time_milliseconds();

  selection_sort::sort(array_1);

  const long elapsed_1 = utils::get_current_unix_time_milliseconds() - start_1;

  // merge sort

  const long start_2 = utils::get_current_unix_time_milliseconds();

  merge_sort::sort(array_2);

  const long elapsed_2 = utils::get_current_unix_time_milliseconds() - start_2;

  // quick sort

  const long start_3 = utils::get_current_unix_time_milliseconds();

  quick_sort::sort(array_3);

  const long elapsed_3 = utils::get_current_unix_time_milliseconds() - start_3;

  // multithread merge sort

  const long start_4 = utils::get_current_unix_time_milliseconds();

  multithread_merge_sort::sort(array_4);

  const long elapsed_4 = utils::get_current_unix_time_milliseconds() - start_4;

  // multithread merge sort

  const long start_5 = utils::get_current_unix_time_milliseconds();

  multithread_quick_sort::sort(array_5);

  const long elapsed_5 = utils::get_current_unix_time_milliseconds() - start_5;

  if (LOGS)
  {
    cout << "Selection sort execution time: "
         << elapsed_1 << " ms" << endl;
    cout << "Merge sort execution time: "
         << elapsed_2 << " ms" << endl;
    cout << "Quick sort execution time: "
         << elapsed_3 << " ms" << endl;
    cout << "Multithread merge sort execution time: "
         << elapsed_4 << " ms" << endl;
    cout << "Multithread quick sort execution time: "
         << elapsed_5 << " ms" << endl;

    if (TOTAL_LOGS)
    {
      utils::show_array(array_1);
      utils::show_array(array_2);
      utils::show_array(array_3);
      utils::show_array(array_4);
      utils::show_array(array_5);
    }
  }
}

void separator()
{
  cout << endl;
}
