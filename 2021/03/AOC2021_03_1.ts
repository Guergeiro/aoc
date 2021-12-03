import { readInput } from "../utils.ts";
import {
  convertToDecimal,
  getBinary,
  makeMatrix,
  populateMap,
} from "./shared.ts";

export async function AOC2021_03_1(file: string) {
  const splittedString = (await readInput(file)).split("\n");

  return calculatePowerConsumption(splittedString);
}

export function calculatePowerConsumption(splittedString: string[]) {
  const matrix = makeMatrix(splittedString);

  const mapOfSums = populateMap(matrix);

  const gamma = getBinary(mapOfSums, function (sum) {
    if (sum > Math.round(matrix.length / 2)) {
      return "1";
    }
    return "0";
  });

  const epsilion = getBinary(mapOfSums, function (sum) {
    if (sum > Math.round(matrix.length / 2)) {
      return "0";
    }
    return "1";
  });

  return convertToDecimal(gamma) * convertToDecimal(epsilion);
}
