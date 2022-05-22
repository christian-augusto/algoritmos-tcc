const { LOGS, TOTAL_LOGS, MIN_ARRAY_LENGTH, MAX_ARRAY_LENGTH } = require("./constants");
const monothreadSort = require("./monothread-sort");
const utils = require("./utils");

function main() {
  for (let size = MIN_ARRAY_LENGTH; size <= MAX_ARRAY_LENGTH; size++) {
    sort(size);
  }
}

function sort(size) {
  console.log(`size: ${size}`);

  const array1 = utils.genRandomIntegersArray(size);
  const array2 = utils.copyArray(array1);

  if (LOGS) {
    if (TOTAL_LOGS) {
      console.log(array1);
    }
  }

  const start1 = new Date();

  monothreadSort.selectionSort.sort(array1);

  const elapsed1 = utils.timeElapsedSinceSeconds(start1);

  const start2 = new Date();

  monothreadSort.mergeSort.sort(array2);

  const elapsed2 = utils.timeElapsedSinceSeconds(start2);

  if (LOGS) {
    console.log(`Selection sort execution time: ${elapsed1}s`);
    console.log(`Merge sort execution time: ${elapsed2}s`);
    console.log("\n");

    if (TOTAL_LOGS) {
      console.log(array1);
      console.log(array2);
    }
  }
}

main();
