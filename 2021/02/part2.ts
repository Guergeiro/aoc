import { readInput } from "../utils.ts";
import type { Action } from "./Action.ts";

type Submarine = "forward" | "aim" | "depth";

export async function AOC2021_02_2(file: string) {
  const splittedString = (await readInput(file)).split("\n");

  const matrix = splittedString.map(function (line) {
    const [action, value] = line.split(" ");

    return { action: action as Action, value: parseInt(value) };
  });
  return calculate(matrix);
}

export function calculate(matrix: Array<{ action: Action; value: number }>) {
  const map = getMapFromMatrix(matrix);
  return map.get("forward")! * map.get("depth")!;
}

function getMapFromMatrix(matrix: Array<{ action: Action; value: number }>) {
  const m = new Map<Submarine, number>(
    [
      ["forward", 0],
      ["aim", 0],
      ["depth", 0],
    ],
  );
  for (const { action, value } of matrix) {
    switch (action) {
      case "up": {
        m.set("aim", m.get("aim")! - value);
        break;
      }
      case "down":
        m.set("aim", m.get("aim")! + value);
        break;
      default:
        m.set(action, m.get(action)! + value);
        m.set("depth", m.get("aim")! * value + m.get("depth")!);
    }
  }
  return m;
}
