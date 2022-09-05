namespace quick_sort
{
  using namespace std;

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

  void quickSort(vector<int> *array, int low, int high)
  {
    if (low < high)
    {
      int pi = partition(array, low, high);

      quickSort(array, low, pi - 1);
      quickSort(array, pi + 1, high);
    }
  }

  void sort(vector<int> *array)
  {
    quickSort(array, 0, array->size() - 1);
  }
}
