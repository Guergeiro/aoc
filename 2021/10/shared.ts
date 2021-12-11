const openPairs = new Map<string, string>([
  ["(", ")"],
  ["[", "]"],
  ["{", "}"],
  ["<", ">"],
]);
const closePairs = new Map<string, string>([
  [")", "("],
  ["]", "["],
  ["}", "{"],
  [">", "<"],
]);
const corruptedValue = new Map<string, number>([
  [")", 3],
  ["]", 57],
  ["}", 1197],
  [">", 25137],
]);
const missingValue = new Map<string, number>([
  [")", 1],
  ["]", 2],
  ["}", 3],
  [">", 4],
]);

export function isPair(left?: string, right?: string) {
  if (left == null) {
    return false;
  }
  if (right == null) {
    return false;
  }
  const rightCandidate = openPairs.get(left);
  if (rightCandidate == null) {
    return false;
  }
  return rightCandidate === right;
}

export function isNotPair(left?: string, right?: string) {
  return isPair(left, right) === false;
}

export function isOpen(char: string) {
  return openPairs.has(char);
}

export function isClose(char: string) {
  return closePairs.has(char);
}

export function getCorruptedValue(char: string) {
  const v = corruptedValue.get(char);
  if (v == null) {
    throw new Error();
  }
  return v;
}

export function getCorruptedChar(line: string) {
  const queue: string[] = [];
  for (const char of line) {
    if (isOpen(char) === true) {
      queue.push(char);
    }
    if (isClose(char) === true) {
      const left = queue.pop();
      if (isNotPair(left, char) === true) {
        return char;
      }
    }
  }
}

export function getMissingClosingChars(line: string) {
  const allClosingChars: string[] = [];

  for (const char of line) {
    if (isOpen(char) === true) {
      allClosingChars.push(openPairs.get(char)!);
    }
    if (isClose(char) === true) {
      allClosingChars.pop();
    }
  }
  return allClosingChars.reverse();
}

export function getMissingValue(char: string) {
  const v = missingValue.get(char);
  if (v == null) {
    throw new Error();
  }
  return v;
}
