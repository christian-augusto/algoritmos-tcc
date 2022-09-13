namespace multithread_quick_sort
{
  using namespace std;
  using namespace constants;

  int threads_count = 0;

  int partition(vector<int> *array, int low, int high)
  {
    int pivot = array->operator[](high);

    int i = low - 1;

    for (int j = low; j <= high - 1; j++)
    {
      if (array->operator[](j) < pivot)
      {
        i++;
        int aux = array->operator[](i);
        array->operator[](i) = array->operator[](j);
        array->operator[](j) = aux;
      }
    }

    int aux = array->operator[](i + 1);
    array->operator[](i + 1) = array->operator[](high);
    array->operator[](high) = aux;

    return i + 1;
  }

  void quick_sort(vector<int> *array, int low, int high)
  {
    if (low < high)
    {
      int pi = partition(array, low, high);

      vector<thread> threads;

      if (threads_count < THREADS_NUMBER)
      {
        threads_count++;

        thread t(quick_sort, array, low, pi - 1);

        threads.push_back(move(t));
      }
      else
      {
        quick_sort(array, low, pi - 1);
      }

      if (threads_count < THREADS_NUMBER)
      {
        threads_count++;

        thread t(quick_sort, array, pi + 1, high);

        threads.push_back(move(t));
      }
      else
      {
        quick_sort(array, pi + 1, high);
      }

      for (thread &t : threads)
      {
        t.join();
      }

      threads_count -= threads.size();
    }
  }

  void sort(vector<int> *array)
  {
    quick_sort(array, 0, array->size() - 1);
  }
}
