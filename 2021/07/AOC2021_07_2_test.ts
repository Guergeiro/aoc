import { assertEquals } from "../deps.ts";
import { execute } from "./AOC2021_07_2.ts";

const { test } = Deno;

test("Test Example", function () {
  const input = [16, 1, 2, 0, 4, 2, 7, 1, 2, 14];

  assertEquals(execute(input), 168);
});
