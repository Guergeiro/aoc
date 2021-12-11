import { readInput } from "../utils.ts";
import { getCorruptedChar, getCorruptedValue } from "./shared.ts";

export async function AOC2021_10_1(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const incomplete: string[] = [];

  for (const line of input) {
    const corruptedChar = getCorruptedChar(line);
    if (corruptedChar != null) {
      incomplete.push(corruptedChar);
    }
  }

  return incomplete.reduce(function (sum, cur) {
    return sum += getCorruptedValue(cur);
  }, 0);
}
