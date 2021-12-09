import { readInput } from "../utils.ts";
import { buildMatrix, isLowPoint } from "./shared.ts";

export async function AOC2021_09_1(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const matrix = buildMatrix(input);

  const lowestPoints: number[] = [];
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (isLowPoint(matrix, i, j) === true) {
        lowestPoints.push(matrix[i][j]);
      }
    }
  }
  return lowestPoints.reduce(function (sum, cur) {
    return sum += 1 + cur;
  }, 0);
}
