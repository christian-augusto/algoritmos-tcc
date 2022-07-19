const { THREADS_NUMBER } = require("../constants");
const { intDivision } = require("../utils");

async function sort(array) {
  const arrayLen = array.length;

  const promises = [];

  for (let i = 0; i < THREADS_NUMBER; i++) {
    const low = i * (arrayLen / 4);
    const high = (i + 1) * (arrayLen / 4) - 1;
    const mid = low + intDivision(high - low, 2);

    promises.push(
      new Promise(function (resolve) {
        mergeSort(array, low, mid);
        mergeSort(array, mid + 1, high);

        merge(array, low, mid, high);

        resolve(true);
      }),
    );
  }

  await Promise.all(promises);

  const low_1 = 0;
  const mid_1 = intDivision(arrayLen / 2 - 1, 2);
  const high_1 = arrayLen / 2 - 1;

  merge(array, low_1, mid_1, high_1);

  const low_2 = arrayLen / 2;
  const mid_2 = arrayLen / 2 + intDivision(arrayLen - 1 - arrayLen / 2, 2);
  const high_2 = arrayLen - 1;

  merge(array, low_2, mid_2, high_2);

  const low_3 = 0;
  const mid_3 = intDivision(arrayLen - 1, 2);
  const high_3 = arrayLen - 1;

  merge(array, low_3, mid_3, high_3);
}

function mergeSort(array, low, high) {
  if (low < high) {
    const mid = low + intDivision(high - low, 2);

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
