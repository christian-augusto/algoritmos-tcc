const { MIN_NUMBER, MAX_NUMBER } = require("../constants");

function genRandomNumber(min, max) {
  return Math.random() * (max - min) + min;
}

function genRandomIntegersArray(size) {
  const array = [];

  for (let i = 0; i < size; i++) {
    array.push(parseInt(genRandomNumber(MIN_NUMBER, MAX_NUMBER)));
  }

  return array;
}

function copyArray(array) {
  const newArray = [];

  for (let i = 0; i < array.length; i++) {
    newArray.push(array[i]);
  }

  return newArray;
}

function timeElapsedSinceMiliSeconds(start) {
  return new Date() - start;
}

function intDivision(a, b) {
  return parseInt(a / b);
}

module.exports = {
  genRandomIntegersArray,
  copyArray,
  timeElapsedSinceMiliSeconds,
  intDivision,
};
