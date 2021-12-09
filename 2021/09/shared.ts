export function isLowPoint(matrix: number[][], row: number, col: number) {
  const currentNumber = matrix[row][col];
  if (canGrow(matrix, row - 1, col) === true) {
    // check previous row
    if (matrix[row - 1][col] <= currentNumber) {
      return false;
    }
  }
  if (canGrow(matrix, row + 1, col) === true) {
    // check next row
    if (matrix[row + 1][col] <= currentNumber) {
      return false;
    }
  }
  if (canGrow(matrix, row, col - 1) === true) {
    // check previous col
    if (matrix[row][col - 1] <= currentNumber) {
      return false;
    }
  }
  if (canGrow(matrix, row, col + 1) === true) {
    // check next col
    if (matrix[row][col + 1] <= currentNumber) {
      return false;
    }
  }
  return true;
}

export function canGrow(matrix: number[][], row: number, col: number) {
  if (row < 0) {
    return false;
  }
  if (row > matrix.length - 1) {
    return false;
  }
  if (col < 0) {
    return false;
  }
  if (col > matrix[row].length - 1) {
    return false;
  }
  return true;
}

export function buildMatrix(lines: string[]) {
  return lines.map(function (line) {
    return line.split("").map(function (value) {
      return parseInt(value);
    });
  });
}
