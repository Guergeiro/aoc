import { readInput } from "../utils.ts";
import { buildMatrix, canGrow, isLowPoint } from "./shared.ts";

type Point = {
  value: number;
  row: number;
  col: number;
};

export async function AOC2021_09_2(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const matrix = buildMatrix(input);

  const lowestPoints: Point[] = [];
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (isLowPoint(matrix, i, j) === true) {
        lowestPoints.push({
          value: matrix[i][j],
          row: i,
          col: j,
        });
      }
    }
  }

  const basin: number[][] = [];
  for (const point of lowestPoints) {
    basin.push(findBasin(matrix, point.row, point.col));
  }

  const [a, b, c] = basin.sort(function (a, b) {
    return b.length - a.length;
  });

  return a.length * b.length * c.length;
}

function findBasin(
  matrix: number[][],
  row: number,
  col: number,
) {
  const basin: number[] = [];
  const currentNumber = matrix[row][col];
  if (currentNumber === -1) {
    // Already went through it
    return basin;
  }
  if (currentNumber === 9) {
    return basin;
  }
  basin.push(currentNumber);
  matrix[row][col] = -1;
  if (canGrow(matrix, row - 1, col) === true) {
    const upperBounds = findBasin(matrix, row - 1, col);
    basin.push(...upperBounds);
  }
  if (canGrow(matrix, row + 1, col) === true) {
    const lowerBounds = findBasin(matrix, row + 1, col);
    basin.push(...lowerBounds);
  }
  if (canGrow(matrix, row, col - 1) === true) {
    const leftBounds = findBasin(matrix, row, col - 1);
    basin.push(...leftBounds);
  }
  if (canGrow(matrix, row, col + 1) === true) {
    const rightBounds = findBasin(matrix, row, col + 1);
    basin.push(...rightBounds);
  }

  return basin;
}
