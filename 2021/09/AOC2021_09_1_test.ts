import { assertEquals } from "../deps.ts";
import { calculate } from "./AOC2021_09_1.ts";

const { test } = Deno;

test("Test Example 1", function () {
  const input = [
    "2199943210",
    "3987894921",
    "9856789892",
    "8767896789",
    "9899965678",
  ];

  assertEquals(calculate(input), 15);
});
