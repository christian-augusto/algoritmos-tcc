namespace utils
{
  using namespace std;
  using namespace std::chrono;

  void show_array(vector<int> *array)
  {
    for (int i = 0; i < array->size(); i++)
    {
      cout << array->operator[](i) << " ";
    }

    cout << "" << endl;
  }

  int random_int(int min, int max, unsigned someSeed)
  {
    return rand() % max + min;
  }

  vector<int> *generate_random_array(int size, int min, int max)
  {
    vector<int> *array = new vector<int>;

    for (int i = 0; i < size; i++)
    {
      array->push_back(random_int(min, max, i));
    }

    return array;
  }

  vector<int> *copy_array(vector<int> *array)
  {
    vector<int> *array_2 = new vector<int>;

    for (int i = 0; i < array->size(); i++)
    {
      array_2->push_back(array->operator[](i));
    }

    return array_2;
  }

  steady_clock::time_point get_current_unix_time_milliseconds()
  {
    return steady_clock::now();
  }
}
