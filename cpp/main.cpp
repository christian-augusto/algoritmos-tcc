#include <iostream>
#include <ctime>
#include <vector>
#include "./constants/constants.cpp"
#include "./utils/utils.cpp"
#include "./monothread_sort/selection_sort.cpp"
#include "./monothread_sort/merge_sort.cpp"

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

  selection_sort::sort(array_1);

  // merge sort

  merge_sort::sort(array_2);

  if (LOGS)
  {
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
