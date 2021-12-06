import { Line } from "./Line.ts";

export function findBoundaries(lines: Line[]) {
  let maxX = 0;
  let maxY = 0;

  for (const line of lines) {
    for (const coordinate of line.getTuple()) {
      if (coordinate.x > maxX) {
        maxX = coordinate.x;
      }
      if (coordinate.y > maxY) {
        maxY = coordinate.y;
      }
    }
  }
  return { maxX, maxY };
}
