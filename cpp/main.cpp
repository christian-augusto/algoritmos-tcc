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
using namespace std::chrono;

void sort(int arraySize);
void separator();

int main()
{
  for (int i = 0; i < ARRAY_SIZES_LENGTH; i++)
  {
    int arraySize = ARRAY_SIZES[i];

    cout << "arraySize: " << arraySize << endl;

    sort(arraySize);

    separator();
  }

  return 0;
}

void sort(int arraySize)
{
  vector<int> *array_1 = utils::generate_random_array(arraySize, RANDOM_INT_MIN, RANDOM_INT_MAX);
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

  const steady_clock::time_point start_1 = utils::get_current_unix_time_milliseconds();

  selection_sort::sort(array_1);

  const steady_clock::time_point end_1 = utils::get_current_unix_time_milliseconds();

  const long elapsed_1 = duration_cast<milliseconds>(end_1 - start_1).count();

  // merge sort

  const steady_clock::time_point start_2 = utils::get_current_unix_time_milliseconds();

  merge_sort::sort(array_2);

  const steady_clock::time_point end_2 = utils::get_current_unix_time_milliseconds();

  const long elapsed_2 = duration_cast<milliseconds>(end_2 - start_2).count();

  // quick sort

  const steady_clock::time_point start_3 = utils::get_current_unix_time_milliseconds();

  quick_sort::sort(array_3);

  const steady_clock::time_point end_3 = utils::get_current_unix_time_milliseconds();

  const long elapsed_3 = duration_cast<milliseconds>(end_3 - start_3).count();

  // multithread merge sort

  const steady_clock::time_point start_4 = utils::get_current_unix_time_milliseconds();

  multithread_merge_sort::sort(array_4);

  const steady_clock::time_point end_4 = utils::get_current_unix_time_milliseconds();

  const long elapsed_4 = duration_cast<milliseconds>(end_4 - start_4).count();

  // multithread merge sort

  const steady_clock::time_point start_5 = utils::get_current_unix_time_milliseconds();

  multithread_quick_sort::sort(array_5);

  const steady_clock::time_point end_5 = utils::get_current_unix_time_milliseconds();

  const long elapsed_5 = duration_cast<milliseconds>(end_5 - start_5).count();

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
