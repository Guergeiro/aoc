import { readInput } from "../utils.ts";
import { permutations } from "../deps.ts";

export async function AOC2021_08_2(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

const CHAR_CODES = "abcdefg".split("").map(function (char) {
  return char.charCodeAt(0);
});

const keyFromCharCodes = (...chars: number[]): string => {
  return chars.sort(function (a, b) {
    return a - b;
  }).join();
};

const mapKeysToDigits = (
  [a, b, c, d, e, f, g]: number[],
): Record<string, number> => {
  return {
    [keyFromCharCodes(a, b, c, e, f, g)]: 0,
    [keyFromCharCodes(c, f)]: 1,
    [keyFromCharCodes(a, c, d, e, g)]: 2,
    [keyFromCharCodes(a, c, d, f, g)]: 3,
    [keyFromCharCodes(b, c, d, f)]: 4,
    [keyFromCharCodes(a, b, d, f, g)]: 5,
    [keyFromCharCodes(a, b, d, e, f, g)]: 6,
    [keyFromCharCodes(a, c, f)]: 7,
    [keyFromCharCodes(a, b, c, d, e, f, g)]: 8,
    [keyFromCharCodes(a, b, c, d, f, g)]: 9,
  };
};

const keyFromChars = (chars: string): string => {
  return chars
    .split("")
    .map(function (char) {
      return char.charCodeAt(0);
    })
    .sort(function (a, b) {
      return a - b;
    })
    .join();
};

export function calculate(input: string[]) {
  let outputValuesTotal = 0;
  for (const line of input) {
    const [inputs, outputs] = line.split(" | ").map(function (part) {
      return part.split(" ");
    });
    for (const permutation of permutations(CHAR_CODES)) {
      const keyToDigits = mapKeysToDigits(permutation);
      const key = inputs.some(function (chars) {
        return !(keyFromChars(chars) in keyToDigits);
      });
      if (key === true) {
        continue;
      }
      outputValuesTotal += parseInt(
        outputs
          .map(function (chars) {
            return keyFromChars(chars);
          })
          .map(function (key) {
            return keyToDigits[key];
          })
          .join(""),
      );
      break;
    }
  }
  return outputValuesTotal;
}
