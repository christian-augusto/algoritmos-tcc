const { LOGS, TOTAL_LOGS, ARRAY_SIZES } = require("./constants");
const monothreadSort = require("./monothread-sort");
const multithreadSort = require("./multithread-sort");
const utils = require("./utils");

async function main() {
  for (let i = 0; i < ARRAY_SIZES.length; i++) {
    const size = ARRAY_SIZES[i];

    await sort(size);
  }
}

async function sort(size) {
  console.log(`size: ${size}`);

  const array1 = utils.genRandomIntegersArray(size);
  const array2 = utils.copyArray(array1);
  const array3 = utils.copyArray(array1);
  const array4 = utils.copyArray(array1);

  if (LOGS) {
    if (TOTAL_LOGS) {
      console.log(array1);
    }
  }

  // selection sort

  const start1 = new Date();

  monothreadSort.selectionSort.sort(array1);

  const elapsed1 = utils.timeElapsedSinceMiliSeconds(start1);

  // merge sort

  const start2 = new Date();

  monothreadSort.mergeSort.sort(array2);

  const elapsed2 = utils.timeElapsedSinceMiliSeconds(start2);

  // quick sort

  const start3 = new Date();

  monothreadSort.quickSort.sort(array3);

  const elapsed3 = utils.timeElapsedSinceMiliSeconds(start3);

  // multithread merge sort

  const start4 = new Date();

  await multithreadSort.mergeSort.sort(array4);

  const elapsed4 = utils.timeElapsedSinceMiliSeconds(start4);

  if (LOGS) {
    console.log(`Selection sort execution time: ${elapsed1} ms`);
    console.log(`Merge sort execution time: ${elapsed2} ms`);
    console.log(`Quick sort execution time: ${elapsed3} ms`);
    console.log(`Multithread merge sort execution time: ${elapsed4} ms`);
    console.log("\n");

    if (TOTAL_LOGS) {
      console.log(array1);
      console.log(array2);
      console.log(array3);
      console.log(array4);
    }
  }
}

main();
