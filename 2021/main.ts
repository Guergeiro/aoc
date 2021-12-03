import { AOC2021_01_1 } from "./01/AOC2021_01_1.ts";
import { AOC2021_01_2 } from "./01/AOC2021_01_2.ts";
import { AOC2021_02_1 } from "./02/AOC2021_02_1.ts";
import { AOC2021_02_2 } from "./02/AOC2021_02_2.ts";
import { AOC2021_03_1 } from "./03/AOC2021_03_1.ts";
import { AOC2021_03_2 } from "./03/AOC2021_03_2.ts";

async function main() {
  console.log(
    await AOC2021_01_1(await Deno.realPath("./01/input_example.txt")),
  );
  console.log(
    await AOC2021_01_1(await Deno.realPath("./01/input_question.txt")),
  );
  console.log(
    await AOC2021_01_2(await Deno.realPath("./01/input_example.txt")),
  );
  console.log(
    await AOC2021_01_2(await Deno.realPath("./01/input_question.txt")),
  );
  console.log(
    await AOC2021_02_1(await Deno.realPath("./02/input_example.txt")),
  );
  console.log(
    await AOC2021_02_1(await Deno.realPath("./02/input_question.txt")),
  );
  console.log(
    await AOC2021_02_2(await Deno.realPath("./02/input_example.txt")),
  );
  console.log(
    await AOC2021_02_2(await Deno.realPath("./02/input_question.txt")),
  );
  console.log(
    await AOC2021_03_1(await Deno.realPath("./03/input_example.txt")),
  );

  console.log(
    await AOC2021_03_1(await Deno.realPath("./03/input_question.txt")),
  );
  console.log(
    await AOC2021_03_2(await Deno.realPath("./03/input_example.txt")),
  );
  console.log(
    await AOC2021_03_2(await Deno.realPath("./03/input_question.txt")),
  );
}

if (import.meta.main === true) {
  await main();
}
