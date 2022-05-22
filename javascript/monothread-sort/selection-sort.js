function sort(array) {
  for (let i = 0; i < array.length; i++) {
    let index = i;

    for (let j = i + 1; j < array.length; j++) {
      if (array[j] < array[index]) {
        index = j;
      }
    }

    if (index != i) {
      let aux = array[i];

      array[i] = array[index];
      array[index] = aux;
    }
  }
}

module.exports = {
  sort,
};
