import { readInput } from "../utils.ts";
import { calculate } from "./shared.ts";
export async function AOC2021_07_1(file: string) {
  const input = await readInput(file);
  const parsedInput = input.split(",").map(function (v) {
    return parseInt(v);
  });
  return execute(parsedInput);
}

export function execute(input: number[]) {
  return calculate(input, function (acc, curr, depth) {
    return acc += Math.abs(curr - depth);
  });
}
