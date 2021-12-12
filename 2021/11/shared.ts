import { Octupus } from "./Octupus.ts";

export function simulate(
  octupus: Octupus[],
  steps: number,
  fn: (step: number) => boolean = () => false,
) {
  let step = 0;
  for (step = 1; step <= steps; step += 1) {
    for (const oct of octupus) {
      oct.increaseEnergy(step);
    }
    if (fn(step) === true) {
      // Custom function asks us to exit early
      break;
    }
    let stillHasFlash = false;
    do {
      stillHasFlash = octupus.some(function (oct) {
        return oct.needsToFlash();
      });
      for (const oct of octupus) {
        oct.flash(step);
      }
    } while (stillHasFlash);
  }
  return step;
}

export function getOctupus(input: string[]) {
  const matrix = buildMatrix(input);

  const octupusMatrix = buildOctupusMatrix(matrix);

  buildOctupusConnections(octupusMatrix);

  const flatted = octupusMatrix.flat();

  return flatted;
}

function buildMatrix(input: string[]) {
  return input.map(function (line) {
    return line.split("").map(function (oct) {
      return parseInt(oct);
    });
  });
}

function buildOctupusMatrix(matrix: number[][]) {
  return matrix.map(function (line) {
    return line.map(function (value) {
      return new Octupus(value);
    });
  });
}

function buildOctupusConnections(octupusMatrix: Octupus[][]) {
  for (let row = 0; row < octupusMatrix.length; row += 1) {
    for (let col = 0; col < octupusMatrix[row].length; col += 1) {
      const octupus = octupusMatrix[row][col];
      const hasUpperBound = row > 0;
      const hasLowerBound = row < octupusMatrix.length - 1;
      const hasLeftBound = col > 0;
      const hasRightBound = col < octupusMatrix[row].length - 1;
      if (hasUpperBound) {
        octupus.addAdjacent(octupusMatrix[row - 1][col]);
      }
      if (hasLowerBound) {
        octupus.addAdjacent(octupusMatrix[row + 1][col]);
      }
      if (hasLeftBound) {
        octupus.addAdjacent(octupusMatrix[row][col - 1]);
      }
      if (hasRightBound) {
        octupus.addAdjacent(octupusMatrix[row][col + 1]);
      }
      if (hasUpperBound && hasLeftBound) {
        octupus.addAdjacent(octupusMatrix[row - 1][col - 1]);
      }
      if (hasUpperBound && hasRightBound) {
        octupus.addAdjacent(octupusMatrix[row - 1][col + 1]);
      }
      if (hasLowerBound && hasLeftBound) {
        octupus.addAdjacent(octupusMatrix[row + 1][col - 1]);
      }
      if (hasLowerBound && hasRightBound) {
        octupus.addAdjacent(octupusMatrix[row + 1][col + 1]);
      }
    }
  }
  return octupusMatrix;
}
