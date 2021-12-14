export function numberOfPaths(
  caves: Map<string, string[]>,
  from: string,
  visited: string[],
  dups = false,
) {
  if (from === "end") {
    return 1;
  }
  let sum = 0;

  for (const to of caves.get(from) || []) {
    const thisVisit = [...visited];
    let dupped = dups === true;
    if (to.toLowerCase() === to) {
      if (thisVisit.includes(to) === true) {
        if (dupped && to !== "start" && to !== "end") {
          dupped = false;
        } else {
          continue;
        }
      } else {
        thisVisit.push(to);
      }
    }
    sum += numberOfPaths(caves, to, thisVisit, dupped);
  }

  return sum;
}

export function buildCaves(input: string[]) {
  const caves = new Map<string, string[]>();

  for (const line of input) {
    const [left, right] = line.split("-");
    const leftCaves = caves.get(left) || [];
    const rightCaves = caves.get(right) || [];
    rightCaves.push(left);
    leftCaves.push(right);
    caves.set(left, leftCaves);
    caves.set(right, rightCaves);
  }
  return caves;
}
