const { THREADS_NUMBER } = require("../constants");

let threadsCount = 0;

async function sort(array) {
  await quickSort(array, 0, array.length - 1);
}

async function quickSort(array, low, high) {
  if (low < high) {
    const pi = partition(array, low, high);

    const promises = [];

    if (threadsCount < THREADS_NUMBER) {
      threadsCount++;

      promises.push(
        new Promise(async function (resolve, reject) {
          await quickSort(array, low, pi - 1);

          resolve(true);
        }),
      );
    } else {
      await quickSort(array, low, pi - 1);
    }

    if (threadsCount < THREADS_NUMBER) {
      threadsCount++;

      promises.push(
        new Promise(async function (resolve, reject) {
          await quickSort(array, pi + 1, high);

          resolve(true);
        }),
      );
    } else {
      await quickSort(array, pi + 1, high);
    }

    await Promise.all(promises);

    threadsCount -= promises.length;
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
