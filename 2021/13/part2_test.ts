import { assertEquals } from "../deps.ts";
import { calculate } from "./part2.ts";

const { test } = Deno;

test("Test Example 1", function () {
  const input = [
    "6,10",
    "0,14",
    "9,10",
    "0,3",
    "10,4",
    "4,11",
    "6,0",
    "6,12",
    "4,1",
    "0,13",
    "10,12",
    "3,4",
    "3,0",
    "8,4",
    "1,10",
    "2,14",
    "8,10",
    "9,0",
    "",
    "fold along y=7",
  ];

  assertEquals(calculate(input), 17);
});
test("Test Example 2", function () {
  const input = [
    "6,10",
    "0,14",
    "9,10",
    "0,3",
    "10,4",
    "4,11",
    "6,0",
    "6,12",
    "4,1",
    "0,13",
    "10,12",
    "3,4",
    "3,0",
    "8,4",
    "1,10",
    "2,14",
    "8,10",
    "9,0",
    "",
    "fold along y=7",
    "fold along x=5",
  ];

  assertEquals(calculate(input), 16);
});
