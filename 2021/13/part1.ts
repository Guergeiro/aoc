import { readInput } from "../utils.ts";
import { Paper } from "./Paper.ts";
import { buildDots, buildFolds, buildMatrix } from "./shared.ts";

export async function AOC2021_13_1(file: string) {
  const input = (await readInput(file)).split("\n");
  return calculate(input);
}

export function calculate(input: string[]) {
  const dots = buildDots(input);
  const folds = buildFolds(input);
  const matrix = buildMatrix(dots, folds);

  const paper = new Paper(matrix);
  const oneFold = folds.shift();
  if (oneFold!.x == null) {
    return paper.foldHorizontal(oneFold!.y).getNumberOfDots();
  }
  return paper.foldVertical(oneFold!.x).getNumberOfDots();
}
