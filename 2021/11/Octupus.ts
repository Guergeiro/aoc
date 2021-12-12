export class Octupus {
  #energy: number;
  #needsToFlash = false;
  #lastFlashedDay = 0;
  #flashes = 0;
  readonly #adjacent: Octupus[] = [];

  public constructor(energy: number) {
    this.#energy = energy;
  }

  public increaseEnergy(step: number) {
    if (this.#lastFlashedDay === step) {
      // Already flashed today
      return;
    }
    this.#energy += 1;
    if (this.#energy <= 9) {
      // Do not flag me as ready to flash
      return;
    }
    this.#needsToFlash = true;
  }

  public addAdjacent(octupus: Octupus) {
    this.#adjacent.push(octupus);
  }

  public flash(step: number) {
    if (this.#needsToFlash === false) {
      // I'm not ready to flash
      return;
    }
    this.#lastFlashedDay = step;
    this.#needsToFlash = false;
    for (const adj of this.#adjacent) {
      // Since I flashed, I'll ask everyone to increase their energy!
      adj.increaseEnergy(step);
    }
    this.#energy = 0;
    this.#flashes += 1;
  }

  public needsToFlash() {
    return this.#needsToFlash;
  }

  public getFlashes() {
    return this.#flashes;
  }

  public toString() {
    return this.#energy.toString();
  }
}
