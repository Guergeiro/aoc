import { readInput } from "../utils.ts";
import {
  getCorruptedChar,
  getMissingClosingChars,
  getMissingValue,
} from "./shared.ts";

export async function AOC2021_10_2(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const incomplete = input.filter(function (line) {
    return getCorruptedChar(line) == null;
  });

  const allScores: number[] = [];

  for (const line of incomplete) {
    const chars = getMissingClosingChars(line);
    const value = chars.reduce(function (sum, cur) {
      return (sum * 5) + getMissingValue(cur);
    }, 0);
    allScores.push(value);
  }

  const sortedScores = allScores.sort(function (a, b) {
    return a - b;
  });

  return sortedScores[Math.floor(sortedScores.length / 2)];
}
