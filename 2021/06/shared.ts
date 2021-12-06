export function simulateLifecycle(fishtank: number[], days: number) {
  const timers: number[] = Array(9).fill(0);
  for (const fish of fishtank) {
    // Populate timers
    timers[fish] += 1;
  }
  for (let i = 0; i < days; i++) {
    const fishThatReachedZero = timers.shift()!;
    timers[6] += fishThatReachedZero;
    timers.push(fishThatReachedZero);
  }
  return timers.reduce(function (prev, curr) {
    return prev + curr;
  });
}

export function simulateLifecycleOld(fishtank: number[], days: number) {
  for (let i = 0; i < days; i++) {
    const previous = [...fishtank];
    for (const fish in fishtank) {
      previous[fish] -= 1;
      if (previous[fish] === -1) {
        previous[fish] = 6;
        previous.push(8);
      }
    }
    fishtank = [...previous];
  }
  return fishtank;
}
