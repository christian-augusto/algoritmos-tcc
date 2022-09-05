namespace selection_sort
{
  using namespace std;

  void sort(vector<int> *array)
  {
    for (int i = 0; i < array->size(); i++)
    {
      int index = i;

      for (int j = i + 1; j < array->size(); j++)
      {
        if (array->operator[](j) < array->operator[](index))
        {
          index = j;
        }
      }

      if (index != i)
      {
        int aux = array->operator[](i);

        array->operator[](i) = array->operator[](index);
        array->operator[](index) = aux;
      }
    }
  }
}
