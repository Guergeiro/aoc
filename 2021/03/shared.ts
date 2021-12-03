export function makeMatrix(splittedString: string[]) {
  const matrix = splittedString.map(function (row) {
    const cols = row.split("");

    return cols;
  });

  return matrix;
}

export function populateMap(matrix: string[][]) {
  const m = new Map<number, number>();

  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (m.has(j) === false) {
        m.set(j, 0);
      }
      m.set(j, m.get(j)! + parseInt(matrix[i][j]));
    }
  }

  return m;
}

export function getBinary(m: Map<number, number>, fn: (sum: number) => string) {
  const binary: string[] = [];
  for (const sum of m.values()) {
    binary.push(fn(sum));
  }
  return binary.join("");
}

export function convertToDecimal(binary: string) {
  return parseInt(binary, 2);
}
