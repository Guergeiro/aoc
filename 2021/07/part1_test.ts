import { assertEquals } from "../deps.ts";
import { execute } from "./part1.ts";

const { test } = Deno;

test("Test Example", function () {
  const input = [16, 1, 2, 0, 4, 2, 7, 1, 2, 14];

  assertEquals(execute(input), 37);
});
