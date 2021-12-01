import { readInput } from "../utils.ts";

export async function AOC2021_01_1(file: string) {
  const splittedString = (await readInput(file)).split("\n");
  const measurements = splittedString.map(function (v) {
    return parseInt(v);
  });

  return calculate(measurements);
}

export function calculate(measurements: number[]) {
  let increased = 0;
  measurements.reduce(function (prev, cur) {
    if (prev < cur) {
      increased += 1;
    }
    return cur;
  });

  return increased;
}
