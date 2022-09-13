const { LOGS, TOTAL_LOGS, ARRAY_SIZES } = require("./constants");
const monothreadSort = require("./monothread-sort");
const multithreadSort = require("./multithread-sort");
const utils = require("./utils");

async function main() {
  for (let i = 0; i < ARRAY_SIZES.length; i++) {
    const size = ARRAY_SIZES[i];

    console.log(`size: ${size}`);

    await sort(size);

    separator();
  }
}

async function sort(size) {
  const array_1 = utils.genRandomIntegersArray(size);
  const array_2 = utils.copyArray(array_1);
  const array_3 = utils.copyArray(array_1);
  const array_4 = utils.copyArray(array_1);
  const array_5 = utils.copyArray(array_1);

  if (LOGS && TOTAL_LOGS) {
    console.log(array_1);
    console.log(array_2);
    console.log(array_3);
    console.log(array_4);
    console.log(array_5);
  }

  // selection sort

  const start_1 = new Date();

  monothreadSort.selectionSort.sort(array_1);

  const elapsed_1 = utils.timeElapsedSinceMiliSeconds(start_1);

  // merge sort

  const start_2 = new Date();

  monothreadSort.mergeSort.sort(array_2);

  const elapsed_2 = utils.timeElapsedSinceMiliSeconds(start_2);

  // quick sort

  const start_3 = new Date();

  monothreadSort.quickSort.sort(array_3);

  const elapsed_3 = utils.timeElapsedSinceMiliSeconds(start_3);

  // multithread merge sort

  const start_4 = new Date();

  await multithreadSort.mergeSort.sort(array_4);

  const elapsed_4 = utils.timeElapsedSinceMiliSeconds(start_4);

  // multithread quick sort

  const start_5 = new Date();

  await multithreadSort.quickSort.sort(array_5);

  const elapsed_5 = utils.timeElapsedSinceMiliSeconds(start_5);

  if (LOGS) {
    console.log(`Selection sort execution time: ${elapsed_1} ms`);
    console.log(`Merge sort execution time: ${elapsed_2} ms`);
    console.log(`Quick sort execution time: ${elapsed_3} ms`);
    console.log(`Multithread merge sort execution time: ${elapsed_4} ms`);
    console.log(`Multithread quick sort execution time: ${elapsed_5} ms`);

    if (TOTAL_LOGS) {
      console.log(array_1);
      console.log(array_2);
      console.log(array_3);
      console.log(array_4);
      console.log(array_5);
    }
  }
}

function separator() {
  console.log();
}

main();
