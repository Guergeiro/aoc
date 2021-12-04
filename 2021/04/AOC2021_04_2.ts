import { readInput } from "../utils.ts";
import { Bingo } from "./Bingo.ts";
import { Builder } from "./Builder.ts";
import { makeBoardMatrix } from "./utils.ts";

export async function AOC2021_04_2(file: string) {
  const splittedString = (await readInput(file)).split("\n");
  const outNumbers = splittedString.shift()!.split(",").map(function (play) {
    return parseInt(play);
  });

  const stringGames = makeBoardMatrix(splittedString);

  const bingo = Builder.buildBingo(stringGames);
  return winnerScore(bingo, outNumbers);
}

export function winnerScore(bingo: Bingo, outNumbers: number[]) {
  while (bingo.everyWinner() === false && outNumbers.length !== 0) {
    const next = outNumbers.shift()!;
    bingo.next(next);
  }

  return bingo.getLastWinnerScore();
}
