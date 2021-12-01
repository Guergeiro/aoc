export async function readInput(fileLocation: string) {
  const input = await Deno.readTextFile(fileLocation);
  const splittedString: string[] = input.split("\n").slice(0, -1);
  return splittedString.join("\n");
}
