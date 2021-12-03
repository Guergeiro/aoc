import { readInput } from "../utils.ts";
import {
  convertToDecimal,
  getBinary,
  makeMatrix,
  populateMap,
} from "./shared.ts";

export async function AOC2021_03_2(file: string) {
  const splittedString = (await readInput(file)).split("\n");

  return calculatePowerConsumption(splittedString);
}

export function calculatePowerConsumption(splittedString: string[]) {
  const maxNumCols = splittedString.reduce(function (prev, curr) {
    if (curr.length > prev) {
      return curr.length;
    }
    return prev;
  }, 0);

  const [gamma] = calculate(
    [...splittedString],
    maxNumCols,
    function (char1, char2) {
      return char1 === char2;
    },
  );
  const [epsilion] = calculate(
    [...splittedString],
    maxNumCols,
    function (char1, char2) {
      return char1 !== char2;
    },
  );

  return convertToDecimal(gamma) * convertToDecimal(epsilion);
}

function calculate(
  arr: string[],
  maxNumCols: number,
  fn: (char1: string, char2: string) => boolean,
) {
  for (let i = 0; i < maxNumCols; i++) {
    const matrix = makeMatrix(arr);
    const m = populateMap(matrix);
    const binaryChar = getBinary(m, function (sum) {
      if (sum >= Math.round(matrix.length / 2)) {
        return "1";
      }
      return "0";
    })[i];
    const validArrays = arr.filter(function (row) {
      return fn(row[i], binaryChar);
    });
    arr.length = 0;
    arr.push(...validArrays);
    if (arr.length === 1) {
      break;
    }
  }
  return arr;
}
