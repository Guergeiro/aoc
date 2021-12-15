type Dot = {
  x: number;
  y: number;
};

type Fold = Record<string, number>;

export function buildDots(input: string[]) {
  const dots: Dot[] = [];
  for (const line of input) {
    if (line.length === 0) {
      break;
    }
    const [x, y] = line.split(",");
    dots.push({
      x: parseInt(x),
      y: parseInt(y),
    });
  }
  return dots;
}

export function buildFolds(input: string[]) {
  const folds: Fold[] = [];
  for (const line of input) {
    if (line.startsWith("fold along ") === false) {
      continue;
    }
    const fString = line.replace("fold along ", "");
    const [dir, value] = fString.split("=");
    folds.push({
      [dir]: parseInt(value),
    });
  }
  return folds;
}

export function buildMatrix(dots: Dot[], folds: Fold[]) {
  const maxRow = folds.reduce(function (max, cur) {
    const { y } = cur;
    if (y == null) {
      return max;
    }
    if (y > max) {
      return y;
    }
    return max;
  }, 0);
  const maxCol = folds.reduce(function (max, cur) {
    const { x } = cur;
    if (x == null) {
      return max;
    }
    if (x > max) {
      return x;
    }
    return max;
  }, 0);

  const matrix: string[][] = Array(maxRow * 2 + 1).fill(null).map(function () {
    return Array(maxCol * 2 + 1).fill(null).map(function () {
      return " ";
    });
  });
  for (const dot of dots) {
    matrix[dot.y][dot.x] = "#";
  }

  return matrix;
}
