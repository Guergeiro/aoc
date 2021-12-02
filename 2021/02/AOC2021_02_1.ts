import { readInput } from "../utils.ts";
import type { Action } from "./Action.ts";

export async function AOC2021_02_1(file: string) {
  const splittedString = (await readInput(file)).split("\n");

  const matrix = splittedString.map(function (line) {
    const [action, value] = line.split(" ");

    return { action: action as Action, value: parseInt(value) };
  });
  return calculate(matrix);
}

export function calculate(matrix: Array<{ action: Action; value: number }>) {
  const map = getMapFromMatrix(matrix);
  const depth = map.get("down")! - map.get("up")!;
  return map.get("forward")! * depth;
}

function getMapFromMatrix(matrix: Array<{ action: Action; value: number }>) {
  const m = new Map<Action, number>(
    [
      ["forward", 0],
      ["down", 0],
      ["up", 0],
    ],
  );
  for (const { action, value } of matrix) {
    const prev = m.get(action);
    if (prev == null) {
      m.set(action, value);
      continue;
    }

    m.set(action, prev + value);
  }
  return m;
}
