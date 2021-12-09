import { readInput } from "../utils.ts";

export async function AOC2021_08_1(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const sum = input.reduce(function (sum, light) {
    const [_firstPart, secondPart] = light.split(" | ");
    const output = secondPart.split(" ");
    return sum += output.reduce(function (s, part) {
      switch (part.length) {
        case 2:
        case 3:
        case 4:
        case 7:
          return s += 1;
        default:
          return s;
      }
    }, 0);
  }, 0);

  return sum;
}
