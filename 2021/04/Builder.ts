import { Bingo } from "./Bingo.ts";
import { Board } from "./Board.ts";
import { Cell } from "./Cell.ts";

export abstract class Builder {
  public static buildBingo(strBingo: string[][][]) {
    return new Bingo(strBingo.map((board) => {
      return this.buildBoard(board);
    }));
  }

  private static buildBoard(strBoard: string[][]) {
    return new Board(strBoard.map((row) => {
      return this.buildRow(row);
    }));
  }

  private static buildRow(strRow: string[]) {
    return (strRow.map((cell) => {
      return this.buildCell(cell);
    }));
  }

  private static buildCell(strCell: string) {
    return new Cell(parseInt(strCell));
  }
}
