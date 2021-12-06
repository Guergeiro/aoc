import { assertEquals } from "../deps.ts";
import { calculatePowerConsumption } from "./AOC2021_03_2.ts";

const { test } = Deno;

test("Test Example", function () {
  const m = [
    "00100",
    "11110",
    "10110",
    "10111",
    "10101",
    "01111",
    "00111",
    "11100",
    "10000",
    "11001",
    "00010",
    "01010",
  ];

  assertEquals(calculatePowerConsumption(m), 230);
});