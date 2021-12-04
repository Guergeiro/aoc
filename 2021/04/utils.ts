export function makeBoardMatrix(lines: string[]) {
  const stringGames: string[][] = [];

  let idx = -1;
  for (const row of lines) {
    if (row.length === 0) {
      // New game
      stringGames.push([]);
      idx += 1;
      continue;
    }
    stringGames[idx].push(row);
  }

  const finalGames = stringGames.map(function (game) {
    return game.map(function (gameRow) {
      const row = gameRow.split(" ");
      return row.filter(function (cell) {
        return isNaN(parseInt(cell)) === false;
      });
    });
  });

  return finalGames;
}
