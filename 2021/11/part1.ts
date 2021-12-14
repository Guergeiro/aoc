import { readInput } from "../utils.ts";
import { getOctupus, simulate } from "./shared.ts";

export async function AOC2021_11_1(file: string) {
  const input = (await readInput(file)).split("\n");

  return calculate(input);
}

export function calculate(input: string[]) {
  const octupus = getOctupus(input);

  simulate(octupus, 100, function (step) {
    let stillHasFlash = false;
    do {
      stillHasFlash = octupus.some(function (oct) {
        return oct.needsToFlash();
      });

      for (const oct of octupus) {
        oct.flash(step);
      }
    } while (stillHasFlash);
    return false;
  });

  return octupus.reduce(function (sum, cur) {
    return sum += cur.getFlashes();
  }, 0);
}
