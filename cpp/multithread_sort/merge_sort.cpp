namespace multithread_merge_sort
{
  using namespace std;
  using namespace constants;

  void merge(vector<int> *array, int low, int mid, int high)
  {
    int i, j, k;
    int n1 = mid - low + 1;
    int n2 = high - mid;

    vector<int> *left = new vector<int>;
    vector<int> *right = new vector<int>;

    for (i = 0; i < n1; i++)
    {
      left->push_back(array->operator[](low + i));
    }

    for (i = 0; i < n2; i++)
    {
      right->push_back(array->operator[](mid + 1 + i));
    }

    i = 0;
    j = 0;
    k = low;
    while (i < n1 && j < n2)
    {
      if (left->operator[](i) <= right->operator[](j))
      {
        array->operator[](k) = left->operator[](i);
        i++;
      }
      else
      {
        array->operator[](k) = right->operator[](j);
        j++;
      }
      k++;
    }

    while (i < n1)
    {
      array->operator[](k) = left->operator[](i);
      k++;
      i++;
    }

    while (j < n2)
    {
      array->operator[](k) = right->operator[](j);
      k++;
      j++;
    }
  }

  void mergeSort(vector<int> *array, int low, int high)
  {
    if (low < high)
    {
      int mid = low + (high - low) / 2;

      mergeSort(array, low, mid);
      mergeSort(array, mid + 1, high);

      merge(array, low, mid, high);
    }
  }

  void thread_sort(vector<int> *array, int low, int mid, int high)
  {
    mergeSort(array, low, mid);
    mergeSort(array, mid + 1, high);

    merge(array, low, mid, high);
  }

  void sort(vector<int> *array)
  {
    int arrayLen = array->size();
    vector<thread> threads;

    for (int i = 0; i < THREADS_NUMBER; i++)
    {
      int low = i * (arrayLen / 4);
      int high = (i + 1) * (arrayLen / 4) - 1;
      int mid = low + (high - low) / 2;

      thread t(thread_sort, array, low, mid, high);

      threads.push_back(move(t));
    }

    for (thread &t : threads)
    {
      t.join();
    }

    merge(array, 0, (arrayLen / 2 - 1) / 2, arrayLen / 2 - 1);
    merge(array, arrayLen / 2, arrayLen / 2 + (arrayLen - 1 - arrayLen / 2) / 2, arrayLen - 1);
    merge(array, 0, (arrayLen - 1) / 2, arrayLen - 1);
  }
}
