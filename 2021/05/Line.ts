import { Coordinate } from "./Coordinate.ts";

export class Line {
  readonly #coordinate1: Coordinate;
  readonly #coordinate2: Coordinate;

  public constructor(coordinate1: Coordinate, coordinate2: Coordinate) {
    this.#coordinate1 = coordinate1;
    this.#coordinate2 = coordinate2;
  }

  public getTuple() {
    return [this.#coordinate1, this.#coordinate2];
  }

  get coordinate1() {
    return this.#coordinate1;
  }

  get coordinate2() {
    return this.#coordinate2;
  }
}
