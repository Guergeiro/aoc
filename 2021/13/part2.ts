import { readInput } from "../utils.ts";
import { Paper } from "./Paper.ts";
import { buildDots, buildFolds, buildMatrix } from "./shared.ts";

export async function AOC2021_13_2(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const dots = buildDots(input);
  const folds = buildFolds(input);
  const matrix = buildMatrix(dots, folds);

  const paper = folds.reduce(function (paper, cur) {
    if (cur.x == null) {
      return paper.foldHorizontal(cur.y);
    }
    return paper.foldVertical(cur.x);
  }, new Paper(matrix));
  return paper.toString();
}
