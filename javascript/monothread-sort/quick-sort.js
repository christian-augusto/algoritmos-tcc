function sort(array) {
  quickSort(array, 0, array.length - 1);
}

function quickSort(array, low, high) {
  if (low < high) {
    const pi = partition(array, low, high);

    quickSort(array, low, pi - 1);
    quickSort(array, pi + 1, high);
  }
}

function partition(array, low, high) {
  const pivot = array[high];

  let i = low - 1;

  for (let j = low; j <= high - 1; j++) {
    if (array[j] < pivot) {
      i++;
      const aux = array[i];
      array[i] = array[j];
      array[j] = aux;
    }
  }

  const aux = array[i + 1];
  array[i + 1] = array[high];
  array[high] = aux;

  return i + 1;
}

module.exports = {
  sort,
};
