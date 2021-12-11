import { readInput } from "../utils.ts";
import { Coordinate } from "./Coordinate.ts";
import { findBoundaries } from "./utils.ts";
import { Grid } from "./Grid.ts";
import { Line } from "./Line.ts";

export async function AOC2021_05_1(file: string) {
  const splittedString = (await readInput(file)).split("\n");
  return execute(splittedString);
}

export function execute(input: string[]) {
  const lines: Line[] = [];
  for (const row of input) {
    const [start, end] = row.split(" -> ");
    const [x1, y1] = start.split(",");
    const [x2, y2] = end.split(",");

    const coordinate1 = new Coordinate(parseInt(x1), parseInt(y1));
    const coordinate2 = new Coordinate(parseInt(x2), parseInt(y2));
    lines.push(new Line(coordinate1, coordinate2));
  }

  const { maxX, maxY } = findBoundaries(lines);

  const grid = new Grid(
    Array(maxY + 1).fill(null).map(function () {
      return Array(maxX + 1).fill(0);
    }),
  );
  for (const line of lines) {
    grid.markLine(line);
  }

  return grid.getSumOverlaps();
}
