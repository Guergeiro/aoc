import { assertEquals } from "../deps.ts";
import { simulateLifecycle, simulateLifecycleOld } from "./shared.ts";

const { test } = Deno;

test("Test Example 18 days", function () {
  const input = [3, 4, 3, 1, 2];

  assertEquals(simulateLifecycle(input, 18), 26);
});
test("Test Example 256 days", function () {
  const input = [3, 4, 3, 1, 2];

  assertEquals(simulateLifecycle(input, 256), 26984457539);
});
test({
  name: "Test Example 256 days (Old)",
  ignore: true,
  fn: function () {
    const input = [3, 4, 3, 1, 2];

    assertEquals(simulateLifecycleOld(input, 256).length, 26984457539);
  },
});
