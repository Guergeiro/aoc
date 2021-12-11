import { assertEquals } from "../deps.ts";
import { calculate } from "./part2.ts";

const { test } = Deno;

test("Test Example 2", function () {
  const input = [
    "2199943210",
    "3987894921",
    "9856789892",
    "8767896789",
    "9899965678",
  ];

  assertEquals(calculate(input), 1134);
});
