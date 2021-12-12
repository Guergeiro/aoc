import { assertEquals } from "../deps.ts";
import { calculate } from "./part2.ts";

const { test } = Deno;

test("Test Example 2", function () {
  const input = [
    "5483143223",
    "2745854711",
    "5264556173",
    "6141336146",
    "6357385478",
    "4167524645",
    "2176841721",
    "6882881134",
    "4846848554",
    "5283751526",
  ];

  assertEquals(calculate(input), 195);
});
