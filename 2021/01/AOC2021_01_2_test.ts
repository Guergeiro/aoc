import { assertEquals } from "../deps.ts";
import { calculate } from "./AOC2021_01_2.ts";

const { test } = Deno;

test("Test Example", function () {
  const input = [
    199,
    200,
    208,
    210,
    200,
    207,
    240,
    269,
    260,
    263,
  ];

  assertEquals(calculate(input), 5);
});

test("Test Everything Decreased", function () {
  const input = [
    269,
    263,
    260,
    240,
    210,
    208,
    207,
    200,
    200,
    199,
  ];

  assertEquals(calculate(input), 0);
});

test("Test Everything Increased", function () {
  const input = [
    199,
    200,
    201,
    207,
    208,
    210,
    240,
    260,
    263,
    269,
  ];

  assertEquals(calculate(input), 7);
});
