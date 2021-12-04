import { Cell } from "./Cell.ts";

export class Board {
  readonly #board: Cell[][];

  public constructor(board: Cell[][]) {
    this.#board = board;
  }

  public mark(value: number) {
    for (const row of this.#board) {
      for (const cell of row) {
        cell.mark(value);
      }
    }
  }

  public getScore() {
    let sum = 0;
    for (const row of this.#board) {
      if (this.completed(row) === false) {
        sum += this.getUnmarkedSum(row);
      }
    }
    return sum;
  }

  public isFinished() {
    for (const row of this.#board) {
      if (this.completed(row)) {
        return true;
      }
    }
    for (const col of this.getTranspose()) {
      if (this.completed(col)) {
        return true;
      }
    }
    return false;
  }

  private completed(strip: Cell[]) {
    return strip.every(function (cell) {
      return cell.isFinished();
    });
  }

  private getTranspose() {
    return this.#board[0].map((_, j) => {
      return this.#board.map((row) => row[j]);
    });
  }

  private getUnmarkedSum(strip: Cell[]) {
    return strip.reduce(function (sum, cur) {
      return sum + cur.getScore();
    }, 0);
  }
}
