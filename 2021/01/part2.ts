import { readInput } from "../utils.ts";

export async function AOC2021_01_2(file: string) {
  const splittedString = (await readInput(file)).split("\n");

  const measurements = splittedString.map(function (v) {
    return parseInt(v);
  });

  return calculate(measurements);
}

export function calculate(measurements: number[]) {
  let increased = 0;

  let prev = measurements[0];
  for (let i = 3; i < measurements.length; i++) {
    if (prev < measurements[i]) {
      increased += 1;
    }
    prev = measurements[i - 2];
  }

  return increased;
}
