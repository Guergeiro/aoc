import { AOC2021_01_1 } from "./01/AOC2021_01_1.ts";
import { AOC2021_01_2 } from "./01/AOC2021_01_2.ts";
import { AOC2021_02_1 } from "./02/AOC2021_02_1.ts";
import { AOC2021_02_2 } from "./02/AOC2021_02_2.ts";
import { AOC2021_03_1 } from "./03/AOC2021_03_1.ts";
import { AOC2021_03_2 } from "./03/AOC2021_03_2.ts";
import { AOC2021_04_1 } from "./04/AOC2021_04_1.ts";
import { AOC2021_04_2 } from "./04/AOC2021_04_2.ts";
import { AOC2021_05_1 } from "./05/AOC2021_05_1.ts";
import { AOC2021_05_2 } from "./05/AOC2021_05_2.ts";
import { AOC2021_06_1 } from "./06/AOC2021_06_1.ts";
import { AOC2021_06_2 } from "./06/AOC2021_06_2.ts";
import { AOC2021_07_1 } from "./07/AOC2021_07_1.ts";
import { AOC2021_07_2 } from "./07/AOC2021_07_2.ts";

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
  //https://github.com/FourLineCode/advent-of-code-2021/blob/main/solutions/8.js
}

if (import.meta.main === true) {
  await main();
}
