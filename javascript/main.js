const { LOGS, TOTAL_LOGS, ARRAY_SIZES } = require("./constants");
const monothreadSort = require("./monothread-sort");
const utils = require("./utils");

function main() {
  ARRAY_SIZES.forEach(function (size) {
    sort(size);
  });
}

function sort(size) {
  console.log(`size: ${size}`);

  const array1 = utils.genRandomIntegersArray(size);
  const array2 = utils.copyArray(array1);
  const array3 = utils.copyArray(array1);

  if (LOGS) {
    if (TOTAL_LOGS) {
      console.log(array1);
    }
  }

  // selection sort

  const start1 = new Date();

  monothreadSort.selectionSort.sort(array1);

  const elapsed1 = utils.timeElapsedSinceSeconds(start1);

  // merge sort

  const start2 = new Date();

  monothreadSort.mergeSort.sort(array2);

  const elapsed2 = utils.timeElapsedSinceSeconds(start2);

  // quick sort

  const start3 = new Date();

  monothreadSort.mergeSort.sort(array3);

  const elapsed3 = utils.timeElapsedSinceSeconds(start3);

  if (LOGS) {
    console.log(`Selection sort execution time: ${elapsed1}s`);
    console.log(`Merge sort execution time: ${elapsed2}s`);
    console.log(`Quick sort execution time: ${elapsed3}s`);
    console.log("\n");

    if (TOTAL_LOGS) {
      console.log(array1);
      console.log(array2);
      console.log(array3);
    }
  }
}

main();
