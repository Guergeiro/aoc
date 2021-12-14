import { AOC2021_01_1 } from "./01/part1.ts";
import { AOC2021_01_2 } from "./01/part2.ts";
import { AOC2021_02_1 } from "./02/part1.ts";
import { AOC2021_02_2 } from "./02/part2.ts";
import { AOC2021_03_1 } from "./03/part1.ts";
import { AOC2021_03_2 } from "./03/part2.ts";
import { AOC2021_04_1 } from "./04/part1.ts";
import { AOC2021_04_2 } from "./04/part2.ts";
import { AOC2021_05_1 } from "./05/part1.ts";
import { AOC2021_05_2 } from "./05/part2.ts";
import { AOC2021_06_1 } from "./06/part1.ts";
import { AOC2021_06_2 } from "./06/part2.ts";
import { AOC2021_07_1 } from "./07/part1.ts";
import { AOC2021_07_2 } from "./07/part2.ts";
import { AOC2021_08_1 } from "./08/part1.ts";
import { AOC2021_08_2 } from "./08/part2.ts";
import { AOC2021_09_1 } from "./09/part1.ts";
import { AOC2021_09_2 } from "./09/part2.ts";
import { AOC2021_10_1 } from "./10/part1.ts";
import { AOC2021_10_2 } from "./10/part2.ts";
import { AOC2021_11_1 } from "./11/part1.ts";
import { AOC2021_11_2 } from "./11/part2.ts";
import { AOC2021_12_1 } from "./12/part1.ts";
import { AOC2021_12_2 } from "./12/part2.ts";

async function main() {
  console.log(
    "01_1",
    await AOC2021_01_1(await Deno.realPath("./01/input_question.txt")),
  );
  console.log(
    "01_2",
    await AOC2021_01_2(await Deno.realPath("./01/input_question.txt")),
  );
  console.log(
    "02_1",
    await AOC2021_02_1(await Deno.realPath("./02/input_question.txt")),
  );
  console.log(
    "02_2",
    await AOC2021_02_2(await Deno.realPath("./02/input_question.txt")),
  );
  console.log(
    "03_1",
    await AOC2021_03_1(await Deno.realPath("./03/input_question.txt")),
  );
  console.log(
    "03_2",
    await AOC2021_03_2(await Deno.realPath("./03/input_question.txt")),
  );
  console.log(
    "04_1",
    await AOC2021_04_1(await Deno.realPath("./04/input_question.txt")),
  );
  console.log(
    "04_2",
    await AOC2021_04_2(await Deno.realPath("./04/input_question.txt")),
  );
  console.log(
    "05_1",
    await AOC2021_05_1(await Deno.realPath("./05/input_question.txt")),
  );
  console.log(
    "05_2",
    await AOC2021_05_2(await Deno.realPath("./05/input_question.txt")),
  );
  console.log(
    "06_1",
    await AOC2021_06_1(await Deno.realPath("./06/input_question.txt")),
  );
  console.log(
    "06_2",
    await AOC2021_06_2(await Deno.realPath("./06/input_question.txt")),
  );
  console.log(
    "07_1",
    await AOC2021_07_1(await Deno.realPath("./07/input_question.txt")),
  );
  console.log(
    "07_2",
    await AOC2021_07_2(await Deno.realPath("./07/input_question.txt")),
  );
  console.log(
    "08_1",
    await AOC2021_08_1(await Deno.realPath("./08/input_question.txt")),
  );
  // Inspired by: https://github.com/FourLineCode/advent-of-code-2021/blob/main/solutions/8.js
  console.log(
    "08_2",
    await AOC2021_08_2(await Deno.realPath("./08/input_question.txt")),
  );
  console.log(
    "09_1",
    await AOC2021_09_1(await Deno.realPath("./09/input_question.txt")),
  );
  console.log(
    "09_2",
    await AOC2021_09_2(await Deno.realPath("./09/input_question.txt")),
  );
  console.log(
    "10_1",
    await AOC2021_10_1(await Deno.realPath("./10/input_question.txt")),
  );
  console.log(
    "10_2",
    await AOC2021_10_2(await Deno.realPath("./10/input_question.txt")),
  );
  console.log(
    "11_1",
    await AOC2021_11_1(await Deno.realPath("./11/input_question.txt")),
  );
  console.log(
    "11_2",
    await AOC2021_11_2(await Deno.realPath("./11/input_question.txt")),
  );
  console.log(
    "12_1",
    await AOC2021_12_1(await Deno.realPath("./12/input_question.txt")),
  );
  console.log(
    "12_2",
    await AOC2021_12_2(await Deno.realPath("./12/input_question.txt")),
  );
}

if (import.meta.main === true) {
  await main();
}
