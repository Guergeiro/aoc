import { Line } from "./Line.ts";

export class Grid {
  readonly #grid: number[][];

  public constructor(grid: number[][]) {
    this.#grid = grid;
  }

  public markLine(line: Line) {
    if (
      line.coordinate1.x !== line.coordinate2.x &&
      line.coordinate1.y !== line.coordinate2.y
    ) {
      return;
    }
    const [yStart, yEnd] = this.getStartAndEnd([
      line.coordinate1.y,
      line.coordinate2.y,
    ]);
    const [xStart, xEnd] = this.getStartAndEnd([
      line.coordinate1.x,
      line.coordinate2.x,
    ]);

    for (let y = yStart; y <= yEnd; y += 1) {
      for (let x = xStart; x <= xEnd; x += 1) {
        this.#grid[y][x] += 1;
      }
    }
  }

  public markDiagonal(line: Line) {
    if (
      Math.abs(line.coordinate1.x - line.coordinate2.x) !==
        Math.abs(line.coordinate1.y - line.coordinate2.y)
    ) {
      return;
    }

    const x1 = line.coordinate1.x;
    const x2 = line.coordinate2.x;
    const y1 = line.coordinate1.y;
    const y2 = line.coordinate2.y;

    const deltaX = x2 === x1 ? 0 : x2 < x1 ? -1 : 1;
    const deltaY = y2 === y1 ? 0 : y2 < y1 ? -1 : 1;

    for (
      let x = x1, y = y1;
      x !== x2 + deltaX && y !== y2 + deltaY;
      x += deltaX, y += deltaY
    ) {
      this.#grid[y][x] += 1;
    }
  }

  public getGrid() {
    return this.#grid;
  }

  public print() {
    const grid: string[] = [];
    for (const row of this.#grid) {
      const rowStr: string[] = [];
      for (const cell of row) {
        rowStr.push(`${cell}`);
      }
      grid.push(rowStr.join(" "));
    }
    return grid.join("\n");
  }

  public getSumOverlaps() {
    let sum = 0;

    for (const row of this.#grid) {
      for (const cell of row) {
        if (cell > 1) {
          sum += 1;
        }
      }
    }

    return sum;
  }

  private getStartAndEnd(pair: number[]) {
    return pair.sort(function (a, b) {
      return a - b;
    });
  }
}
