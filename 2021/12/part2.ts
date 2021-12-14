import { readInput } from "../utils.ts";
import { buildCaves, numberOfPaths } from "./shared.ts";

export async function AOC2021_12_2(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const caves = buildCaves(input);
  const sum = numberOfPaths(caves, "start", ["start"], true);
  return sum;
}
