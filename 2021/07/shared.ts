export function calculate(
  input: number[],
  logic: (acc: number, curr: number, depth: number) => number,
) {
  const min = Math.min(...input);
  const max = Math.max(...input);

  const totals: number[] = [];
  for (let depth = min; depth < max; depth += 1) {
    const sum = input.reduce(function (acc, curr) {
      return logic(acc, curr, depth);
    }, 0);
    totals.push(sum);
  }
  return Math.min(...totals);
}
