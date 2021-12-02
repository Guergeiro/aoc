import { assertEquals } from "../deps.ts";
import { Action } from "./Action.ts";
import { calculate } from "./AOC2021_02_2.ts";

const { test } = Deno;

test("Test Example", function () {
  const m: Array<{ action: Action; value: number }> = [
    {
      action: "forward",
      value: 5,
    },
    {
      action: "down",
      value: 5,
    },
    {
      action: "forward",
      value: 8,
    },
    {
      action: "up",
      value: 3,
    },
    { action: "down", value: 8 },
    { action: "forward", value: 2 },
  ];

  assertEquals(calculate(m), 900);
});
