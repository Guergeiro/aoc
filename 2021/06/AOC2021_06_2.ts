import { readInput } from "../utils.ts";
import { simulateLifecycle } from "./shared.ts";

export async function AOC2021_06_2(file: string) {
  const input = (await readInput(file));
  const fish = input.split(",").map(function (i) {
    return parseInt(i);
  });
  return simulateLifecycle(fish, 256);
}
