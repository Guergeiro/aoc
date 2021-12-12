import { readInput } from "../utils.ts";
import { getOctupus,  simulate } from "./shared.ts";

export async function AOC2021_11_2(file: string) {
  const input = (await readInput(file)).split("\n");

  return calculate(input);
}

export function calculate(input: string[]) {
  const octupus = getOctupus(input);

  const lastStep = simulate(octupus, Infinity, function () {
    const allFlash = octupus.every(function (oct) {
      return oct.needsToFlash();
    });

    return allFlash;
  });

  // Normalize for extra step
  return lastStep - 10;
}
