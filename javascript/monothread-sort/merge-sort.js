function sort(array) {
  mergeSort(array, 0, array.length - 1);
}

function mergeSort(array, low, high) {
  if (low < high) {
    const mid = parseInt(low + (high - low) / 2);

    mergeSort(array, low, mid);
    mergeSort(array, mid + 1, high);

    merge(array, low, mid, high);
  }
}

function merge(array, low, mid, high) {
  let i, j, k;
  const n1 = mid - low + 1;
  const n2 = high - mid;

  const left = new Array(n1);
  const right = new Array(n2);

  for (i = 0; i < n1; i++) {
    left[i] = array[low + i];
  }

  for (j = 0; j < n2; j++) {
    right[j] = array[mid + 1 + j];
  }

  i = 0;
  j = 0;
  k = low;
  for (; i < n1 && j < n2; ) {
    if (left[i] <= right[j]) {
      array[k] = left[i];
      i++;
    } else {
      array[k] = right[j];
      j++;
    }
    k++;
  }

  for (; i < n1; ) {
    array[k] = left[i];
    k++;
    i++;
  }

  for (; j < n2; ) {
    array[k] = right[j];
    k++;
    j++;
  }
}

module.exports = {
  sort,
};
