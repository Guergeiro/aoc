import { assertEquals } from "../deps.ts";
import { simulateLifecycle, simulateLifecycleOld } from "./shared.ts";

const { test } = Deno;

test("Test Example 18 days", function () {
  const input = [3, 4, 3, 1, 2];

  assertEquals(simulateLifecycle(input, 18), 26);
});
test("Test Example 80 days", function () {
  const input = [3, 4, 3, 1, 2];

  assertEquals(simulateLifecycle(input, 80), 5934);
});
test("Test Example 18 days (old)", function () {
  const input = [3, 4, 3, 1, 2];

  assertEquals(simulateLifecycleOld(input, 18).length, 26);
});
test("Test Example 80 days (Old)", function () {
  const input = [3, 4, 3, 1, 2];

  assertEquals(simulateLifecycleOld(input, 80).length, 5934);
});
