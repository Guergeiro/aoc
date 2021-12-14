import { assertEquals } from "../deps.ts";
import { calculate } from "./part2.ts";

const { test } = Deno;

test("Test example 1", function () {
  const input = ["start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"];
  assertEquals(calculate(input), 36);
});
test("Test example 2", function () {
  const input = [
    "dc-end",
    "HN-start",
    "start-kj",
    "dc-start",
    "dc-HN",
    "LN-dc",
    "HN-end",
    "kj-sa",
    "kj-HN",
    "kj-dc",
  ];
  assertEquals(calculate(input), 103);
});
test("Test example 3", function () {
  const input = [
    "fs-end",
    "he-DX",
    "fs-he",
    "start-DX",
    "pj-DX",
    "end-zg",
    "zg-sl",
    "zg-pj",
    "pj-he",
    "RW-he",
    "fs-DX",
    "pj-RW",
    "zg-RW",
    "start-pj",
    "he-WI",
    "zg-he",
    "pj-fs",
    "start-RW",
  ];
  assertEquals(calculate(input), 3509);
});
