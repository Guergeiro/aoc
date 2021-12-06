export class Coordinate {
  readonly #x: number;
  readonly #y: number;

  public constructor(x: number, y: number) {
    this.#x = x;
    this.#y = y;
  }

  get x() {
    return this.#x;
  }

  get y() {
    return this.#y;
  }
}
