export class Paper {
  readonly #matrix: string[][];

  public constructor(matrix: string[][]) {
    this.#matrix = matrix;
  }

  public foldHorizontal(y: number) {
    const top = [...this.#matrix].slice(0, y);
    const bottom = this.mirrorHorizontal([...this.#matrix].slice(y + 1));
    const newMatrix = [...top].map((row, i) => {
      return row.map((cel, j) => {
        if (this.isDot(bottom[i][j])) {
          return bottom[i][j];
        }
        return cel;
      });
    });

    const paper = new Paper(newMatrix);
    return paper;
  }

  public foldVertical(x: number) {
    const left = [...this.#matrix].map(function (row) {
      return row.slice(0, x);
    });
    const right = this.mirrorVertical([...this.#matrix].map(function (row) {
      return row.slice(x + 1);
    }));

    const newMatrix = [...left].map((row, i) => {
      return row.map((cel, j) => {
        if (this.isDot(right[i][j])) {
          return "#";
        }
        return cel;
      });
    });

    const paper = new Paper(newMatrix);
    return paper;
  }

  public getNumberOfDots() {
    let sum = 0;
    for (const line of this.#matrix) {
      for (const cel of line) {
        if (this.isDot(cel)) {
          sum += 1;
        }
      }
    }
    return sum;
  }

  private isDot(cell: string) {
    return cell === "#";
  }

  private mirrorHorizontal(matrix: string[][]) {
    return matrix.reverse();
  }

  private mirrorVertical(matrix: string[][]) {
    return matrix.map(function (row) {
      return row.reverse();
    });
  }

  public toString() {
    return this.#matrix.map(function (row) {
      return row.join("");
    });
  }
}
