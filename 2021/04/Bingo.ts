import { Board } from "./Board.ts";

export class Bingo {
  readonly #boards: Board[];
  readonly #winners: Board[] = [];
  #previousValue = 0;

  public constructor(boards: Board[]) {
    this.#boards = boards;
  }

  public hasWinner() {
    return this.#boards.some((board) => {
      return this.boardIsFinished(board);
    });
  }

  public everyWinner() {
    let everyWinner = true;
    for (const board of this.#boards) {
      const isFinished = this.boardIsFinished(board);
      if (isFinished === false) {
        everyWinner = false;
        continue;
      }
      this.addToWinners(board);
    }
    return everyWinner;
  }

  public next(value: number) {
    for (const board of this.#boards) {
      board.mark(value);
    }
    this.#previousValue = value;
  }

  public getWinnerScore() {
    for (const board of this.#boards) {
      if (board.isFinished()) {
        return this.calculateScore(board);
      }
    }
  }

  public getLastWinnerScore() {
    const lastWinner = this.#winners.pop()!;
    return this.calculateScore(lastWinner);
  }

  private addToWinners(board: Board) {
    const alreadyExists = this.#winners.some(function (b) {
      return b === board;
    });
    if (alreadyExists === false) {
      this.#winners.push(board);
    }
  }

  private calculateScore(board: Board) {
    return board.getScore() * this.#previousValue;
  }

  private boardIsFinished(board: Board) {
    return board.isFinished();
  }
}
