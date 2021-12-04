export class Cell {
  readonly #value: number;
  #marked = false;

  public constructor(value: number) {
    this.#value = value;
  }

  public mark(value: number) {
    if (this.#marked === false) {
      this.#marked = this.#value === value;
    }
  }

  public isFinished() {
    return this.#marked;
  }

  public getScore() {
    if (this.#marked === false) {
      return this.#value;
    }
    return 0;
  }
}
